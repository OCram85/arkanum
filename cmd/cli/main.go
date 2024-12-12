package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)


func main()  {
	cmd := &cli.Command{
		Name: "",
	}

	if err := cmd.Run(context.Background(), os.Args); err!= nil {
		log.Fatal(err)
	}
}
