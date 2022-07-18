package main

import (
		  "fmt"
		  "os"
		  "strings"
)

func main() {
		  if IsOriDir("."){ 
					 fmt.Println("Ori dir!")
		  }
		  //printFilenames(".")
		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// CUE read

// CUE write

// Check dir

func IsOriDir(path string) bool {
		  pathFiles, _ := os.ReadDir(path)

		  length := len(pathFiles)

		  for i := 0; i < length; i++ {
					 if pathFiles[i].Name() == ".orichalcum" && pathFiles[i].IsDir() {
								return true
					 }
					 if pathFiles[i].Name() == "home" && pathFiles[i].IsDir() { // not sure how to best stop searhcing up, will probably be a config
								return false
					 }
		  }

		  newPath := strings.Join([]string{path, ".."}, "/")
		  return IsOriDir(newPath)
}

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
