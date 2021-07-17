package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

// Incoming update when telegram sends a webhook event:
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
  Message struct {
    Text string `json:"text"`
    MessageID int64 `json:"message_id"`
    User struct {
      ID int `json:"id"`
      Username string `json:"username"`
    } `json:"from"`
    Chat struct {
      ID int64 `json:"id"`
      Type string `json:"type"`
    } `json:"chat"`
  } `json:"message"`
}

// Chat struct 
// https://core.telegram.org/bots/api#chat
// Used here: get chat method result
type Chat struct {
  ID int64 `json:"id"`
  Title string `json:"title,omitempty"`
  InviteLink string `json:"invite_link,omitempty"`
  Permissions struct {
    CanInvite bool `json:"can_invite_users,omitempty"`
  } `json:"permissions,omitempty"`
}

// This handler is called everytime telegram sends a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
  // Decode the JSON response body
  body := &webhookReqBody{}
  if err := json.NewDecoder(req.Body).Decode(body); err != nil {
    fmt.Println("could not decode request body", err)
    return
  }
  msg, msgType := UpdateGame(body.Message.Chat.ID, body.Message.User.ID, body.Message.User.Username, body.Message.Text)
  switch msgType {
    case "message":
      SendMessage(body.Message.Chat.ID, msg)
    case "reply":
      SendReply(body.Message.Chat.ID, body.Message.MessageID, msg)
    case "photo":
      SendPhoto(body.Message.Chat.ID, msg)
  }
  return
}

// The main function starts server
func main() {
  fmt.Println("start")
  port := ":8080"
  http.ListenAndServe(port, http.HandlerFunc(Handler))
}
