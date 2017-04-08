package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  dups := make(map[string]int32)
  file := bufio.NewScanner(os.Stdin)
  for file.Scan() {
    dups[file.Text()]++
  }

  for k, v := range dups {
    if v > 1 {
     fmt.Printf("The line with text '%s' appeared %d times\n", k, v)
    }
  }
}
