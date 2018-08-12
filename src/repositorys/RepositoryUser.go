package repositorys

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jemmycalak/go_gin_govendor/src/models"
)

func ShowUsers(db *sql.DB) (models.Users, error) {
	querys := `SELECT * FROM "t_user" `

	var musers models.Users

	rows, err := db.Query(querys)
	if err != nil {
		fmt.Printf("Query error")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var nuser models.User

		err := rows.Scan(&nuser.Id, &nuser.Firstname, &nuser.Lastname, &nuser.Email, &nuser.Password, &nuser.ImageProfile, &nuser.CreateAt, &nuser.UpdateAt)
		if err != nil {
			fmt.Println("error loop data")
			return nil, err
		}

		musers = append(musers, nuser)
	}

	return musers, nil
}

func AddUser(db *sql.DB, model *models.User) error {
	querys := `INSERT INTO "t_user" 
	(
		"firtsname", "lastname", "email", "password", "imageprofile", "createat", "updateat"
	) 
	VALUES 
	(
		$1, $2, $3, $4, $5, $6, $7
	)`

	statement, err := db.Prepare(querys)
	if err != nil {
		fmt.Println("Error Prepare Query")
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(model.Firstname, model.Lastname, model.Email, model.Password, model.ImageProfile, model.CreateAt, model.UpdateAt)

	log.Println(model.Firstname, model.Lastname)
	if err != nil {
		fmt.Println("Error Execute Query")
		return err
	}

	return nil
}

func UpdateUser(db *sql.DB, iduser string, model *models.User) error {
	query := `UPDATE "t_user" SET "firtsname"=$1, "lastname"=$2, "email"=$3, "password"=$4, "imageprofile"=$5, "updateat"=$6 WHERE "iduser"=$7`

	statment, err := db.Prepare(query)

	if err != nil {
		fmt.Println("Error prepare update")
		return err
	}

	defer statment.Close()

	_, err = statment.Exec(model.Firstname, model.Lastname, model.Email, model.Password, model.ImageProfile, model.UpdateAt, iduser)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, iduser string) error {
	query := `DELETE FROM "t_user" WHERE "iduser"=$1`

	statment, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Error prepare delete")
		return err
	}

	defer statment.Close()
	_, err = statment.Exec(iduser)

	if err != nil {
		fmt.Println("Error Exc query")
		return err
	}

	return nil

}

func FindUserById(db *sql.DB, iduser int) (*models.User, error) {
	query := `SELECT "iduser", "firtsname", "lastname", "email", "imageprofile", "createat", "updateat" from "t_user" where "iduser" = $1`
	var newmodel models.User

	err := db.QueryRow(query, iduser).Scan(&newmodel.Id, &newmodel.Firstname, &newmodel.Lastname, &newmodel.Email, &newmodel.ImageProfile, &newmodel.CreateAt, &newmodel.UpdateAt)
	if err != nil {
		fmt.Println("error query select")
		return nil, err
	}

	//cara kedua
	// query := `SELECT * FROM "t_user" WHERE "iduser" = $1`
	// niduser := strconv.Itoa(iduser) //convert to string
	// log.Println(niduser)
	// statment, err := db.Prepare(query)
	// if err != nil {
	// 	fmt.Println("error prepare query select")
	// 	return nil, err
	// }
	// defer statment.Close()
	// err = statment.QueryRow(iduser).Scan(&newmodel.Id, &newmodel.Firstname, &newmodel.Lastname, &newmodel.Email, &newmodel.Password, &newmodel.ImageProfile, &newmodel.CreateAt, &newmodel.UpdateAt)
	// if err != nil {
	// 	fmt.Println("error QueryRow query select")
	// 	return nil, err
	// }

	return &newmodel, nil
}

func LoginRepository(db *sql.DB, model *models.LoginStruct) (*models.User, error) {
	query := `SELECT "iduser", "password" FROM "t_user" WHERE "email" = $1`

	var newmodel models.User
	log.Println(model.Email, model.Password)

	err := db.QueryRow(query, model.Email).Scan(&newmodel.Id, &newmodel.Password)
	if err != nil {
		fmt.Println("User not found")
		return nil, err
	}

	return &newmodel, nil
}

func ValidationEmail(db *sql.DB, c *gin.Context, email string) (stts bool) {
	query := `select email from t_user where email = $1`
	var model models.User
	err := db.QueryRow(query, email).Scan(&model.Email)
	if err != nil {
		return true
	}
	return false
}
