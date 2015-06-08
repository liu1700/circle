package controllers

import (
  "circle/app/models"
  "github.com/kellydunn/golang-geo"
  "github.com/revel/revel"
  "strconv"
)

type Feed struct {
  App
}

func (f Feed) GetFeeds(lon float64, lat float64, distance float64) revel.Result {
  response := new(Response)
  response.Success = true
  respFeeds := []models.Feed{}

  feeds := models.GetFeeds()

  userPosition := geo.NewPoint(lat, lon)
  for _, f := range feeds {
    feedPosition := geo.NewPoint(f.Lat, f.Lon)
    km := userPosition.GreatCircleDistance(feedPosition)
    if km > distance {
      continue
    }
    comments := models.GetComments(f.FeedId)
    f.CommentCont = len(comments)
    revel.INFO.Println(km)
    respFeeds = append(respFeeds, f)
  }

  response.Feed = respFeeds
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
  response.Feed = append(response.Feed, *feed)

  return f.RenderJson(response)
}

func (f Feed) GetUserFeeds(userid string) revel.Result {
  response := new(Response)
  response.Success = true

  feeds := models.GetUserFeeds(userid)
  response.Feed = feeds

  return f.RenderJson(response)
}

func (f Feed) CreateFeed() revel.Result {
  response := new(Response)
  response.Success = true

  newFeed := new(models.Feed)

  _ = f.Request.ParseForm()
  newFeed.UserId = f.Request.Form["userid"][0]
  newFeed.Content = f.Request.Form["content"][0]
  if f.Request.Form["imageId"] != nil {
    newFeed.ImageUrl = f.Request.Form["imageId"][0]
  }
  newFeed.Location = f.Request.Form["location"][0]

  newFeed.Lon, _ = strconv.ParseFloat(f.Request.Form["lon"][0], 64)
  newFeed.Lat, _ = strconv.ParseFloat(f.Request.Form["lat"][0], 64)

  revel.INFO.Println(newFeed)

  err := newFeed.NewFeed()
  if err != nil {
    response.Success = false
    response.Error = err.Error()
    return f.RenderJson(response)
  }

  return f.RenderJson(response)
}
