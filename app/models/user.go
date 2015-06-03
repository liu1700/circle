package models

import (
  "fmt"
  "github.com/nu7hatch/gouuid"
  "strconv"
  "time"
)

type (
  Register struct {
    DeviceToken string `json:"device_token"`
    PhoneNumber string `json:"phone_number"`
    SMSCode     string `json:"sms_code"`
  }

  User struct {
    UserId      string `json:"user_id"`
    Nickname    string `json:"nickname"`
    DeviceToken string `json:"device_token"`
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
    AvatarId    string `json:"avatar_id"`
    CreateAt    int64  `json:"create_at"`
  }

  UserLocation struct {
    UserId string  `json:"user_id"`
    Lat    float64 `json:"lat"`
    Lon    float64 `json:"lon"`
  }
)

func (u *User) NewUser() {
  uid, _ := uuid.NewV4()
  u.UserId = uid.String()
  u.CreateAt = time.Now().Unix()
  u.Nickname = strconv.FormatInt(u.CreateAt, 10)
  err := SetUserByPhone(u)
  if err != nil {
    fmt.Println(err.Error())
  }
}
