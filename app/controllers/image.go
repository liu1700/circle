package controllers

import (
  "bytes"
  "github.com/nfnt/resize"
  "github.com/revel/revel"
  "image"
  "image/jpeg"
  "os"
)

type Image struct {
  App
}

const (
  FILE_UPLOAD_PATH = "/Users/liuyuchen/gocode/src/circle/upload/"
  MB               = 1 << 20 // 1MB
)

func (c Image) Upload(width uint, height uint) revel.Result {
  response := new(Response)
  response.Success = true
  e := c.Request.ParseMultipartForm(10 * MB)
  if e != nil {
    revel.INFO.Println(e.Error())
  }
  if imagePart, handler, e := c.Request.FormFile("file"); e == nil {

    // ensure file should be closed later
    defer imagePart.Close()
    imageStream := make([]byte, 10*MB)

    revel.INFO.Println(handler.Filename)
    imagePart.Read(imageStream)
    img, str, err := image.Decode(bytes.NewReader(imageStream))
    if err != nil {
      revel.ERROR.Println(err.Error())
      response.Success = false
      response.Error = "解析图片失败"
    }
    revel.INFO.Println(str)

    newImage := resize.Resize(width, height, img, resize.Lanczos2)

    // create file for write
    if f, e := os.Create(FILE_UPLOAD_PATH + handler.Filename + ".jpg"); e == nil {

      err := jpeg.Encode(f, newImage, nil)
      if err != nil {
        revel.ERROR.Printf(err.Error())
        response.Success = false
        response.Error = "写入图片失败"
      }

    } else {
      revel.ERROR.Printf(e.Error())
      response.Success = false
      response.Error = "上传图片失败"
    }
  } else {
    revel.ERROR.Printf(e.Error())
    response.Success = false
    response.Error = "上传图片失败"
  }
  return c.RenderJson(response)
}
