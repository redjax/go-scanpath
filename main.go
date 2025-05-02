package main

import (
	"log"

	"github.com/alecthomas/kong"
)

type CLI struct {
	ScanPath string `short:"s" long:"scan-path" help:"Directory path to scan." default:"." type:"existingdir"`
	Limit    int    `short:"l" long:"limit-results" help:"Limit the number of results (0 = unlimited)." default:"0"`
}

func main() {
	var cli CLI
	kong.Parse(&cli,
		kong.Name("dirscan"),
		kong.Description("Scan a directory and list items with metadata."),
	)

	err := scanDirectory(cli.ScanPath, cli.Limit)
	if err != nil {
		log.Fatal(err)
	}
}
