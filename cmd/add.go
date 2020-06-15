package cmd

import (
	"fmt"
	"github.com/bst27/feed-webhook/config"
	"github.com/urfave/cli/v2"
	"log"
	"net/url"
	"time"
)

func Add() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Usage:       "add a feed to be monitored",
		Description: "This command adds a feed defined as an URL to the list of monitored feeds.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Usage:    "Define feed URL to be monitored (e.g. https://www.digitalpush.net/feed.xml).",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "type",
				Usage:    "Define feed type (allowed: atom, rss).",
				Required: true,
			},
			&cli.Uint64Flag{
				Name:        "interval",
				Usage:       "Define feed polling interval in seconds.",
				Value:       3600,
				DefaultText: "3600",
			},
			&cli.StringFlag{
				Name:     "config",
				Usage:    "Define path to the config file to save feed settings. (e.g. settings.json)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			feedUrl, err := url.ParseRequestURI(c.String("url"))
			if err != nil {
				fmt.Println("Invalid feed URL")
				return nil
			}

			pollingInterval, err := time.ParseDuration(fmt.Sprintf("%ds", c.Uint64("interval")))
			if err != nil {
				fmt.Println("Invalid polling interval")
				return nil
			}

			if pollingInterval < 1 {
				fmt.Println("Invalid polling interval (< 1)")
				return nil
			}

			feedType := c.String("type")
			if feedType != "atom" && feedType != "rss" {
				fmt.Println("Invalid feed type (allowed: atom, rss)")
				return nil
			}

			configFile := c.String("config")

			conf, err := config.ReadOrInit(configFile)
			if err != nil {
				log.Fatal(err)
			}

			conf.AddFeed(*feedUrl, feedType, pollingInterval)
			err = config.Save(conf, configFile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Added feed monitoring:", feedUrl, pollingInterval, configFile)

			return nil
		},
	}
}
