package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"github.com/utkarsh-1905/go-chain/blockchain"
	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
	"github.com/utkarsh-1905/go-chain/wallet"
)

func main() {
	app := &cli.App{
		Name:    "go-chain",
		Usage:   "A command line interface for Go-Chain",
		Version: "1.1.0",
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
					if c.String("to") == "" || c.String("from") == "" || c.Int("value") == 0 {
						return cli.NewExitError("Please provide all the required parameters", 1)
					}
					transactions.CreateTransaction(c.String("to"), c.String("from"), c.Int("value"), c.String("data"))
					return nil
				},
			},
			{
				Name:  "show",
				Usage: "Show the current Blockchain",
				Action: func(c *cli.Context) error {
					blockchain := helpers.ReadAndUnmarshallBlockchain()
					spew.Dump(blockchain)
					return nil
				},
			},
			{
				Name:  "block",
				Usage: "Create a new block",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "miner",
						Aliases:  []string{"m"},
						Usage:    "The miner of the block",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					if c.String("miner") == "" {
						fmt.Println("Please specify the miner of the block")
						return nil
					} else {
						blockchain.CreateBlock(c.String("miner"))
						fmt.Println("Block Created")
						return nil
					}
				},
			},
			{
				Name:  "wallet",
				Usage: "To create a new wallet on chain",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of wallet",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					wallet.GenerateWallet(c.String("name"))
					spew.Dump("Wallet created, check the " + c.String("name") + ".json file in the root folder")
					spew.Dump("Use the public key for transactions")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
