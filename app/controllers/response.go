package controllers

import (
  "circle/app/models"
)

type (
  Response struct {
    Success bool             `json:"success"`
    Error   string           `json:"errors,omitempty"`
    User    *models.User     `json:"user,omitempty"`
    Feed    []models.Feed    `json:"feeds,omitempty"`
    Comment []models.Comment `json:"comments,omitempty"`
    Image   string           `json:"image,omitempty"`
  }
)
