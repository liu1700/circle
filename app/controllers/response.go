package controllers

import (
  "circle/app/models"
)

type (
  Response struct {
    Success bool         `json:"success"`
    Error   string       `json:"error,omitempty"`
    User    *models.User `json:"user,omitempty"`
  }

  /**
   * Image
   */
  UploadImage struct {
    Id       string `json:"imageId"`
    Filename string `json:"imageName"`
  }
)
