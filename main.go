package main

import (
  "flag"
  "fmt"
  "os"
)

var flagFilename string
var flagBenchmark bool

func init() {
  flag.StringVar(&flagFilename, "f", "", "Path to word.list")
  flag.BoolVar(&flagBenchmark, "b", false, "Benchmark")
  flag.Parse()
}

func main() {
  if len(flagFilename) == 0 {
    fmt.Printf("Usage: %s -f filename\n", os.Args[0])
    os.Exit(1)
  }
}