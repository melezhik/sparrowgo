# sparrowgo

Go interface for Sparrow tasks

# How to

## Create golang task

`task.go`:

```go
package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

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
```

## Create task binary

```bash
go build .
```

## Create Sparrow wrapper

`test.raku`:

```raku
use Sparrow6::DSL;

my $s = task-run ".", %(
  message => "Hello"
);

say "message: ", $s<state><Message>;
```

Output will be:

```
[task run: task.bash - .]
[task stdout]
13:08:47 :: CACHE_ROOT_DIR: /Users/melezhik/.sparrow6/tmp/7307517
13:08:47 :: json: { "message" : "Hello" }
13:08:47 :: 
13:08:47 :: Sparrow says: Hello
13:08:47 :: update /Users/melezhik/.sparrow6/tmp/7307517/state.json
message: Hello from Go
```
