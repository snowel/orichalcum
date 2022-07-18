package main

import (
		  "fmt"
		  "os"
		  "strings"
)

func main() {

		  printFilenames(".")
		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// CUE read

// CUE write

// recursive walk

func printFilenames(rootdir string) {
		  filesAndDirs, ok := os.ReadDir(rootdir)

		  if ok != nil {
					 fmt.Println("Something went wrong...")
					 return
		  }

		  length := len(filesAndDirs)

		  for i := 0; i < length; i++ {
					 if filesAndDirs[i].IsDir() {
								newRootDir := []string{rootdir, filesAndDirs[i].Name()}
								printFilenames(strings.Join(newRootDir, "/"))
					 } else {
								path := strings.Join([]string{rootdir, filesAndDirs[i].Name()}, "/")
								fmt.Println(path)
					 }
		  }
}


// per file hash + object create/update

// per dir create/update ???
