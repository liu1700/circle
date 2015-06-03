package controllers

import (
  "circle/app/models"
  "encoding/json"
  "fmt"
  "github.com/revel/revel"
  "io/ioutil"
  "math/rand"
  "net/http"
  "net/url"
  "regexp"
  "strconv"
  "strings"
  "time"
)

var (
  APIKEY          string
  SMS_SERVICE_URL string
)

type User struct {
  App
}

func init() {
  APIKEY = "b66503b1cb75230cc925a687da189a25"
  SMS_SERVICE_URL = "http://yunpian.com/v1/sms/send.json"
}

func (c User) SendCode() revel.Result {
  registry := new(models.Register)
  response := new(Response)
  response.Success = true

  err := json.NewDecoder(c.Request.Body).Decode(&registry)
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return c.RenderJson(response)
  }

  if !validPhone(registry.PhoneNumber) || registry.DeviceToken == "" {
    response.Success = false
    response.Error = "验证手机出错"
    return c.RenderJson(response)
  }

  user, err := models.GetUserByPhone(registry.PhoneNumber)
  if user.UserId != "" {
    response.Success = false
    response.Error = "手机号已注册"
    return c.RenderJson(response)
  }
  revel.INFO.Println(registry.PhoneNumber)
  go func() {
    err = sendSms(registry.PhoneNumber, randCode(6))
    if err != nil {
      e := models.SetUserRegistry(registry)
      revel.ERROR.Println(e.Error())
    }
  }()
  return c.RenderJson(response)
}

/**
 * 随机length长度的验证码
 */
func randInt(min int, max int) int {
  rand.Seed(time.Now().UTC().UnixNano())
  return min + rand.Intn(max-min)
}

func randCode(length int) string {
  var code string
  for i := 0; i < length; {
    code = code + strconv.Itoa(randInt(1, 10))
    i++
  }
  return code
}

func validPhone(phone string) bool {
  m, _ := regexp.MatchString("^([1][3-8])\\d{9}$", phone)
  return m
}

/**
 * 发送验证码
 */
func sendSms(phone string, code string) error {

  sms := url.Values{}
  sms.Set("apikey", APIKEY)
  sms.Set("mobile", phone)

  sms.Add("text", fmt.Sprintf("【Circle信息分享平台】欢迎使用Circle，您的手机验证码是%s。本条信息无需回复", code))

  client := &http.Client{}
  reader := strings.NewReader(sms.Encode())
  request, err := http.NewRequest("POST", SMS_SERVICE_URL, reader)
  if err != nil {
    revel.ERROR.Println(err.Error())
    return err
  }

  response, err := client.Do(request)
  defer response.Body.Close()
  if err != nil {
    revel.ERROR.Println(err.Error())
    return err
  }

  body, _ := ioutil.ReadAll(response.Body)
  revel.INFO.Println(string(body))
}
