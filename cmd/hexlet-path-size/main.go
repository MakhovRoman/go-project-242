package main

import (
	pathsize "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()

			if path == "" {
				return fmt.Errorf("path argument is required")
			}

			size, err := pathsize.GetSize(path)

			if err != nil {
				return err
			}

			pathsize.PrintSize(size, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
