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
	var recursive bool

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Usage:       "recursive size of directories (default: false)",
				Destination: &recursive,
			},
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "human-readable sizes (auto-select unit) (default: false)",
				Destination: &humanize,
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "include hidden files and directories (default: false)",
				Destination: &includeHidden,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()

			if path == "" {
				return fmt.Errorf("path argument is required")
			}

			size, err := pathsize.GetSize(path, includeHidden, recursive)

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
