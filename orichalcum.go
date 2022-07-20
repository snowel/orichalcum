package main

import (
		  "fmt"
		  "os"
		  "strings"


		  "encoding/json"
)

func main() {

		  orid := WhereIsOriDir(".")

		  fmt.Println(orid)

		  //if IsOriDir("."){ 
		//			 fmt.Println("Ori dir!")
		 // }
		  //printFilenames(".")
		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// CUE read

// CUE/JSON write

func writeFileLog(files *[]OriFile, path string) {

		  oriRoot := WhereIsOriDir(path string)

		  jfile, ok := json.Marshal(*files)
		  if ok != nil {
					 fmt.Println("Something went wrong.")
		  }
		  writeOk := os.WriteFile(oriRoot + "/.orichalcum/log", jfile, 0666)// TODO need a geenric fucntion to fetch the path of the root ori dir
		  if writeOk != nil {
					 fmt.Println("Something went wrong.")
		  }
		  
}

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

func WhereIsOriDir(path string) string {// add error
		  pathFiles, _ := os.ReadDir(path)

		  length := len(pathFiles)

		  for i := 0; i < length; i++ {
					 if pathFiles[i].Name() == ".orichalcum" && pathFiles[i].IsDir() {
								return path
					 }
					 //Could simulaniotusly check for ori dirness
					 //if pathFiles[i].Name() == "home" && pathFiles[i].IsDir() { // not sure how to best stop searhcing up, will probably be a config
					//			return false
					// }
		  }

		  newPath := strings.Join([]string{path, ".."}, "/")
		  return WhereIsOriDir(newPath) 
}



//is there an orichalcum dir in the sub directories of a dir
func ContainsOriDir(path string) bool {
		  
		  return true		  
}

func InitOriDir() {
		  if IsOriDir(".") {
					 fmt.Println("This is already an orichalcum directory.")
					 return
		  }

		  ok := os.Mkdir(".orichalcum", 0777)// permissions are messed up 

		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }

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

// files structs

type OriFile struct {
		  path string // path relative to the root of the ori directory
		  datemod [3]byte // {y, m, d}
		  timemod string // time modified

		  dateCreated [3]byte // same as date init dir if the file is older than the creation of the ori dir
		  timeCreated [3]byte // hour/ minute / seconds

		  isArc bool // file has auto archive option enables
		  archive *OriArc
} 

type OriArc struct {
		  isRED bool // the archived copies are saved as Redundate Error-protected Digital copies
		  isTotal bool // every time the file is changed a copy is saved to a vault
}

// per file hash + object create/update

// per dir create/update ???


