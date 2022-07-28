package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/utkarsh-1905/go-chain/blockchain"
	"github.com/utkarsh-1905/go-chain/transactions"
)

func main() {
	app := &cli.App{
		Name:    "Go-Chain CLI",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start a new Blockchain with the Genesis Block",
				Action: func(c *cli.Context) error {
					blockchain.Genesis()
					fmt.Println("Genesis Block Created")
					return nil
				},
			},
			{
				Name:    "transaction",
				Aliases: []string{"tx"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "to",
						Value:    "",
						Usage:    "The receiver of the transaction",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "from",
						Value:    "",
						Usage:    "The sender of the transaction",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "value",
						Value:    0,
						Usage:    "The value of the transaction",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "data",
						Value: "",
						Usage: "The data of the transaction",
					},
				},
				Usage: "Create a new transaction",
				Action: func(c *cli.Context) error {
					transactions.CreateTransaction(c.String("to"), c.String("from"), c.Int("value"), c.String("data"))
					fmt.Println("Transaction Created")
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "Show the current Blockchain",
				Action: func(c *cli.Context) error {
					bc := blockchain.Blockchain
					fmt.Println(bc)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}