package main

import (
		  "os"
		  "log"
		  "crypto/sha512"
		  "time"
)
// Hash utils
func OpenFile(filename string) []byte{

		  f, ok := os.ReadFile(filename)
		  if ok != nil {
					 log.Fatal(ok)
		  }
		  return f 
}


// Get the file size in bytes of a given file
func FileSize(filename string) int64 {
	f, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	return f.Size() 
}

func HashFile(filename string) [sha512.Size]byte{
		  f := OpenFile(filename)
		  hash := sha512.Sum512(f)
		  return hash 
}


// File stat informatoin.

//TODO : Struct for os file stats
// Gets last mod time
func FileChanged(filename string) time.Time {
	f, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	return f.ModTime() 
}
