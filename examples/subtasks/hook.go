package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  // sparrowgo.DebugOn()

  type Params struct {
    Message string
  }

  p := Params{Message: "hello from main"}

  fmt.Println(p.Message)

  sparrowgo.RunTask("foo",&p)

}
