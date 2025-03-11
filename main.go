package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "install app",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						fmt.Println("nar install <app>")
						return nil
					} else {
						appName := cCtx.Args().First()
						downloadLink, binPath, symlink, appVersion, autoDownload := Fetch(appName)
						Download(downloadLink, binPath, symlink, autoDownload, appName, appVersion)
					}
					return nil
				},
			},
			{
				Name:    "uninstall",
				Aliases: []string{"ui", "remove", "r", "purge", "p"},
				Usage:   "uninstall app",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						fmt.Println("nar uninstall <app>")
						return nil
					} else {
						appName := cCtx.Args().First()
						deleteLink := FetchDel(appName)
						DeleteApp(deleteLink, appName)
					}
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"up", "upgrade"},
				Usage:   "uninstall app",
				Action: func(cCtx *cli.Context) error {
					appName := cCtx.Args().First()
					UpdateDB(appName, FetchCurrentVersionNo(appName))
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "show", "s"},
				Usage:   "list out all apps",
				Action: func(cCtx *cli.Context) error {
					List()
					return nil
				},
			},
			{
				Name:    "oof",
				Aliases: []string{"nil"},
				Usage:   "Updates app",
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
