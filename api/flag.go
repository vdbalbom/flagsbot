package main

import (
  "encoding/json"
  "os"
  "io/ioutil"
  "math/rand"
  "strconv"
)

type Flag struct {
  Link string `json:"link"`
  Names []string `json:"names"`
}

// Read json file
func DrawFlag() Flag {
  flag := Flag{}
  i := rand.Intn(6)+1
  jsonFile, _ := os.Open("flags/" + strconv.Itoa(i) + ".json")
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &flag)
  return flag
}

