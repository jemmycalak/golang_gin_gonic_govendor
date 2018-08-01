package user

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/go_gin_govendor/src/models"
	"github.com/jemmycalak/go_gin_govendor/src/repositorys"
)

var msg, stts string

type UserController struct {
	DB *sql.DB
}

func NewUserContrroller(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Users(c *gin.Context) {
	users, err := repositorys.ShowUsers(uc.DB)
	if err != nil {
		log.Println("error", err)
		c.String(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) AddUser(c *gin.Context) {
	data := models.NewUser()
	c.BindJSON(&data)

	err := repositorys.AddUser(uc.DB, data)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Faild add user")
		return
	}

	msg = "data saved"
	stts = "true"
	result := map[string]string{"status": stts, "msg": msg}

	ResponseSuccess(c, http.StatusOK, result)

}

func (uc *UserController) UpdateUser(c *gin.Context) {
	data := models.NewUser()
	c.BindJSON(&data)

	err := repositorys.UpdateUser(uc.DB, data.Id, data)

	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Something error")
		return
	}
	msg = "data updated"
	stts = "true"
	result := map[string]string{"msg": msg, "status": stts}
	ResponseSuccess(c, http.StatusOK, result)
}

func (uc *UserController) DeletUser(c *gin.Context) {
	data := models.NewUser()
	c.BindJSON(&data)

	err := repositorys.DeleteUser(uc.DB, data.Id)
	log.Println(data.Id)

	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Faild to delete data")
		return
	}

	msg = "date deleted"
	stts = "true"
	result := map[string]string{"msg": msg, "status": stts}
	ResponseSuccess(c, http.StatusOK, result)
}

func (uc *UserController) UserById(c *gin.Context) {
	iduser := c.Params.ByName("iduser")
	miduser, err := strconv.Atoi(iduser) //convert to int
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Error userid")
		return
	}

	model, err := repositorys.FindUserById(uc.DB, miduser)
	if err != nil {
		if err == sql.ErrNoRows {
			ResponseError(c, http.StatusBadRequest, "No data found")
			return
		}
		ResponseError(c, http.StatusInternalServerError, "Faild find data")
		return
	}

	ResponseSuccess(c, http.StatusOK, model)

}

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status": "false",
		"msg":    msg,
	})
}

func ResponseSuccess(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
