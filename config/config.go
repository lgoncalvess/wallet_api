package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDbConnection string = ""
	DB                 string = ""
	Port               int
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	StringDbConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("DB"))

	DB = os.Getenv("DB")

	Port, erro = strconv.Atoi(os.Getenv("PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
}
