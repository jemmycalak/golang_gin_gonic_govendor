package repositorys

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jemmycalak/go_gin_govendor/src/models"
)

func ShowUsers(db *sql.DB) (models.Users, error) {
	querys := `SELECT * FROM "t_user"`

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
