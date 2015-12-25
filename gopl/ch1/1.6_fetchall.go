// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
//
// Modified version, to save the requested url to a TempFile
package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url, ch) // start a goroutine
  }
  for range os.Args[1:] {
    fmt.Println(<-ch) // receive from channel ch
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // send to channel ch
    return
  }

  splitted := strings.Split(url, "/")
  fileName := splitted[len(splitted)-1]
  file, _ := ioutil.TempFile(".", fileName)
  nbytes, err := io.Copy(file, resp.Body)
  resp.Body.Close() // don't leak resources
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
