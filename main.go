package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

const VERSION = "0.0.1"

var createKeyCommand = cli.Command{
	Name:  "createkey",
	Usage: "create a pair of public/private key",
	Flags: []cli.Flag{
		cli.IntFlag{Name: "size, s", Value: 2048, Usage: "key size in bits"},
		cli.StringFlag{Name: "password, p", Value: "", Usage: "password to access the private key"},
	},
	Action: func(c *cli.Context) {
		keyGenerator := new(KeyGenerator)
		fmt.Println("creating key")
		if err := keyGenerator.CreateNewKey(c.Int("size")); err != nil {
			fmt.Println("error while creating the key\n", err)
		}
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "ecrypto"
	app.Usage = "encrypt files"
	app.Commands = []cli.Command{
		createKeyCommand,
	}

	app.Run(os.Args)
}