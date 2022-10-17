package main

import (
	"fmt"

	"github.com/ArchishmanSengupta/neo-cli/cmd"
	"github.com/ArchishmanSengupta/neo-cli/config"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	config.DbConn, _ = config.Connect()
	cmd.Execute()
}
