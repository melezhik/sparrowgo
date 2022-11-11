# sparrowgo

Write SparrowCI tasks on Golang

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

  sparrowgo.DebugOn()

  type Params struct {
    Message string
  }

  type Message struct {
    Message string
  }

  var params Params

  sparrowgo.Config(&params)

  fmt.Printf("Sparrow says: %s\n", params.Message)

  switch sparrowgo.Os() {
    case "darwin":
      fmt.Println("hello Mac")
    case "arch":
      fmt.Println("hello Arch Linux")
    case "debian":
      fmt.Println("hello Debian")
    // so on
  }

  sparrowgo.UpdateState(&Message{Message : "Hello from Go"})

}
```

## Create Sparrow wrapper

`test.raku`:

```raku
use Sparrow6::DSL;

my $s = task-run ".", %(
  message => "Hello from Raku"
);

say "message: ", $s<state><Message>;
```

Run:

```bash
raku test.raku
```

Output will be:

```
[task run: task.bash - .]
[task stdout]
13:08:47 :: Sparrow says: Hello from Raku
message: Hello from Go
```

# Other options


## Enable debug

```go
sparrowgo.DebugOn()
```

## Disable debug

```go
sparrowgo.DebugOff()
```

## Get know OS name

```go
  switch sparrowgo.Os() {
    case "darwin":
      fmt.Println("hello Mac")
    case "arch":
      fmt.Println("hello Arch Linux")
    case "debian":
      fmt.Println("hello Debian")
    // so on
  }
```

To get the list of supported OS follow this [link](https://github.com/melezhik/Sparrow6/blob/master/documentation/development.md#recognizable-os-list)

# Advanced topics

## Hooks and subtasks


Create hook task, `hook.go`

```go
package main

import (
  "github.com/melezhik/sparrowgo"
)

func main() {

  type Params struct {
    Message string
  }

  sparrowgo.RunTask("foo",Params{Message: "hello from main"})

}
```

Create subtask:

```bash
mkdir -p tasks/foo
```

`tasks/foo/task.go`

```go
package main

import (
  "fmt"
  "github.com/melezhik/sparrowgo"
)

func main() {

  type Vars struct {
    Message string
  }

  type Message struct {
    Message string
  }

  var task_vars Vars

  sparrowgo.TaskVars(&task_vars)

  fmt.Printf("foo subtask get this: %s\n",task_vars.Message)

  sparrowgo.UpdateState(Message{Message: "Hello from subtask"})
}
```

Create Sparrow wrapper:

`test.raku`:

```raku
use Sparrow6::DSL;

my $s = task-run ".";

say $s<state><Message>;
```

Run:

```bash
raku test.raku
```

Output will be:

```
[task run: task - .]
[task stdout]
12:59:14 :: foo subtask get this: hello from main
Hello from subtask
```
