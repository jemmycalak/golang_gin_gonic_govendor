package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/go_gin_govendor/src/handlers"
	"github.com/jemmycalak/go_gin_govendor/src/models"
	"github.com/jemmycalak/go_gin_govendor/src/repositorys"
	"github.com/jemmycalak/go_gin_govendor/src/utils"
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
	if err := c.BindJSON(&data); err != nil {
		ResponseError(c, http.StatusBadRequest, "Data not complete")
		return
	}

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

func (uc *UserController) Login(c *gin.Context) {

	var model models.LoginStruct
	// var newModel []models.Login
	err := c.BindJSON(&model)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Data not valid")
	}

	// newModel := models.LoginStruct{
	// 	Email:    model.Email,
	// 	Password: newPassword,
	// }
	// fmt.Println(model)
	newModel, err := repositorys.LoginRepository(uc.DB, &model)
	if err != nil {
		ResponseError(c, http.StatusNotFound, "Email dosen't exsit")
		c.Abort()
		return
	}
	fmt.Println(newModel.Password)

	chekPassword, err := utils.Decrypt(newModel.Password)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "Something wrong, please try again")
		c.Abort()
		return
	}

	if model.Password != chekPassword {
		ResponseError(c, http.StatusConflict, "Password invalid")
		c.Abort()
		return
	}

	newId := newModel.Id
	convToInt, err := strconv.Atoi(newId)
	if err != nil {
		fmt.Println("error convert to integer")
	}

	token, err := handlers.GenerateJWT(convToInt)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "Something wrong, please try again")
		c.Abort()
		return
	}

	stts = "true"
	msg = "login successfuly "
	result := map[string]string{
		"status": stts,
		"msg":    msg,
		"token":  token,
	}

	ResponseSuccess(c, http.StatusOK, result)
}

func (uc *UserController) Register(c *gin.Context) {
	model := models.NewUser()
	err := c.ShouldBindJSON(model)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, "Data not valid")
		c.Abort()
		return
	}

	check := repositorys.ValidationEmail(uc.DB, c, model.Email)
	if !check {
		ResponseError(c, http.StatusBadRequest, "Email has been used")
		c.Abort()
		return
	}

	passEncrypt, err := utils.Encryptor([]byte(model.Password))
	if err != nil {
		log.Println("faild encrypt password")
		ResponseError(c, http.StatusInternalServerError, "Something wrong please try again")
		c.Abort()
		return
	}
	log.Println(passEncrypt)
	newmodel := models.User{
		Firstname:    model.Firstname,
		Lastname:     model.Lastname,
		Email:        model.Email,
		Password:     passEncrypt,
		ImageProfile: model.ImageProfile,
		CreateAt:     model.CreateAt,
		UpdateAt:     model.UpdateAt,
	}

	err = repositorys.AddUser(uc.DB, &newmodel)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, "Register faild please try again")
		c.Abort()
		return
	}

	stts = "true"
	msg = "Register successfuly"
	result := map[string]string{
		"msg":    msg,
		"status": stts,
	}

	ResponseSuccess(c, http.StatusOK, result)

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
