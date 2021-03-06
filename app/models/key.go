package models

import (
  "fmt"
)

const (
  CACHE_BASE        = 0x80000000
  USER_REGISTRY_KEY = CACHE_BASE + 1
  USER_BASE         = CACHE_BASE + 2
  USER_LOCATION     = CACHE_BASE + 3
  FEED_CREATE       = CACHE_BASE + 4
  MESSAGE           = CACHE_BASE + 5
  SESSION           = CACHE_BASE + 6
  FEED              = CACHE_BASE + 7

  FEED_LIST = "feedList"
)

func Session(id string) string {
  return fmt.Sprintf("%d-%s", SESSION, id)
}

func CacheKeyUserRegistry(deviceToken string) string {
  return fmt.Sprintf("%d-%s", USER_REGISTRY_KEY, deviceToken)
}

func CacheKeyUserObjectId(objectId string) string {
  return fmt.Sprintf("%d-%s", USER_BASE, objectId)
}

func CacheKeyUserLocationByUserId(userId string) string {
  return fmt.Sprintf("%d-%s", USER_LOCATION, userId)
}

func CacheKeyFeedById(feedid string) string {
  return fmt.Sprintf("%d-%s", FEED_CREATE, feedid)
}

func CacheFeedForUser(userid string) string {
  return fmt.Sprintf("%d-%s", FEED, userid)
}

func CacheMessageKey(userid string) string {
  return fmt.Sprintf("%d-%s", MESSAGE, userid)
}
