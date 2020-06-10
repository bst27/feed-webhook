package cmd

import (
	"fmt"
	"github.com/bst27/feed-webhook/config"
	"github.com/bst27/feed-webhook/monitoring"
	"github.com/bst27/feed-webhook/reader"
	"github.com/bst27/feed-webhook/webhook"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/context"
	"log"
	"time"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:        "run",
		Usage:       "run application and monitor feeds",
		Description: "This command monitors added feeds and sends webhooks on feed updates.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Usage:    "Define path to the config file to read settings. (e.g. settings.json)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			conf, err := config.Read(c.String("config"))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(fmt.Sprintf("Monitoring %d feeds ...", len(conf.Feeds)))

			m := monitoring.Create()
			for _, f := range conf.Feeds {

				switch f.Type {
				case "atom":
					webhooks := make([]webhook.Receiver, 0)

					for _, w := range conf.GetWebhooks(f.ID) {
						switch w.Type {
						case "web":
							webhooks = append(webhooks, webhook.Web(w.URL))
						case "filesystem":
							webhooks = append(webhooks, webhook.File(w.URL))
						case "stdout":
							webhooks = append(webhooks, webhook.Stdout())
						default:
							panic("Cannot handle hook type: " + w.Type)
						}
					}

					m.AddFeed(f.URL, f.PollingInterval, reader.GetAtom(), webhooks)
				default:
					panic("No reader for feed type: " + f.Type)
				}
			}

			// TODO: Wait for signal to shutdown
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			m.Start(ctx)

			time.Sleep(3600 * time.Second) // TODO: Remove

			return nil
		},
	}
}
