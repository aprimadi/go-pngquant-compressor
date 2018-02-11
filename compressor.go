package main

import (
  "fmt"
  "image/png"
  "os"
  "path/filepath"
  "strings"

  pngquant "github.com/yusukebe/go-pngquant"
)

func main() {
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    if err != nil {
      fmt.Println(err)
      return nil
    }

    if info.IsDir() {
      return nil
    }
    if strings.ToLower(filepath.Ext(info.Name())) == ".png" {
      finput, err := os.Open(path)
      if err != nil {
        fmt.Println(err)
        return nil
      }
      img, err := png.Decode(finput)
      if err != nil {
        fmt.Println(err)
        return nil
      }
      cimg, err := pngquant.Compress(img, "1")
      if err != nil {
        fmt.Println(err)
        return nil
      }
      f, err := os.Create(path)
      err = png.Encode(f, cimg)
      if err != nil {
        fmt.Println(err)
        return nil
      }
    }
    return nil
  })
}
