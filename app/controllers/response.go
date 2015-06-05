package controllers

import (
  "circle/app/models"
)

type (
  Response struct {
    Success bool           `json:"success"`
    Error   string         `json:"errors,omitempty"`
    User    *models.User   `json:"user,omitempty"`
    Feed    []*models.Feed `json:"feed,omitempty"`
    Image   string         `json:"image,omitempty"`
  }
)
