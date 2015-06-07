package controllers

import (
  "bytes"
  "errors"
  "github.com/nfnt/resize"
  "github.com/nu7hatch/gouuid"
  "github.com/revel/revel"
  "image"
  "image/jpeg"
  "os"
)

type Image struct {
  App
}

const (
  AVATAR_UPLOAD_PATH = "/upload/avatar/"
  IMAGE_UPLOAD_PATH  = "/upload/images/"
  MB                 = 1 << 20 // 1MB
)

var ROOT string

func init() {
  ROOT, _ = os.Getwd()
}

func (c Image) UploadAvatar(width uint, height uint) revel.Result {
  response := new(Response)
  response.Success = true

  success, err, _ := processImage(c, ROOT+AVATAR_UPLOAD_PATH, width, height)
  if !success {
    response.Success = false
    response.Error = err.Error()
    return c.RenderJson(response)
  }

  return c.RenderJson(response)
}

func (c Image) UploadImage(width uint, height uint) revel.Result {
  response := new(Response)
  response.Success = true

  success, err, fileName := processImage(c, ROOT+IMAGE_UPLOAD_PATH, width, height)
  if !success {
    response.Success = false
    response.Error = err.Error()
    revel.INFO.Println(response)
    return c.RenderJson(response)
  }

  response.Image = fileName

  return c.RenderJson(response)
}

func processImage(c Image, filePath string, width uint, height uint) (bool, error, string) {
  var fileName string

  e := c.Request.ParseMultipartForm(10 * MB)
  if e != nil {
    revel.INFO.Println(e.Error())
    return false, e, ""
  }
  if imagePart, handler, e := c.Request.FormFile("file"); e == nil {

    // ensure file should be closed later
    defer imagePart.Close()
    imageStream := make([]byte, 10*MB)

    revel.INFO.Println(handler.Filename)
    fileName = handler.Filename
    if fileName == "image" {
      uuid, _ := uuid.NewV4()
      fileName = uuid.String()
    }

    imagePart.Read(imageStream)
    img, str, err := image.Decode(bytes.NewReader(imageStream))
    if err != nil {
      revel.ERROR.Println(err.Error())
      return false, errors.New("解析图片失败"), ""
    }
    revel.INFO.Println(str)

    newImage := resize.Resize(width, height, img, resize.Lanczos2)

    // create file for write
    if f, e := os.Create(filePath + fileName + ".jpg"); e == nil {

      err := jpeg.Encode(f, newImage, nil)
      if err != nil {
        revel.ERROR.Printf(err.Error())
        return false, errors.New("写入图片失败"), ""
      }

    } else {
      revel.ERROR.Printf(e.Error())
      return false, errors.New("上传图片失败"), ""
    }
  } else {
    revel.ERROR.Printf(e.Error())
    return false, errors.New("上传图片失败"), ""
  }

  return true, nil, fileName
}
