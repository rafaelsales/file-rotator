package file_sweeper

import (
  "os"
  "path/filepath"
  "log"
  "sort"
)

func Run(path string, maxSize int64, targetSize int64) {
  currentSize, _ := dirSize(path)
  log.Printf("Path: %s | Current size: %d", path, currentSize)

  if currentSize > maxSize {
    log.Printf( "Path: %s | Over max size - Current size: %d | Max size: %d", path, currentSize, maxSize)
    deleteOldFiles(path, currentSize, targetSize)
  }
}

func dirSize(path string) (int64, error) {
  var size int64
  os.Open(path)
  err := filepath.Walk(path, func(a string, info os.FileInfo, err error) error {
    if !info.IsDir() {
      size += info.Size()
    }
    return err
  })
  return size, err
}

func deleteOldFiles(path string, currentSize int64, targetSize int64) {
  var sizeDeleted int64

  dir, err := os.Open(path)
  if err != nil { log.Fatal(err) }

  entries, err := dir.Readdir(-1)
  if err != nil { log.Fatal(err) }

  sort.SliceStable(entries, func(i, j int) bool {
    return entries[i].ModTime().Before(entries[j].ModTime())
  })

  for _, entry := range entries {
    if !entry.IsDir() {
      filePath := filepath.Join(path, entry.Name())
      modifiedAt := entry.ModTime().UTC().String()
      size := entry.Size()

      log.Printf( "Deleting: %s | Modified at: %s | Size: %d", filePath, modifiedAt, size)
      os.Remove(filePath)
      sizeDeleted += size
    }

    if currentSize - sizeDeleted < targetSize {
      log.Printf( "Deleted a total of %d", sizeDeleted)
      return
    }
  }
}