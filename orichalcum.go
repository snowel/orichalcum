package main

import (
		  "fmt"
		  "os"
		  "log"
		  "strings"
		  "time"
		  "encoding/json"
		  "crypto/sha512"
)

func main() {
		  thisLog := OriLog{}// Need this no matter what
		  var oriRoot string
		  if IsOriDir(".") {
					 oriRoot = WhereIsOriDir(".")
					 LoadLog(oriRoot, &thisLog)
		  } else {
					 InitOriDir()
		  }

		  fmt.Println(oriRoot)

		  fmt.Println("Before update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&thisLog.FileEntries)
		  fmt.Println(oriRoot)
		  UpdateOriDir(oriRoot, &thisLog)
		  fmt.Println("After update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&thisLog.FileEntries)

		  WriteLog(&thisLog, oriRoot)

		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// JSON read
func LoadLog(oriroot string, logHandle *OriLog) {
		  fileBytes := OpenFile(oriroot + "/.orichalcum/log")
		  json.Unmarshal(fileBytes, logHandle)
}

// JSON write
func WriteLog(logHandle *OriLog, oriRoot string) {

		  jfile, ok := json.Marshal(*logHandle)
		  if ok != nil {
					 fmt.Println("Something went wrong.")
		  }
		  writeOk := os.WriteFile(oriRoot + "/.orichalcum/log", jfile, 0666)// TODO need a geenric fucntion to fetch the path of the root ori dir
		  if writeOk != nil {
					 fmt.Println("Something went wrong.")
		  }
		  
}



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

// Meta information update on tracked file update.
func OnUpdate(logHandle *OriLog) {
		  logHandle.Meta.DateMod = time.Now().Unix()
}

//Find index of last occurence of a character in a string(for finding the name of the file without the path)

// Better would be slices lib. Developing withou network connection.
func Contains[E comparable](slice []E, elem E) bool {
		  length := len(slice)

		  for i := 0; i < length; i++ {
					 if slice[i] == elem {
								return true
					 }
		  }

		  return false
}

func PathpairContains(slice []PathPair, elem string, path int) bool {
		  length := len(slice)

		  if path == 1 {// abspath
					 for i := 0; i < length; i++ {
								if slice[i].asbPath == elem {
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

