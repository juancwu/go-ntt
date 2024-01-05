package main

import (
	charmlog "github.com/charmbracelet/log"
	dotenv "github.com/joho/godotenv"

	"github.com/juancwu/go-ntt/cmd"
	"github.com/juancwu/go-ntt/config"
	"github.com/juancwu/go-ntt/util"
)

func main() {
	// load env
	if err := dotenv.Load(); err != nil {
		charmlog.Fatal(err)
	}

	// setup logger so that the rest of the application can log things in a formatted manner
	util.InitLog()

	// setup typed safe environment variables
	if err := config.InitEnv(); err != nil {
		util.Log().Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		util.Log().Fatal(err)
	}
}
