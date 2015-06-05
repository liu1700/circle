package controllers

import (
  "circle/app/models"
  "github.com/revel/revel"
  "strconv"
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

  _ = f.Request.ParseForm()
  newFeed.UserId = f.Request.Form["userid"][0]
  newFeed.Content = f.Request.Form["content"][0]
  newFeed.ImageUrl = f.Request.Form["imageUrl"][0]
  newFeed.Location = f.Request.Form["location"][0]

  newFeed.Lon, _ = strconv.ParseFloat(f.Request.Form["lon"][0], 64)
  newFeed.Lat, _ = strconv.ParseFloat(f.Request.Form["lat"][0], 64)

  err := newFeed.NewFeed()
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return f.RenderJson(response)
  }

  return f.RenderJson(response)
}
