package controllers

import (
  "circle/app/models"
  "github.com/revel/revel"
)

type Comment struct {
  App
}

func (c Comment) PostComment(feedId string) revel.Result {
  response := new(Response)
  response.Success = true

  newComment := new(models.Comment)

  _ = c.Request.ParseForm()
  newComment.FeedId = feedId
  newComment.UserId = c.Request.Form["userid"][0]
  newComment.Content = c.Request.Form["content"][0]
  newComment.Nickname = c.Request.Form["nickname"][0]

  err := newComment.NewComment()
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return c.RenderJson(response)
  }
  revel.INFO.Println(response)
  return c.RenderJson(response)
}

func (c Comment) GetComments(feedId string) revel.Result {
  response := new(Response)
  response.Success = true

  comments := models.GetComments(feedId)
  response.Comment = comments
  return c.RenderJson(response)
}
