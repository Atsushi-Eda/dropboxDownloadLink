package main
import (
  "strings"
  "github.com/atotto/clipboard"
)
func main() {
  text, err := clipboard.ReadAll()
  if err != nil {
    panic(err)
  }
  text = strings.Replace(text, "www.dropbox.com", "dl.dropboxusercontent.com", 1)
  text = strings.Replace(text, "?dl=0", "", 1)
  if err := clipboard.WriteAll(text); err != nil {
    panic(err)
  }
}