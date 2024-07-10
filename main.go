package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"

	"github.com/dunglas/frankenphp"
	opts "github.com/urfave/cli/v2"
)

const VERSION = "v0.0.1"

var app = &opts.App{
	Name: "Whoa",
	Description: "Whoa is a modern PHP runtime and production webserver",
	HelpName: "whoa",
	Version: VERSION,
	Copyright: "© 2024 Tristan Isham",
	Suggest: true,
	Before: func(ctx *opts.Context) error {
		if _, ok := os.LookupEnv("WHOA_DEBUG"); ok {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	},
	Commands: []*opts.Command{
		{
			Name: "run",
			Usage: "Run a PHP script from the command line",
			Aliases: []string{"r"},
			Description: "Run a PHP script as a PHP command line script",
			Args: true,
			ArgsUsage: " <script.php>",
			Action: func(ctx *opts.Context) error {
				script := ctx.Args().First()
				os.Exit(frankenphp.ExecuteScriptCLI(script, ctx.Args().Slice()))
				return nil
			},
		},
		{
			Name: "fs",
			Usage: "Start a PHP file-based server",
			Description: "Start a traditional PHP server with file-based routing",
			Args: true,
			ArgsUsage: " </path/to/dir>",
			Flags: []opts.Flag{
				&opts.BoolFlag{
					Name: "sym",
					Usage: "Enable symblinks (may generate errors due to caching if symlinks are changed and Whoa isn't restarted)",
				},
			},
			Action: func(ctx *opts.Context) error {
				dir := ctx.Args().First()

				if err := frankenphp.Init(); err != nil {
					panic(err)
				}
				defer frankenphp.Shutdown()

				target := dir
				if newTarget, err := filepath.Abs(dir); err == nil {
					target = newTarget
				}
			
				http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					req, err := frankenphp.NewRequestWithContext(r, frankenphp.WithRequestDocumentRoot(target, ctx.Bool("sym")))
					if err != nil {
						panic(err)
					}
			
					if err := frankenphp.ServeHTTP(w, req); err != nil {
						panic(err)
					}
				})
				if err := http.ListenAndServe(":8080", nil); err != nil {
					return err
				}

				return nil
			},
		},
	},
}

func main() {
	
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	
}
