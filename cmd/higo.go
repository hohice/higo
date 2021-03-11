package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hohice/higo/internal/pkg/project"
	herrors "github.com/hohice/higo/pkg/errors"
	"github.com/hohice/higo/pkg/version"
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		{
			Name:  "new",
			Usage: "create new project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name, n",
					Usage: "use name",
				},
			},
			Action: newCommand,
		},
		{
			Name:  "version",
			Usage: "print higo version",
			Action: func(c *cli.Context) error {
				fmt.Println("version:", version.Version())
				fmt.Println("commit:", version.Commit())
				fmt.Println("buildTs:", version.BuildTs())
				return nil
			},
		},
	}
)

func newCommand(c *cli.Context) error {
	pname := c.String("n")
	if len(pname) == 0 {
		return errors.New("missing -n")
	}

	return project.NewProject(pname)
}

func main() {
	app := cli.NewApp()
	app.Usage = "go project initializater"
	app.Version = version.Version()
	app.Commands = commands

	defer herrors.PanicHandler()

	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
