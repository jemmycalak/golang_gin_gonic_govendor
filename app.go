package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/go_gin_govendor/src/config"
	"github.com/jemmycalak/go_gin_govendor/src/controllers/user"
	"github.com/jemmycalak/go_gin_govendor/src/routers"
)

func main() {

	db, err := config.ConnectToDb()

	if err != nil {
		log.Fatal("error", err)
	}

	ruser := user.NewUserContrroller(db)
	r := gin.Default()
	// r.Use(routers.AuhtToken())

	routers.UserRouters(r, ruser)
	fmt.Println("app was running")
	r.Run(":8000")
}
