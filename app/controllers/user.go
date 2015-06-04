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

  if registry.DeviceToken == "" {
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
    code := randCode(6)
    registry.SMSCode = code

    revel.INFO.Println(code)

    // err = sendSms(registry.PhoneNumber, code)
    // if err != nil {
    // e := models.SetUserRegistry(registry)
    //   revel.ERROR.Println(e.Error())
    // }
    _ = models.SetUserRegistry(registry)
  }()
  return c.RenderJson(response)
}

/**
 * 用户注册
 */
func (c User) Registry(device string, smscode string) revel.Result {
  var err error

  req := new(models.User)

  response := new(Response)
  response.Success = true

  // 验证smscode
  if device == "" || smscode == "" {
    response.Success = false
    response.Error = "错误的请求"
    return c.RenderJson(response)
  }

  reg, err := models.GetUserRegistry(device)
  if err != nil {
    response.Success = false
    response.Error = "手机与验证码不匹配"
    return c.RenderJson(response)
  }

  if reg.DeviceToken != device || reg.SMSCode != smscode {
    response.Success = false
    response.Error = "验证码有误"
    return c.RenderJson(response)
  }

  // 删除注册用缓存
  _ = models.DelUserRegistry(reg)

  // 解析内容
  err = json.NewDecoder(c.Request.Body).Decode(&req)
  if err != nil {
    response.Success = false
    response.Error = "错误的请求"
    return c.RenderJson(response)
  }

  err = req.NewUser()
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return c.RenderJson(response)
  }

  c.Session["user"] = req.DeviceToken

  err = models.SetUserByPhone(req)
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return c.RenderJson(response)
  }

  return c.RenderJson(response)
}

/**
 * 用户登录
 */
func (c User) SignIn() revel.Result {
  var (
    err  error
    user *models.User
  )

  type signin struct {
    Account  string `json:"phone"`
    Password string `json:"password"`
  }
  req := new(signin)

  _ = c.Request.ParseForm()

  response := new(Response)
  response.Success = true
  req.Account = c.Request.Form["phone"][0]
  req.Password = c.Request.Form["password"][0]

  user, err = models.GetUserByPhone(req.Account)

  if user == nil {
    response.Success = false
    response.Error = "用户不存在"
    return c.RenderJson(response)
  }

  if req.Password != user.Password {
    response.Success = false
    response.Error = "密码有误"
    return c.RenderJson(response)
  }

  c.Session["user"] = user.DeviceToken
  revel.INFO.Println(user)
  return c.RenderJson(response)
}

/**
 * 用户登出
 */
func (c User) SignOut() revel.Result {
  response := new(Response)
  response.Success = true

  for key := range c.Session {
    delete(c.Session, key)
  }

  return c.RenderJson(response)
}

/**
 * 用户更新名字
 */
func (c User) UpdateNickname() revel.Result {
  req := new(models.User)
  response := new(Response)
  response.Success = true

  // 解析内容
  err := json.NewDecoder(c.Request.Body).Decode(&req)
  if err != nil {
    response.Success = false
    response.Error = "错误的请求"
    return c.RenderJson(response)
  }

  user, err := models.GetUserByPhone(req.PhoneNumber)
  if user.UserId == "" {
    response.Success = false
    response.Error = "用户不存在"
    return c.RenderJson(response)
  }
  user.Nickname = req.Nickname
  models.SetUserByPhone(user)

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
  return err
}
