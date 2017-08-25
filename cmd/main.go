package main

import (
  "log"
  "github.com/fsnotify/fsnotify"
  //"github.com/urfave/cli"
  "../internal/file_sweeper"
  "fmt"
  "strings"
  "errors"
  "os"
)

func main() {
    maxSize, err := parseSize(os.Args[1])
  if err != nil { log.Fatal(err) }
  targetSize, err := parseSize(os.Args[2])
  if err != nil { log.Fatal(err) }
  path := os.Args[3]

  log.Printf("Initializing. Path: %s | Max size: %d | Target size: %d", path, maxSize, targetSize)

  onWrite := func() { file_sweeper.Run(path, maxSize, targetSize) }
  onWrite()
  fileWatcher(path, onWrite)
}

func parseSize(formattedSize string) (bytes int64, err error) {
  formattedSize = strings.TrimSpace(formattedSize)

  _, err = fmt.Sscanf(formattedSize, "%dKB", &bytes)
  if err == nil { return bytes * 1024, nil }

  _, err = fmt.Sscanf(formattedSize, "%dMB", &bytes)
  if err == nil { return bytes * 1024 * 1024, nil}

  _, err = fmt.Sscanf(formattedSize, "%dGB", &bytes)
  if err == nil { return bytes * 1024 * 1024 * 1024, nil }

  return -1, errors.New("Unknown size format: " + formattedSize)
}

func fileWatcher(path string, onWrite func()) {
  watcher, err := fsnotify.NewWatcher()
  if err != nil { log.Fatal(err) }
  defer watcher.Close()

  done := make(chan bool)
  go func() {
    for {
      select {
      case event := <-watcher.Events:
        if event.Op == fsnotify.Write || event.Op == fsnotify.Create {
          log.Printf("Write detected: %s", event.Name)
          onWrite()
        }
      case err := <-watcher.Errors:
        log.Println("Error: %s", err)
      }
    }
  }()

  err = watcher.Add(path)
  if err != nil { log.Fatal(err) }
  <-done
}
