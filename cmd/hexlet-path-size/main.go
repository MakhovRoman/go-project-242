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
	var humanize bool
	var includeHidden bool

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "Add size output in human-readable format",
				Destination: &humanize,
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "include hidden files and directories",
				Destination: &includeHidden,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()

			if path == "" {
				return fmt.Errorf("path argument is required")
			}

			size, err := pathsize.GetSize(path, includeHidden)

			if err != nil {
				return err
			}

			pathsize.PrintSize(size, path, humanize)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
