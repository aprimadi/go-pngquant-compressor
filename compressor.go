package main

import (
  "fmt"
  "image/png"
  "os"
  "path/filepath"
  "regexp"
  "strings"

  pngquant "github.com/yusukebe/go-pngquant"
)

func main() {
  re, err := regexp.Compile("\\/original\\/")
  if err != nil {
    panic(err)
  }

  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      panic(err)
    }
    if re.Find([]byte(path)) != nil {
      fmt.Println(fmt.Sprintf("Skipping %s", path))
      return nil
    }

    if info.IsDir() {
      return nil
    }
    ext := strings.ToLower(filepath.Ext(info.Name()))
    if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
      finput, err := os.Open(path)
      if err != nil {
        panic(err)
      }
      img, err := png.Decode(finput)
      if err != nil {
        fmt.Println(fmt.Sprintf("Skipping %s: not a png image", path))
        return nil
      }
      cimg, err := pngquant.Compress(img, "1")
      if err != nil {
        panic(err)
      }
      f, err := os.Create(path)
      err = png.Encode(f, cimg)
      if err != nil {
        panic(err)
      }
      fmt.Println(fmt.Sprintf("Processing: %s", path))
    }
    return nil
  })
}
