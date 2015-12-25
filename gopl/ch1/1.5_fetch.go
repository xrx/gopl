// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Fetch prints the content found at each specified URL.
//
// Modified version with solutions for exercises 1.7, 1.8, 1.9
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    if !strings.HasPrefix(url, "http://"){
      url = "http://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    _, err = io.Copy(os.Stdout, resp.Body)
    status := resp.Status
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("\nStatus: %s\n", status)
  }
}