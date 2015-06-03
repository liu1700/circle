package models

import (
  "fmt"
  "github.com/nu7hatch/gouuid"
  "strconv"
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

func (f *Feed) NewFeed() {

}
