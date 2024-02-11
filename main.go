package main

import (
	"fmt"
	"os"
)

func main() {
  err := SaveData1("./saveData1", []byte("Hello, World"))
  if err != nil {
    fmt.Println(err)
  }
}

func SaveData1(path string, data []byte) error {
  fp, err := os.OpenFile(path, os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0664)
  if err != nil {
    return err
  }

  defer fp.Close()
  _, err = fp.Write(data)
  return err
}
