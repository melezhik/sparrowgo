package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  // sparrowgo.DebugOn()

  type Vars struct {
    Message string
  }

  type Hello struct {
    Hello string
  }

  type Message struct {
    Message Hello
  }

  var task_vars Vars

  sparrowgo.TaskVars(&task_vars)

  fmt.Printf("foo subtask get this: %s\n",task_vars.Message)

  sparrowgo.UpdateState(Message{
     Message: Hello{
        Hello: "Hello from subtask",
      },
  })
}
