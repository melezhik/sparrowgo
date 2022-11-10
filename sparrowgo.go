package sparrowgo

import (

  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "log"

)

var crd = os.Getenv("cache_root_dir")

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func Config (data interface{}) {

  fmt.Printf("CACHE_ROOT_DIR: %s\n",crd)

  jsonFile, err := os.Open(fmt.Sprintf("%s/config.json",crd))
  if err != nil {
    log.Fatal(err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  fmt.Printf("json: %s\n",byteValue)

  if err := json.Unmarshal(byteValue, &data); err != nil {
    panic(err)
  }

}

func UpdateState (data interface{}) {

  file, _ := json.MarshalIndent(data, "", " ")

  _ = ioutil.WriteFile(fmt.Sprintf("%s/state.json",crd), file, 0644)

  fmt.Printf("update %s/state.json\n",crd)

  return 
}
