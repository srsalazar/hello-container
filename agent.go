package main

import (
  "fmt"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
  "golang.org/x/net/context"
)

func main() {
  // Generate a new client
  cli, err := client.NewEnvClient()
    if err != nil {
      panic(err)
    }

  // Get all running containers
  containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
  if err != nil {
    panic(err)
  }

  //Iterate over all container and print running processes
  for _, container := range containers {
    fmt.Println(container.ID)
    resp, err :=cli.ContainerTop(context.Background(),container.ID, []string{})
    if err != nil {
      panic(err)
    }
    for _, title := range resp.Titles{
      fmt.Printf("%s ", title)
    }
    fmt.Printf("\n")
    for _, process := range resp.Processes{
      for _, column := range process {
        fmt.Printf("%s ",column)
      }
      fmt.Printf("\n")
    }
  }
}
