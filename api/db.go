package main

import (
  "gopkg.in/mgo.v2"
  "os"
)

var Col *mgo.Collection

func CreateInstance(inst interface{}, collection string) bool {
  // Connect to mongo
  session, err := mgo.Dial("mongo" + GetDBPort())
  if err != nil {
    os.Exit(1)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  // Get collection
  Col = session.DB(GetDBName()).C(collection)
  // Create instance
  if Col.Insert(inst) == nil {
    return true
  }
  return false
}

func SearchInstance(id int64, collection string, inst interface{}) bool {
  // Connect to mongo
  session, err := mgo.Dial("mongo" + GetDBPort())
  if err != nil {
    os.Exit(1)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  // Get collection
  Col = session.DB(GetDBName()).C(collection)
  // Search instance
  if Col.FindId(id).One(inst) == nil {
    return true
  }
  return false
}

func UpdateInstance(id int64, inst interface{}, collection string) bool {
  // Connect to mongo
  session, err := mgo.Dial("mongo" + GetDBPort())
  if err != nil {
    os.Exit(1)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  // Get collection
  Col = session.DB(GetDBName()).C(collection)
  // Update instance
  if Col.UpdateId(id, inst) == nil {
    return true
  }
  return false
}
