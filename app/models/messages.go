package models

import (
  "github.com/nu7hatch/gouuid"
)

type Message struct {
  MessageId string `jsong:"messageId"`
  FeedId    string `json:"feedId"`
  UserId    string `json:"userId"`
  Nickname  string `json:"nickname"`
  Type      int    `json:"type"`
  Checked   int    `json:"checked"`
}

func (m *Message) AddMessage() error {
  uid, _ := uuid.NewV4()
  m.MessageId = uid.String()
  return AddToMessageList(*m)
}
