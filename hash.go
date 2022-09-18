package main

import (
		  "os"
		  "log"
		  "crypto/sha512"
)
// Hash utils
func OpenFile(filename string) []byte{

		  f, ok := os.ReadFile(filename)
		  if ok != nil {
					 log.Fatal(ok)
		  }
		  return f 
}

func FileSize(filename string) int64 {
		  return os.Stat(filename).Size() 
}

func HashFile(filename string) [sha512.Size]byte{
		  f := OpenFile(filename)
		  hash := sha512.Sum512(f)
		  return hash 
}
