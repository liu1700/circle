package models

import ()

type Comment struct {
  FeedId   string `json:"feedId"`
  UserId   string `json:"userId"`
  Nickname string `json:"nickname"`
  Content  string `json:"content"`
}

func (c *Comment) NewComment() error {
  return SaveComment(*c)
}
