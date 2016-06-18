package main

import (
	"fmt"
	"github.com/matthew-andrews/semver-as-a-service/semver"
	"github.com/matthew-andrews/semver-as-a-service/sources"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "semver"
	app.Usage = "cli to find the latest version of any GitHub repository"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "source",
			Usage: "name of source to look package up in, only 'github' currently supported",
			Value: "github",
		},
		cli.StringFlag{
			Name:  "satisfies",
			Usage: "semver pattern or aliases (e.g. 'latest') to match versions against",
			Value: "latest",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "identifier of the package, for example spf13/hugo",
		},
	}
	app.Action = func(c *cli.Context) error {
		source, err := sources.New(c.String("source"))
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("%s", err), 1)
		}

		version, err := semver.Semver(source, c.String("id"), c.String("satisfies"))
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("%s", err), 1)
		}
		fmt.Println(version)
		return nil
	}
	app.Run(os.Args)
}
