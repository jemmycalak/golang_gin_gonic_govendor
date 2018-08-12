package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jemmycalak/go_gin_govendor/src/controllers/user"
	"github.com/jemmycalak/go_gin_govendor/src/handlers"
)

func UserRouters(r *gin.Engine, ur *user.UserController) {

	//versi apps
	r.Use(handlers.ApikeyValidator())
	r.POST("/api/v1/user/login", ur.Login)
	r.POST("/api/v1/user/register", ur.Register)

	v1 := r.Group("/api/v1")
	v1.Use(handlers.TokenValidator())
	{
		v1.GET("/", ur.Users)
		v1.POST("/AddUser", ur.AddUser)
		v1.POST("/UpdateUser", ur.UpdateUser)
		v1.POST("/DeleteUser", ur.DeletUser)
		v1.GET("/FindUserId/:iduser", ur.UserById)
	}
	r.NoRoute(func(c *gin.Context) {
		user.ResponseError(c, http.StatusNotFound, "I don't know what are you looking for")
	})
}
