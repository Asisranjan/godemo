package main
import (
        "github.com/urfave/cli"
        "./plugins"
)


func NewCli() *cli.App {
  app := cli.NewApp()
  app.Name = "Simple pizza CLI"
  app.Usage = "An example CLI for ordering pizzas"
  app.Author = "Jeroenouw"
  app.Version = "1.0.0"
  allCmds := []cli.Command{}
  allCmds = append(allCmds,plugins.ApplicationCommands...)
  app.Commands = allCmds
  return app
} 
