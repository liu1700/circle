package models

import (
  _cache "github.com/revel/revel/cache"
  // "log"
  "time"
)

// user registry
func SetUserRegistry(registry *Register) error {
  key := CacheKeyUserRegistry(registry.DeviceToken)
  return _cache.Set(key, registry, time.Duration(time.Minute*30))
}

func GetUserRegistry(deviceToken string) (*Register, error) {
  key := CacheKeyUserRegistry(deviceToken)
  registry := new(Register)
  err := _cache.Get(key, &registry)
  return registry, err
}

func DelUserRegistry(registry *Register) error {
  key := CacheKeyUserRegistry(registry.DeviceToken)
  return _cache.Delete(key)
}

// user
func SetUserByPhone(user *User) error {
  key := CacheKeyUserObjectId(user.PhoneNumber)
  return _cache.Set(key, user, _cache.FOREVER)
}

func GetUserByPhone(phone string) (*User, error) {
  key := CacheKeyUserObjectId(phone)
  user := new(User)
  err := _cache.Get(key, &user)
  return user, err
}

func DelUserByPhone(phone string) error {
  key := CacheKeyUserObjectId(phone)
  return _cache.Delete(key)
}

// user location
func UpdateUserLocation(location *UserLocation) error {
  key := CacheKeyUserLocationByUserId(location.UserId)
  return _cache.Set(key, location, _cache.FOREVER)
}

func GetUserLocation(id string) (*UserLocation, error) {
  key := CacheKeyUserLocationByUserId(id)
  location := new(UserLocation)
  err := _cache.Get(key, &location)
  return location, err
}

// feed
func SaveFeed(feed *Feed) error {
  feedIds := []string{}
  key := CacheKeyFeedById(feed.FeedId)
  _ = _cache.Get(FEED_LIST, &feedIds)
  feedIds = append(feedIds, feed.FeedId)

  _ = _cache.Set(FEED_LIST, value, expires)

  return _cache.Set(key, feed, _cache.FOREVER)
}

func GetFeeds() []*Feed {
  feedIds := []string{}
  _ = _cache.Get(FEED_LIST, &feedIds)
  getter, _ := _cache.GetMulti(feedIds...)
  feeds := make([]*Feed, len(feedIds))
  for index, key := range feedIds {
    _ = getter.Get(key, feeds[index])
  }
  return feeds
}
