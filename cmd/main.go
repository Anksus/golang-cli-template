package main

import (
	"os"

	"github.com/Anksus/golang-cli-template.git/internal/service"
	"github.com/urfave/cli/v2"
)

func main(){
	cliApp := cli.NewApp()


	cliApp.Name = "Sample app"
	cliApp.Version = "1.0.0"
	cliApp.Usage = "Normal and simple CLI app"

	cliApp.Commands = cli.Commands{
		{
			Name: "server",
			Usage: "Starts our server",
			Action: func(c *cli.Context)error{
				return service.PrintNum()
				
			},
			
		},

	}
	if err := cliApp.Run(os.Args); err!=nil{
		panic(err)
	}
}