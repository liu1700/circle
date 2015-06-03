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
