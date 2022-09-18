package main

import (
		  "fmt"
		  "os"
		  "log"
		  "strings"
		  "time"
)

func main() {
		  sessionLog := OriLog{}// Need this no matter what
		  var toOriRoot string
		  if IsOriRepo(".") {
					 toOriRoot = WhereIsOriRoot(".")
					 LoadLog(toOriRoot, &sessionLog)
		  } else {
					 InitOriDir()
		  }

		  fmt.Println(toOriRoot)

		  fmt.Println("Before update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&sessionLog.FileEntries)
		  fmt.Println(toOriRoot)
		  UpdateOriDir(toOriRoot, &sessionLog)
		  fmt.Println("After update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&sessionLog.FileEntries)

		  WriteLog(&sessionLog, toOriRoot)

		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
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

// For a given path pair, check it the absolute path or relative paths contain the searched elem.
// Abs path is specified by 1 and rel by 2.
func PathpairContains(slice []PathPair, elem string, path byte) bool {
		  length := len(slice)

		  if path == 1 {// abspath
					 for i := 0; i < length; i++ {
								if slice[i].absPath == elem {
										  return true
								}
					 }
		  } else if path == 2 {
					 for i := 0; i < length; i++ {
								if slice[i].path == elem {
										  return true
								}
					 }

		  }

		  return false
}


func LastIndex(name string, elem string) int {
		  // SHould not have any utf8 problems as, for now, we're woking with a unix filesystem
		  for i, sub := range name {
					 if sub == elem {
								return i
					 }
		  }
		  return -1
}

