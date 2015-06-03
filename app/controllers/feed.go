package controllers

import (
  "circle/app/models"
  "encoding/json"
  "github.com/revel/revel"
)

type Feed struct {
  App
}

func (f Feed) GetFeeds() revel.Result {
  response := new(Response)
  response.Success = true

  feeds := models.GetFeeds()
  response.Feed = feeds

  return f.RenderJson(response)
}

func (f Feed) GetFeed(feedId string) revel.Result {
  response := new(Response)
  response.Success = true

  feed, err := models.GetFeed(feedId)
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return f.RenderJson(response)
  }
  response.Feed = append(response.Feed, feed)

  return f.RenderJson(response)
}

func (f Feed) CreateFeed() revel.Result {
  response := new(Response)
  response.Success = true

  newFeed := new(models.Feed)
  err := json.NewDecoder(f.Request.Body).Decode(&newFeed)
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return f.RenderJson(response)
  }

  err = newFeed.NewFeed()
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return f.RenderJson(response)
  }

  return f.RenderJson(response)
}
