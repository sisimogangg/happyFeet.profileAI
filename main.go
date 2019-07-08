package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	_kidsRepo "github.com/sisimogangg/happyFeet.profileAI/kids/repository"
	_profileController "github.com/sisimogangg/happyFeet.profileAI/profile/controller"
	_profileRepo "github.com/sisimogangg/happyFeet.profileAI/profile/repository"
	_profileservice "github.com/sisimogangg/happyFeet.profileAI/profile/service"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func buildMySQLConnString() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(connectionStr)
	return connectionStr
}

func main() {

	dbConn, err := sql.Open(`mysql`, buildMySQLConnString())
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := mux.NewRouter()

	profileRepo := _profileRepo.NewMysqlProfileRepository(dbConn)
	kidsRepo := _kidsRepo.NewMySQLKidsRepository(dbConn)

	timeContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	profileService := _profileservice.NewProfileService(profileRepo, kidsRepo, timeContext)

	_profileController.NewProfileHandler(router, profileService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
