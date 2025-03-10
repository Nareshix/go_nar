package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "install app",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						fmt.Println("nar install <app>")
						return nil
					}else{
						downloadLink, binPath, symlink,deleteCmd,autoDownload :=  Fetch(cCtx.Args().First())
						Download(downloadLink, binPath, symlink,deleteCmd,autoDownload)
					}
					return nil
				},
			},
			{
				Name:    "uninstall",
				Aliases: []string{"u", "remove", "r", "purge", "p"},
				Usage:   "uninstall app",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
