package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  // sparrowgo.DebugOn()

  // var state interface {}

  state := make(map[string](map[string]interface{}))

  sparrowgo.GetState(&state)

  fmt.Printf("Task state: %s\n",state["Message"]["Hello"])

}
