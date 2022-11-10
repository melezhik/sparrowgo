package sparrowgo

import (

  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "log"

)

var crd = os.Getenv("cache_root_dir")

type settings struct {
  debug bool
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

// Set task output data.

func UpdateState (data interface{}) {

  log_(">>> UpdateState | cache_root_dir: %s\n",crd)

  file, _ := json.MarshalIndent(data, "", " ")

  _ = ioutil.WriteFile(fmt.Sprintf("%s/state.json",crd), file, 0644)

  log_(">>> UpdateState | update %s/state.json\n",crd)

  return
}
