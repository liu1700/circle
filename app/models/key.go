package models

import (
  "fmt"
)

const (
  CACHE_BASE        = 0x80000000
  USER_REGISTRY_KEY = CACHE_BASE + 1
  USER_BASE         = CACHE_BASE + 2
  USER_LOCATION     = CACHE_BASE + 3
  FEED              = CACHE_BASE + 7
)

func CacheKeyUserRegistry(deviceToken string) string {
  return fmt.Sprintf("%d-%s", USER_REGISTRY_KEY, deviceToken)
}

func CacheKeyUserObjectId(objectId string) string {
  return fmt.Sprintf("%d-%s", USER_BASE, objectId)
}

func CacheKeyUserLocationByUserId(userId string) string {
  return fmt.Sprintf("%d-%s", USER_LOCATION, userId)
}
