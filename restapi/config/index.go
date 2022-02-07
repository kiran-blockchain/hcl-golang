package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SetupPort()string {
	err:= godotenv.Load()
	 if err!=nil{
		 fmt.Println(err)
		 fmt.Println("unable to load envoriment")
		 return (":3000") //return default :3000 as port number
	 }
	 return os.Getenv("PORT")
}