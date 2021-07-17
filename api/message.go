package main

import (
  "encoding/json"
  "bytes"
  "errors"
  "net/http"
  "os"
)

// Req body used here to send message method:
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
  ChatID int64 `json:"chat_id"`
  Text string `json:"text"`
}

// Req body used here to send reply message method:
// https://core.telegram.org/bots/api#sendmessage
type sendReplyReqBody struct {
  ChatID int64 `json:"chat_id"`
  Text string `json:"text"`
  MessageID int64 `json:"reply_to_message_id"`
}

// Req body used here to send photo method:
// https://core.telegram.org/bots/api#sendphoto
type sendPhotoReqBody struct {
  ChatID int64 `json:"chat_id"`
  Photo string `json:"photo"`
}

// Send a message <msg> to the chat with ID <chatID> 
func SendMessage(chatID int64, msg string) error {
  // Create the request body struct
  reqBody := &sendMessageReqBody {
    ChatID: chatID,
    Text:   msg,
  }

  // Create the JSON body from the struct
  reqBytes, err := json.Marshal(reqBody)
  if err != nil {
    return err
  }

  // Send a request with your token
  res, err := http.Post("https://api.telegram.org/" + os.Getenv("BOT_TOKEN") + "/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

  if err != nil {
    return err
  }
  if res.StatusCode != http.StatusOK {
    return errors.New("Unexpected status" + res.Status)
  }
  return nil
}

// Send a reply message <msg> to the chat with ID <chatID> 
func SendReply(chatID int64, messageID int64, msg string) error {
  // Create the request body struct
  reqBody := &sendReplyReqBody {
    ChatID: chatID,
    MessageID: messageID,
    Text:   msg,
  }

  // Create the JSON body from the struct
  reqBytes, err := json.Marshal(reqBody)
  if err != nil {
    return err
  }

  // Send a request with your token
  res, err := http.Post("https://api.telegram.org/" + os.Getenv("BOT_TOKEN") + "/sendMessage", "application/json", bytes.NewBuffer(reqBytes))

  if err != nil {
    return err
  }
  if res.StatusCode != http.StatusOK {
    return errors.New("Unexpected status" + res.Status)
  }
  return nil
}

// Send a photo <photo> to the chat with ID <chatID> 
func SendPhoto(chatID int64, photo string) error {
  // Create the request body struct
  reqBody := &sendPhotoReqBody {
    ChatID: chatID,
    Photo:   photo,
  }

  // Create the JSON body from the struct
  reqBytes, err := json.Marshal(reqBody)
  if err != nil {
    return err
  }

  // Send a request with your token
  res, err := http.Post("https://api.telegram.org/" + os.Getenv("BOT_TOKEN") + "/sendPhoto", "application/json", bytes.NewBuffer(reqBytes))

  if err != nil {
    return err
  }
  if res.StatusCode != http.StatusOK {
    return errors.New("Unexpected status" + res.Status)
  }
  return nil
}

