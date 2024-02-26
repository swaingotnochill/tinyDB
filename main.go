package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
  err := SaveData1("./saveData1", []byte("Hello, World"))
  if err != nil {
    fmt.Println(err)
  }

  err = SaveData2("./saveData2", []byte("Hello, World"))
  if err != nil {
    fmt.Println(err)
  }

  err = SaveData3("./saveData3", []byte("Hello, Mother Flippers."))
  if err != nil {
    fmt.Println(err)
  }

 lg, err := LogCreate("./saveData4")
 if err != nil {
    fmt.Println(err)
 }
  err = appendLogs(lg, "Pushing to the logs696969....")
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

func SaveData2(path string, data []byte) error {
  rand.Seed(time.Now().UnixNano())
  tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())

  fp, err := os.OpenFile(tmp, os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0664)
  if err != nil {
    return err
  }
  defer fp.Close()

  _, err = fp.Write(data)
  if err != nil {
    os.Remove(tmp)
    return err
  }
  return os.Rename(tmp, path)
}

func SaveData3(path string, data []byte) error {
  rand.Seed(time.Now().UnixNano())
  tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())

  fp, err := os.OpenFile(tmp, os.O_WRONLY | os.O_TRUNC | os.O_CREATE, 0664)
  if err != nil {
    return err
  }
  defer fp.Close()

  _, err = fp.Write(data)
  if err != nil {
    os.Remove(tmp)
    return err
  }

  err = fp.Sync() // fsync
  if err != nil {
    os.Remove(tmp)
    return err
  }
  return os.Rename(tmp, path)
}

func LogCreate(path string) (*os.File, error) {
  return os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0664)
}

func appendLogs(fp *os.File, line string) error {
  buf := []byte(line)
  buf = append(buf, '\n')

  _, err := fp.Write(buf)
  if err != nil {
    return err
  }
  return fp.Sync()
}