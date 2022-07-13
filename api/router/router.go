package router

import (
	"api/internal/pkg/handler"
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"
	"api/pkg/share/validators"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func (r *Router) Routes() {

	r.DB.MigrateDBWithGorm()

	validators.CustomValidate()
	r.DB.MigrateDBWithGorm()

	hUserCus := handler.NewCustomerHandler(r.DB)
	hUserAdmin := handler.NewAdminHandler(r.DB)
	hCategory := handler.NewCategoryHandler(r.DB)
	hBrand := handler.NewBrandHandler(r.DB)
	hProduct := handler.NewProductHandler(r.DB)
	api := r.Engine.Group("/api")
	{
		api.GET("/get_categories", hCategory.GetCategoryList)
		api.GET("/get_brands", hBrand.GetBrandList)
		api.GET("/search_product", hProduct.GetProductList)
		api.GET("/product_info/:product_id", hProduct.GetProductInfo)
		customerAPI := api.Group("/customer")
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
		adminAPI := api.Group("/admin")
		{
			adminAPI.POST("/login", hUserAdmin.Login)
			adminAPI.Use(middleware.AuthMiddleware(r.DB), middleware.AuthAdminMiddleware(r.DB))
			{
				adminAPI.GET("/profile", hUserAdmin.GetProfile)
				adminAPI.POST("/create_user", hUserAdmin.CreateUser)
				adminAPI.PATCH("/update_profile", hUserAdmin.UpdateProfile)
				adminAPI.PATCH("/change_password", hUserAdmin.ChangePassWord)
				categoryAPI := adminAPI.Group("/category")
				{
					categoryAPI.POST("/create_category", hCategory.CreateCategory)
					categoryAPI.PATCH("/update_category", hCategory.UpdateCategory)
					categoryAPI.DELETE("/delete_category", hCategory.DeleteCategory)
				}
				brandAPI := adminAPI.Group("/brand")
				{
					brandAPI.POST("/create_brand", hBrand.CreateBrand)
					brandAPI.PATCH("/update_brand", hBrand.UpdateBrand)
					brandAPI.DELETE("/delete_brand", hBrand.DeleteBrand)
				}
				productAPI := adminAPI.Group("/product")
				{
					productAPI.POST("/create_product", hProduct.CreateProduct)
					productAPI.PATCH("/update_product", hProduct.UpdateProduct)
					productAPI.DELETE("/delete_product", hProduct.DeleteProduct)
				}
			}
		}
	}

}
func NewRouter() Router {
	var r Router
	r.Engine = gin.Default()
	database, err := db.NewDB()
	if err != nil {
		return Router{}
	}
	r.DB = database
	r.Routes()
	return r
}
