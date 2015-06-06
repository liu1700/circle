package controllers

import (
  "circle/app/models"
  "github.com/revel/revel"
)

type Message struct {
  App
}

func (m Message) Get(userId string) revel.Result {
  response := new(Response)
  response.Success = true

  msgs := models.GetMessages(userId)

  response.Message = msgs

  return m.RenderJson(response)
}

func (m Message) Check(userId string, messageId string) revel.Result {
  response := new(Response)
  response.Success = true
  models.SetMsgChecked(userId, messageId)
  return m.RenderJson(response)
}
