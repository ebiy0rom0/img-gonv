package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envfile = "./gonv.cnf"

var (
	OutputPath string
)

func init() {
	if err := godotenv.Load(envfile); err != nil {
		log.Fatal(err)
	}

	OutputPath = os.Getenv("OUTPUT_PATH")
	if _, err := os.Stat(OutputPath); os.IsNotExist(err) {
		if err := os.Mkdir(OutputPath, 0766); err != nil {
			log.Fatal(err)
		}
	}
}
