package main

import (
  "github.com/melezhik/sparrowgo"
)

func main() {

  type Params struct {
    Message string
  }

  sparrowgo.RunTask("foo",Params{Message: "hello from main"})

  sparrowgo.SetStdout("hook: here we go\nnext line")

}
