package controllers

import (
  "github.com/nu7hatch/gouuid"
  "github.com/revel/revel"
  "io"
  "os"
  "path"
)

type Image struct {
  App
}

const (
  FILE_UPLOAD_PATH = "/Users/liuyuchen/gocode/src/circle/upload/"
)

func (c Image) Upload() revel.Result {
  response := new(Response)
  response.Success = true
  c.Request.ParseMultipartForm(100000)
  m := c.Request.MultipartForm
  revel.INFO.Println(m)
  if image, handler, e := c.Request.FormFile("file"); e == nil {

    // ensure file should be closed later
    defer image.Close()

    // generate uuid string as iamge file name
    uuid, _ := uuid.NewV4()

    u := uuid.String()

    // create a new file name with origin extensions
    filename := u + path.Ext(handler.Filename)

    // create file for write
    if f, e := os.Create(FILE_UPLOAD_PATH + filename); e == nil {

      b := make([]byte, 1024*1024)

      for {
        if n, e := image.Read(b); e == nil {
          if _, e = f.Write(b[:n]); e != nil {
            response.Success = false
            response.Error = "上传图片失败"
            f.Close()
          }
        } else if e == io.EOF {
          f.Close()
          break
        }
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
