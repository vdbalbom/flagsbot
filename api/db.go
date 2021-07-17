package main

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"
  "os"
  "fmt"
  "context"
  "time"
)

func CreateInstance(inst interface{}, collection string) bool {
  // Link
  link := "mongodb://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@mongo:27017/" + os.Getenv("DB_NAME") + "?authSource=" + os.Getenv("DB_NAME")
  fmt.Println(link)

  // Connect to mongo
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(link))
  defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
  }()
  if err != nil {
    fmt.Println(err)
    return false
  }
  coll := client.Database(os.Getenv("DB_NAME")).Collection(collection)

  // Create instance
  _, err = coll.InsertOne(context.TODO(), inst)
  if err != nil {
    fmt.Println(err)
    return false
  }
  return true
}

func SearchInstance(id int64, collection string, inst interface{}) bool {
  // Link
  link := "mongodb://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@mongo:27017/" + os.Getenv("DB_NAME") + "?authSource=" + os.Getenv("DB_NAME")
  fmt.Println(link)

  // Connect to mongo
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(link))
  defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
  }()
  if err != nil {
    fmt.Println(err)
    return false
  }
  coll := client.Database(os.Getenv("DB_NAME")).Collection(collection)

  // Search instance
  filter := bson.D{{"_id", id}}
  err = coll.FindOne(context.TODO(), filter).Decode(inst)
  if err != nil {
    fmt.Println(err)
    return false
  }
  return true
}

func UpdateInstance(id int64, inst interface{}, collection string) bool {
  // Link
  link := "mongodb://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@mongo:27017/" + os.Getenv("DB_NAME") + "?authSource=" + os.Getenv("DB_NAME")
  fmt.Println(link)

  // Connect to mongo
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(link))
  defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
  }()
  if err != nil {
    fmt.Println(err)
    return false
  }
  coll := client.Database(os.Getenv("DB_NAME")).Collection(collection)

  // Update instance
  filter := bson.D{{"_id", id}}
  opts := options.Replace().SetUpsert(true)
  _, err = coll.ReplaceOne(context.TODO(), filter, inst, opts)
  if err != nil {
    fmt.Println(err)
    return false
  }
  return true
}
