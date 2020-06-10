package cmd

import (
	"fmt"
	"github.com/bst27/feed-webhook/config"
	"github.com/urfave/cli/v2"
	"log"
	"net/url"
)

func AddHook() *cli.Command {
	return &cli.Command{
		Name:        "add-webhook",
		Usage:       "add a webhook to a monitored feed",
		Description: "This command adds a webhook to an already monitored feed.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "feed",
				Usage:    "Define ID of the feed to add the webhook (use list command to get a list of defined feeds)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "type",
				Usage:    "Define webhook type (allowed: web, filesystem, stdout).",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "config",
				Usage:    "Define path to the config file to save settings. (e.g. settings.json)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "url",
				Usage: "Define target URL for web type webhooks",
			},
			&cli.StringFlag{
				Name:  "directory",
				Usage: "Define target directory for filesystem type webhooks",
			},
		},
		Action: func(c *cli.Context) error {
			feed := c.String("feed")

			webhookType := c.String("type")
			if webhookType != "web" && webhookType != "filesystem" && webhookType != "stdout" {
				fmt.Println("Invalid webhook type (allowed: web, filesystem, stdout)")
				return nil
			}

			configFile := c.String("config")

			conf, err := config.ReadOrInit(configFile)
			if err != nil {
				log.Fatal(err)
			}

			if !conf.HasFeed(feed) {
				fmt.Println("Invalid feed id (feed not found)")
				return nil
			}

			switch webhookType {
			case "web":
				webhookUrl, err := url.ParseRequestURI(c.String("url"))
				if err != nil {
					fmt.Println("Invalid webhook URL")
					return nil
				}

				conf.AddWebHook(feed, webhookUrl.String())
			case "filesystem":
				directory := c.String("directory")
				if directory == "" {
					fmt.Println("Invalid directory")
					return nil
				}

				conf.AddFilesystemHook(feed, directory)
			case "stdout":
				conf.AddStdoutHook(feed)
			default:
				panic("Cannot handle webhook type: " + webhookType)
			}

			err = config.Save(conf, configFile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Added webhook:", configFile)

			return nil
		},
	}
}
