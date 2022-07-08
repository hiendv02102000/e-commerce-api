package router

import (
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"

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
	if err != nil {
		fmt.Println(err)
	}
	r.DB.MigrateDBWithGorm()

	// h := handler.NewHTTPHandler(r.DB)

	webAPI := r.Engine.Group("/app")
	{

		musicAPI := webAPI.Group("/account")
		{
			musicAPI.Use(middleware.AuthMiddleware(r.DB))
			{
			}
		}

	}
}
func NewRouter() Router {
	var r Router
	r.Routes()
	return r
}
