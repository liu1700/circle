package models

import (
  "github.com/nu7hatch/gouuid"
  "time"
)

type Feed struct {
  FeedId    string  `json:"feed_id"`
  UserId    string  `json:"user_id"`
  Content   string  `json:"content"`
  ImageUrl  string  `json:"image_url,omitempty"`
  CreateAt  int64   `json:"create_at"`
  ExpiredAt int64   `json:"expired_at"`
  Expired   bool    `json:"expired"`
  Lon       float64 `json:"lon"`
  Lat       float64 `json:"lat"`
  Loation   string  `json:"location"`
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
