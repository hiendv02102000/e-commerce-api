package router

import (
	"api/internal/pkg/handler"
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"
	"api/pkg/share/validators"

	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Routes() {
	var err error
	r.Engine = gin.Default()
	r.DB, err = db.NewDB()
	r.DB.MigrateDBWithGorm()
	if err != nil {
		fmt.Println(err)
	}
	validators.CustomValidate()
	r.DB.MigrateDBWithGorm()

	hUserCus := handler.NewCustomerHandler(r.DB)
	hUserAdmin := handler.NewAdminHandler(r.DB)
	customerAPI := r.Engine.Group("/customer")
	{

		customerAPI.POST("/login", hUserCus.Login)
		customerAPI.POST("/register", hUserCus.CreateUser)
		customerAPI.Use(middleware.AuthMiddleware(r.DB))
		{
			customerAPI.GET("/profile", hUserCus.GetProfile)
			customerAPI.PATCH("/update_profile", hUserCus.UpdateProfile)
			customerAPI.PATCH("/change_password", hUserAdmin.ChangePassWord)
		}

	}
	adminAPI := r.Engine.Group("/admin")
	{
		adminAPI.POST("/login", hUserAdmin.Login)
		adminAPI.Use(middleware.AuthMiddleware(r.DB), middleware.AuthAdminMiddleware(r.DB))
		{
			adminAPI.GET("/profile", hUserAdmin.GetProfile)
			adminAPI.POST("/create_user", hUserAdmin.CreateUser)
			adminAPI.PATCH("/update_profile", hUserAdmin.UpdateProfile)
			adminAPI.PATCH("/change_password", hUserAdmin.ChangePassWord)
		}
	}
}
func NewRouter() Router {
	var r Router
	r.Routes()
	return r
}
