package main

import (
  "strconv"
)

type Player struct {
  ID int `json:"_id" bson:"_id"`
  Name string `json:"name" bson:"name"`
  Score int `json:"score" bson:"score"`
}

type Game struct {
  ID int64 `json:"_id" bson:"_id"`
  Status string `json:"status" bson:"status"`
  Players []Player `json:"players" bson: "players"`
  Answers []string `json:"answers" bson: "answers"`
}

func UpdateGame(gameID int64, playerID int, playerName string, cmd string) (string, string) {
  if len(cmd) == 0 {
    return "",""
  }
  if cmd[0] != '/' {
    return "",""
  }
  switch cmd {
    case "/startflaggame":
      return CreateGame(gameID)
    case "/iamin":
      return AddPlayer(gameID, playerID, playerName)
    case "/everybodyisin":
      return StartGame(gameID)
    case "/nextflag":
      return NextFlag(gameID)
    case "/endflaggame":
      return EndGame(gameID)
    default:
      return CheckAnswer(gameID, playerID, cmd)
  }
}

func GetGameStatus(id int64) string {
  var game Game
  if SearchInstance(id, "games", &game) {
    return game.Status
  }
  return "not created"
}

func CreateGame(id int64) (string, string) {
  status := GetGameStatus(id)
  game := Game {
    ID: id,
    Status: "waiting players",
    Players: make([]Player, 0, 0),
    Answers: make([]string, 0, 0),
  }
  if status == "not created" {
    if (CreateInstance(game, "games")) {
      return "Let's play! Who is in?", "message"
    }
    return "OPS! something went wrong...", "reply"
  } else if status == "finished" {
    if UpdateInstance(id, game, "games"){
      return "Let's play! Who is in?", "message"
    }
    return "OPS! something went wrong...", "reply"
  }
  return "Not the moment...", "reply"
}

func AddPlayer(gameID int64, playerID int, playerName string) (string, string) {
  status := GetGameStatus(gameID)
  if status == "waiting players" {
    var game Game
    if SearchInstance(gameID, "games", &game) {
      if len(game.Players) < 10 {
        for i := 0; i < len(game.Players); i++ {
          if game.Players[i].ID == playerID {
              return "You are already in.", "reply"
          }
        }
        player := Player {
          ID: playerID,
          Name: playerName,
          Score: 0,
        }
        game.Players = append(game.Players, []Player{player}...)
        if UpdateInstance(gameID, game, "games") {
          return playerName + " is in the game!", "message"
        }
        return "OPS! something went wrong...", "reply"
      }
      return "Maximum of 30 players, sorry.", "reply"
    }
    return "OPS! something went wrong...", "reply"
  }
  return "Not the moment...", "reply"
}

func StartGame(id int64) (string, string) {
  status := GetGameStatus(id)
  if status == "waiting players" {
    var game Game
    if SearchInstance(id, "games", &game) {
      if len(game.Players) > 0 {
        flag := DrawFlag()
        game.Status = "waiting answer"
        game.Answers = flag.Names
        if UpdateInstance(id, game, "games") {
          return flag.Link, "photo"
        }
        return "OPS! something went wrong...", "reply"
      }
      return "OPS! I need at least 1 player in the game!", "reply"
    }
    return "OPS! something went wrong...", "reply"
  }
  return "Not the moment...", "reply"
}

func NextFlag(id int64) (string, string) {
  status := GetGameStatus(id)
  if status == "waiting ask next flag" {
    var game Game
    if SearchInstance(id, "games", &game) {
      flag := DrawFlag()
      game.Status = "waiting answer"
      game.Answers = flag.Names
      if UpdateInstance(id, game, "games") {
        return flag.Link, "photo"
      }
      return "OPS! something went wrong...", "reply"
    }
    return "OPS! something went wrong...", "reply"
  }
  return "Not the moment...", "reply"
}

func CheckAnswer(gameID int64, playerID int, aws string) (string, string) {
  if GetGameStatus(gameID) != "waiting answer" {
    return "", ""
  }
  var game Game
  if !SearchInstance(gameID, "games", &game) {
    return "OPS! something went wrong...", "reply"
  }
  i := 0
  for i = 0; i < len(game.Players); i++ {
    if game.Players[i].ID == playerID {
      break
    }
  }
  if i == len(game.Players) {
    return "You are not in the game.", "reply"
  }
  for j := 0; j < len(game.Answers); j++ {
    if aws == "/" + game.Answers[j] {
      game.Status = "waiting ask next flag"
      game.Players[i].Score++
      if UpdateInstance(gameID, game, "games") {
        return "CORRECT! =D", "reply"
      }
      return "OPS! something went wrong1...", "reply"
    }
  }
  return "", ""
}

func EndGame(id int64) (string, string) {
  status := GetGameStatus(id)
  if status == "finished" || status == "not created" {
    return "The game has already ended.", "reply"
  }
  var game Game
  if SearchInstance(id, "games", &game) {
    finalScores := "FINAL SCORES:\n"
    for i := 0; i < len(game.Players); i++ {
      finalScores = finalScores + game.Players[i].Name + " " + strconv.Itoa(game.Players[i].Score) + "\n"
    }
    game.Players = make([]Player,0,0)
    game.Answers = make([]string,0,0)
    game.Status = "finished"
    if UpdateInstance(id, game, "games") {
      return finalScores, "message"
    }
  }
  return "OPS! something went wrong...", "reply"
}
