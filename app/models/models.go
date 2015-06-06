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
func SetUserById(user *User) error {
  key := CacheKeyUserObjectId(user.UserId)
  return _cache.Set(key, user, _cache.FOREVER)
}

func SetUserByPhone(user *User) error {
  key := CacheKeyUserObjectId(user.PhoneNumber)
  return _cache.Set(key, user, _cache.FOREVER)
}

func GetUserById(userid string) *User {
  key := CacheKeyUserObjectId(userid)
  user := new(User)
  _ = _cache.Get(key, &user)
  return user
}

func GetUserByPhone(phone string) *User {
  key := CacheKeyUserObjectId(phone)
  user := new(User)
  _ = _cache.Get(key, &user)
  return user
}

func DelUserById(userid string) error {
  key := CacheKeyUserObjectId(userid)
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
  feedIds = append(feedIds, key)

  _ = _cache.Set(FEED_LIST, feedIds, _cache.FOREVER)

  return _cache.Set(key, *feed, _cache.FOREVER)
}

// 批量获取feed
func GetFeeds() []Feed {
  feedIds := []string{}
  _ = _cache.Get(FEED_LIST, &feedIds)
  getter, _ := _cache.GetMulti(feedIds...)

  feeds := make([]Feed, len(feedIds))
  if len(feedIds) == 0 {
    return feeds
  }
  for index, key := range feedIds {
    _ = getter.Get(key, &feeds[index])
  }
  return feeds
}

// 根据id获取feed
func GetFeed(feedId string) (*Feed, error) {
  key := CacheKeyFeedById(feedId)
  feed := new(Feed)
  err := _cache.Get(key, &feed)
  return feed, err
}

//Comment
func SaveComment(c Comment) error {
  comments := []Comment{}
  _ = _cache.Get(c.FeedId, &comments)
  comments = append(comments, c)
  return _cache.Set(c.FeedId, comments, _cache.FOREVER)
}

func GetComments(feedId string) []Comment {
  comments := []Comment{}
  _ = _cache.Get(feedId, &comments)
  return comments
}

// Message
func AddToMessageList(m Message) error {
  msgs := []Message{}
  key := CacheMessageKey(m.UserId)
  _ = _cache.Get(key, &msgs)
  msgs = append(msgs, m)
  return _cache.Set(key, msgs, _cache.FOREVER)
}

func GetMessages(userid string) []Message {
  msgs := []Message{}
  key := CacheMessageKey(userid)
  _ = _cache.Get(key, &msgs)
  return msgs
}

func SetMsgChecked(userid string, msgid string) {
  msgs := []Message{}
  key := CacheMessageKey(userid)
  _ = _cache.Get(key, &msgs)

  for _, m := range msgs {
    if m.MessageId == msgid {
      m.Checked = 1
      break
    }
  }
}
