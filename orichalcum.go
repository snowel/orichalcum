package main

import (
  "time"
  "os"
)

func main() {
  if len(os.Args) == 1 { SelfSync() }

  if len(os.Args) > 1 {
	 switch os.Args[1]{
		case "sync": {break}
		case "set": {break}
		case "help": {break}
		case "debug": {break}
	 }
  }
}

// Meta information update on tracked file update.
func OnUpdate(logHandle *OriLog) {
  logHandle.Meta.DateMod = time.Now().Unix()
}

//Find index of last occurence of a character in a string(for finding the name of the file without the path)
// Better would be slices lib. 
func Contains[E comparable](slice []E, elem E) bool {
  length := len(slice)

  for i := 0; i < length; i++ {
	 if slice[i] == elem {
		return true
	 }
  }
  return false
}

// For a given path pair, check if the absolute path or relative paths contain the searched elem.
// Abs path is specified by 1 and rel by 2.
func PathpairContains(slice []PathPair, elem string, path byte) bool {
  length := len(slice)

  if path == 1 { // abs path
	 for i := 0; i < length; i++ {
		if slice[i].absPath == elem {
		  return true
		}
	 }
  } else if path == 2 { // rel path
	 for i := 0; i < length; i++ {
		if slice[i].path == elem {
		  return true
		}
	 }

  }

  return false
}

// 
func LastIndex(name string, elem string) int {
  // Should not have any utf8 problems as, for now, we're woking with a unix filesystem
  for i, sub := range name {
	 if string(sub) == elem {
		return i
	 }
  }
  return -1
}

