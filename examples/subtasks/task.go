package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  // sparrowgo.DebugOn()

  type State struct { Message string }

  var state State

  sparrowgo.GetState(&state)

  fmt.Printf("task state: %s\n",state.Message)

}
