package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/jemmycalak/go_gin_govendor/src/controllers/user"
)

func CreateUserRouters(r *gin.Engine, ur *user.UserController) {
	r.GET("/", ur.Users)
	r.POST("/AddUser", ur.AddUser)
	r.POST("/UpdateUser", ur.UpdateUser)
	r.POST("/DeleteUser", ur.DeletUser)
	r.GET("/FindUserId/:iduser", ur.UserById)
}
