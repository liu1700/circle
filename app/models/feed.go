package models

import (
  "github.com/nu7hatch/gouuid"
  "time"
)

type Feed struct {
  FeedId      string  `json:"feedId"`
  UserId      string  `json:"userId"`
  UserAvatar  string  `json:"userAvatar"`
  Content     string  `json:"content"`
  ImageUrl    string  `json:"imageUrl,omitempty"`
  CreateAt    int64   `json:"createAt"`
  ExpiredAt   int64   `json:"expiredAt"`
  Expired     bool    `json:"expired"`
  CommentCont int     `json:"commentCount"`
  Lon         float64 `json:"lon"`
  Lat         float64 `json:"lat"`
  Location    string  `json:"location"`
}

const (
  ONE_DAY = 86400
)

func (f *Feed) NewFeed() error {
  fid, _ := uuid.NewV4()
  f.FeedId = fid.String()

  f.CreateAt = time.Now().Unix()
  f.ExpiredAt = f.CreateAt + ONE_DAY
  f.Expired = false

  return SaveFeed(f)
}
