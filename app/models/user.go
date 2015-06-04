package models

import (
  "fmt"
  "github.com/nu7hatch/gouuid"
  "strconv"
  "time"
)

type (
  Register struct {
    DeviceToken string `json:"deviceToken"`
    PhoneNumber string `json:"phone"`
    SMSCode     string `json:"smsCode"`
  }

  User struct {
    UserId      string `json:"userId"`
    Nickname    string `json:"nickname"`
    DeviceToken string `json:"deviceToken"`
    PhoneNumber string `json:"phone"`
    Password    string `json:"password"`
    AvatarId    string `json:"avatarId,omitempty"`
    CreateAt    int64  `json:"createAt"`
  }

  UserLocation struct {
    UserId string  `json:"userId"`
    Lat    float64 `json:"lat"`
    Lon    float64 `json:"lon"`
  }
)

func (u *User) NewUser() error {
  uid, _ := uuid.NewV4()
  u.UserId = uid.String()
  u.CreateAt = time.Now().Unix()
  u.Nickname = strconv.FormatInt(u.CreateAt, 10)
  err := SetUserByPhone(u)
  if err != nil {
    fmt.Println(err.Error())
  }
  return err
}
