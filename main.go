package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/fatih/color"
	"github.com/nomad-software/meme/cli"
	"github.com/nomad-software/meme/image"
	"github.com/nomad-software/meme/output"
)

func main() {
	opt := cli.ParseOptions()

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(output.Stderr, color.RedString(r.(string)))
			os.Exit(1)
		}
	}()
	if opt.Help {
		opt.PrintUsage()

	} else if opt.ListTemplates {
		for _, id := range cli.ImageIds {
			fmt.Fprintln(output.Stdout, color.CyanString("%s", id))
		}

	} else if opt.Valid() {
		st := image.Load(opt)
		st = image.RenderImage(opt, st)

		if opt.ClientID != "" {
			url := image.Upload(opt, st)
			output.Info(url)
		} else {
			file := image.Save(opt, st)
			output.Info(file)
		}
	}
}
