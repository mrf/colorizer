package main

import (
  "os"
  "fmt"
  "github.com/codegangsta/cli"
)

func main() {
  // Map out basic bash colors
  colors := make(map[string]string)
  colors["black"] = "\033[0;30m"
  colors["red"] = "\033[0;31m"
  colors["green"] = "\033[0;32m"
  colors["yellow"] = "\033[0;33m"
  colors["blue"] = "\033[0;34m"

  app := cli.NewApp()
  app.Name = "colorizer"
  app.Usage = "apply color formatting to command line output"
  for index,element := range colors {
    app.Commands = []cli.Command{
      {
        Name: index,
        Usage: "colorize %sgreen",
        Action: func(c *cli.Context) {
          fmt.Println("colors: ",colors)
          fmt.Println(element, c.Args().First())
        },
      },
    }
  }
  app.Run(os.Args)
}
