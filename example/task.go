package main

import (
	"fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  sparrowgo.DebugOn()

  type Params struct {
    Message string
  }

  type Message struct {
    Message string
  }

  type Output struct {
    State Message `json:"state"`
  }

  var params Params

  sparrowgo.Config(&params)

  fmt.Printf("Sparrow says: %s\n", params.Message)

  data := Output{ 
    State: Message{ 
      Message : "Hello from Go" ,
    },
  }

  sparrowgo.UpdateState(data)

}
