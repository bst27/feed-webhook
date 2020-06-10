package cmd

import (
	"fmt"
	"github.com/bst27/feed-webhook/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func List() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Usage:       "list feeds and webhooks defined in config",
		Description: "This command lists feeds and webhooks which have been defined in the config.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Usage:    "Define path to the config file to read settings. (e.g. settings.json)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "type",
				Usage:    "Define type to be listed (allowed: feed, webhook)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			conf, err := config.Read(c.String("config"))
			if err != nil {
				log.Fatal(err)
			}

			listedType := c.String("type")
			if listedType != "feed" && listedType != "webhook" {
				fmt.Println("Invalid type (allowed: feed, webhook)")
				return nil
			}

			tw := tabwriter.NewWriter(os.Stdout, 10, 3, 5, ' ', 0)

			switch listedType {
			case "feed":
				fmt.Fprintln(tw, strings.Join([]string{"Feed ID:", "URL:", "Polling Interval", "Type:"}, "\t"))
				for _, f := range conf.Feeds {
					fmt.Fprintln(tw, strings.Join([]string{f.ID, f.URL, strconv.FormatInt(f.PollingInterval, 10), f.Type}, "\t"))
				}
			case "webhook":
				fmt.Fprintln(tw, strings.Join([]string{"Webhook ID:", "Feed ID:", "URL:", "Type:"}, "\t"))
				for _, w := range conf.Webhooks {
					fmt.Fprintln(tw, strings.Join([]string{w.ID, w.FeedID, w.URL, w.Type}, "\t"))
				}
			}

			tw.Flush()

			return nil
		},
	}
}
