package sparrowgo

import (

  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "log"

)

var crd = os.Getenv("cache_root_dir")
var cd = os.Getenv("cache_dir")

type settings struct {
  debug bool
}

type task_state struct {
  State interface{}  `json:"state"`
}

var settings_ = settings {debug: false}

func DebugOn () {
  settings_.debug = true
}

func DebugOff () {
  settings_.debug = false
}

func log_ (pattern string, line string) {

  if settings_.debug ==true {
    fmt.Printf(pattern,line)
  }
}

func Os () string {

  return os.Getenv("os")
}

// Return task config as struct.

func Config (data interface{}) {

  log_(">>> Config | cache_root_dir: %s\n",crd)

  jsonFile, err := os.Open(fmt.Sprintf("%s/config.json",crd))

  if err != nil {
    log.Fatal(err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  log_(">>> Config | json: %s\n",string(byteValue))

  if err := json.Unmarshal(byteValue, &data); err != nil {
    panic(err)
  }

}

// Return task variables as struct.

func TaskVars (data interface{}) {

  log_(">>> TaskVars | cache_dir: %s\n",cd)

  jsonFile, err := os.Open(fmt.Sprintf("%s/variables.json",cd))

  if err != nil {
    log.Fatal(err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  log_(">>> TaskVars | json: %s\n",string(byteValue))

  if err := json.Unmarshal(byteValue, &data); err != nil {
    panic(err)
  }

}
// Set task output data.

func UpdateState (data interface{}) {

  log_(">>> UpdateState | cache_root_dir: %s\n",crd)

  js, err := json.Marshal(data)

  if err != nil {
    panic(err)
  }

  log_(">>> UpdateState | data: %s\n",string(js))

  _ = ioutil.WriteFile(fmt.Sprintf("%s/state.json",crd), js, 0644)

  log_(">>> UpdateState | update %s/state.json\n",crd)

  return
}

// Impliment run_task

func RunTask (path string, params interface{}) {

  fmt.Println("task_var_json_begin")

  if params == nil {
    fmt.Println("{ }")
  } else {
    jsonData, err := json.Marshal(params)
    if err != nil {
      panic(err)
    }
    fmt.Printf("%s\n",jsonData)
  }

  fmt.Println("task_var_json_end")

  fmt.Printf("task: %s\n", path)

}

