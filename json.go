package main

import (
		  "fmt"
		  "os"
		  "encoding/json"
)

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
