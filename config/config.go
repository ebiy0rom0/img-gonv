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
}
