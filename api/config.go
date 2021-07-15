package main

import (
  "encoding/json"
  "os"
  "io/ioutil"
)

// You must have the file config.json, see config.example.json to see how this file must be
const CONFIG_PATH = "config.json"

type C struct {
  Token string `json:"token"`
  Port string `json:"port"`
  DBName string `json:"db_name"`
  DBPort string `json:"db_port"`
}

// Read json file
func ReadConfig(config *C) {
  jsonFile, _ := os.Open(CONFIG_PATH)
  defer jsonFile.Close()
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &config)
}

func GetPort() string {
  c := &C{}
  ReadConfig(c)
  return c.Port
}

func GetToken() string {
  c := &C{}
  ReadConfig(c)
  return c.Token
}

func GetDBName() string {
  c := &C{}
  ReadConfig(c)
  return c.DBName
}

func GetDBPort() string {
  c := &C{}
  ReadConfig(c)
  return c.DBPort
}
