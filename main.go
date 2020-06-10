// Deliver webhook notifications for RSS/Atom feeds.
package main

import (
	"github.com/bst27/feed-webhook/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func bootstrap() *cli.App {
	return &cli.App{
		Name:    "feed-webhook",
		Usage:   "Monitor RSS/Atom feeds and deliver webhook notifications when feeds are updated.",
		Version: "1.0.0",
		Commands: []*cli.Command{
			cmd.Add(),
			cmd.AddHook(),
			cmd.Run(),
			cmd.List(),
		},
	}
}

// TODO: Handle application shutdown

func main() {
	app := bootstrap()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
