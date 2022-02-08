package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SetupPort() string {
	err := godotenv.Load() //loads .env file in to the project environment
	if err != nil {
		fmt.Println(err)
		fmt.Println("unable to load envoriment")
		return (":3000") //return default :3000 as port number
	}
	return os.Getenv("PORT")
}

func GetMongoDBURI() string {
	err := godotenv.Load() //loads .env file in to the project environment
	if err != nil {
		fmt.Println(err)
		fmt.Println("unable to load envoriment")
		return ("unable to load envoriment") //return default :3000 as port number
	}
	return os.Getenv("MONGODBURI")
}
