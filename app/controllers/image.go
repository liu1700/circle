package controllers

import (
  "bytes"
  "errors"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/nfnt/resize"
  "github.com/nu7hatch/gouuid"
  "github.com/revel/revel"
  "image"
  "image/png"
  "os"
)

type Image struct {
  App
}

const (
  AWS_REGION         = "ap-northeast-1"
  IMAGE_BUCKET_NAME  = "circle-android"
  AVATAR_UPLOAD_PATH = "avatars/"
  IMAGE_UPLOAD_PATH  = "images/"
  MB                 = 1 << 20 // 1MB
)

var ROOT string

func init() {
  ROOT, _ = os.Getwd()
}

func (c Image) UploadAvatar(width uint, height uint) revel.Result {
  response := new(Response)
  response.Success = true

  success, err, _ := processImage(c, AVATAR_UPLOAD_PATH, width, height)
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

  success, err, fileName := processImage(c, IMAGE_UPLOAD_PATH, width, height)
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

    fileName = handler.Filename
    if fileName == "image" {
      uuid, _ := uuid.NewV4()
      fileName = uuid.String()
    }

    imagePart.Read(imageStream)
    img, str, err := image.Decode(bytes.NewReader(imageStream))
    if err != nil {
      revel.ERROR.Println(err.Error())
      return false, errors.New("Error when Decode"), ""
    }

    newImage := resize.Resize(width, height, img, resize.Lanczos2)
    buffer := new(bytes.Buffer)
    e = png.Encode(buffer, newImage)
    if e != nil {
      revel.ERROR.Printf(e.Error())
      return false, errors.New("Error when Encode"), ""
    }

    go func() {
      // Upload image to aws s3
      cred := NewEnvCredentials()
      credValue, err := creds.Get()
      if err != nil {
        revel.ERROR.Println(err.Error())
      }
      svc := s3.New(&aws.Config{
        Region:      aws.String(AWS_REGION),
        Credentials: credValue,
      })

      params := &s3.PutObjectInput{
        Bucket: aws.String(IMAGE_BUCKET_NAME),
        Key:    aws.String(fileName + ".png"),
        Body:   bytes.NewReader(buffer.Bytes()),
      }
      _, err = svc.PutObject(params)
      if err != nil {
        revel.ERROR.Println(err.Error())
      }
    }()

  } else {
    revel.ERROR.Printf(e.Error())
    return false, errors.New("上传图片失败"), ""
  }

  return true, nil, fileName
}
