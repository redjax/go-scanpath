package main

import (
	"log"
	"scanpath/internal/scan"

	"github.com/alecthomas/kong"
)

type CLI struct {
	ScanPath string `short:"p" long:"path" help:"Directory path to scan." default:"." type:"existingdir"`
	Limit    int    `short:"l" long:"limit" help:"Limit the number of results (0 = unlimited)." default:"0"`
	Sort     string `short:"s" long:"sort" help:"Column to sort by (name, size, created, modified, owner, permissions)" default:"name"`
	Order    string `short:"o" long:"order" help:"Sort order: asc or desc" default:"asc"`
	Filter   string `short:"f" long:"filter" help:"Filter results, e.g. --filter 'size <10MB' or --filter 'created >2022-01-01' or --filter 'name ~ someNamePart"`
}

func main() {
	var cli CLI
	kong.Parse(&cli,
		kong.Name("scanpath"),
		kong.Description("Scan a directory and list items with metadata."),
	)

	err := scan.ScanDirectory(cli.ScanPath, cli.Limit, cli.Sort, cli.Order, cli.Filter)
	if err != nil {
		log.Fatal(err)
	}
}
