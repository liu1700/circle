package controllers

import (
  "circle/app/models"
  "github.com/revel/revel"
)

type App struct {
  *revel.Controller
}

func (c App) Index() revel.Result {
  return c.Render()
}

func (c App) Check(deviceId string, userid string) revel.Result {
  response := new(Response)
  response.Success = true

  findUid := models.CheckSession(userid)
  if findUid {
    user := models.GetUserById(userid)
    revel.INFO.Println(user)
    if user.UserId != "" && user.DeviceToken == deviceId {
      response.User = user
      return c.RenderJson(response)
    }

    _ = models.DelSession(userid)
    response.Success = false
    return c.RenderJson(response)
  }
  response.Success = false
  return c.RenderJson(response)
}
