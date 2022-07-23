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

		  thisLog := OriLog{}

		  LoadLog(".orichalcum/log", &thisLog)

		  PrintFileLog(&thisLog.FileEntries)
		  UpdateOriDir(".", &thisLog)
		  PrintFileLog(&thisLog.FileEntries)

		  WriteLog(&thisLog, ".")

		  //if IsOriDir("."){ 
		//			 fmt.Println("Ori dir!")
		 // }
		  //printFilenames(".")
		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// JSON read

func LoadLog(OriLogPath string, logHandle *OriLog) {
		  fileBytes := OpenFile(OriLogPath)
		  json.Unmarshal(fileBytes, logHandle)
}

// JSON write

func WriteLog(logHandle *OriLog, path string) {

		  oriRoot := WhereIsOriDir(path)

		  jfile, ok := json.Marshal(*logHandle)
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

// Retruns the path to the root of the orichalcum directory.
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

// If there this is not a .orichalcum/ directory anywhere aboce the curent directory, init one with the current dir as ori-root.
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

//TODO if i sync from a differnt directory the path to the root ori with bee differnt from a diferent sync...
// I need to pull out the patht to root before runnitn thought the filepaths
// i.e. rootdir needs to alway be the same insice of the Update function
func UpdateOriDir(rootdir string, logHandle *OriLog) {
		  //Open a file slice
		  filesAndDirs, ok := os.ReadDir(rootdir)
		  if ok != nil {
					 fmt.Println("Something went wrong...")
					 return
		  }

		  length := len(filesAndDirs)
		  for i := 0; i < length; i++ {
					 // if the path is a dir, recuse into the dir to keep iterating over every file
					 if filesAndDirs[i].IsDir() {
								newRootDir := strings.Join([]string{rootdir, filesAndDirs[i].Name()}, "/")
								UpdateOriDir(newRootDir, logHandle)
					 } else {
					 // If the path is a file, update the file entry in the log handle.
								path := strings.Join([]string{rootdir, filesAndDirs[i].Name()}, "/")
								UpdateFileEntry(&(logHandle.FileEntries), path)
					 }
		  }
}

// file structs

// The information which orichalcum will save in the .orichalcum directory.
type OriLog struct {
		  FileEntries []OriFile
		  Meta OriMeta
}

// Information about the orichalcum directory itself.
type OriMeta struct {
		  DateInit int64
		  DateMod int64
		  DefualtVaultPath string

		  secret string

}

// Infomation about a given file
type OriFile struct {
		  Path string // path relative to the root of the ori directory
		  Hash [sha512.Size]byte
		  DateMod int64

		  DateCreated int64

		  IsArc bool // file has auto archive option enables
		  Mode ArcMode
		  IsRED bool // the archived copies are saved as Redundate Error-protected Digital copies
		  SizeChangeThresh uint // every time the file is changed a copy is saved to a vault
		  VaultDirs []string
}

// Informaiton about an archive. Used for setting the archive option on an OriFile entry.
type OriArc struct {
		  Mode ArcMode
		  IsRED bool // the archived copies are saved as Redundate Error-protected Digital copies
		  SizeChangeThresh uint // every time the file is changed a copy is saved to a vault
		  VaultDirs []string
}

// Enum for setting the type of auto-archiving.
type ArcMode int
const (
		  None ArcMode = iota
		  Total 
		  Daily
		  Weekly
		  Montly
		  SizeChange// for a size change threshold - Size change is in both directions.
)

// per file hash + object create/update

func TrackFile(fSlice *[]OriFile, path string) {//TODO error
		  newEntry := OriFile{
								Path: path,
								Hash: HashFile(path),
								DateCreated: time.Now().Unix(),
								DateMod: time.Now().Unix(),
								}

		  *fSlice = append(*fSlice, newEntry)
}

func UpdateTrackedFile(entry *OriFile) {
		  if entry.Hash != HashFile(entry.Path) { 
					 entry.DateMod = time.Now().Unix()
					 entry.Hash = HashFile(entry.Path)
		  }
		  /*
		  if *entry.IsArc == true {
					 ArcFile(entry.Path, entry.Archive)
		  }
*/
}

func UpdateFileEntry(fSlice *[]OriFile, path string) {
		  length := len(*fSlice)

		  for i := 0; i < length; i++ {
					 if (*fSlice)[i].Path == path {
								UpdateTrackedFile(&(*fSlice)[i])
								return
					 }
		  }

		  TrackFile(fSlice, path)
}


func SetArc(entry *OriFile, settings *OriArc) {

		  entry.Mode = settings.Mode
		  entry.IsRED = settings.IsRED
		  entry.VaultDirs = settings.VaultDirs
		  entry.SizeChangeThresh = settings.SizeChangeThresh
}


// Hash utils
func OpenFile(filename string) []byte{

		  f, ok := os.ReadFile(filename)
		  
		  if ok != nil {
					 log.Fatal(ok)
		  }

		  return f 
}

// for a filepath, retrun the byte array of the hsha512 sum
func HashFile(filename string) [sha512.Size]byte{

		  f, ok := os.ReadFile(filename)
		  
		  if ok != nil {
					 log.Fatal(ok)
		  }
		  
		  hash := sha512.Sum512(f)
		  return hash 
}

//TODO this function may not be necessairy
// for a file path and an entry, compare if the file is the same or not, true == same hash
func CompareFileHash(filePath string, fileEntry *OriFile) bool {
		  newHash := HashFile(filePath) 
		  oldHash := fileEntry.Hash

		  if newHash == oldHash {
					 return true
		  } else {
					 return false
		  }
}

// Meta information

func OnUpdate(logHandle *OriLog) {
		  logHandle.Meta.DateMod = time.Now().Unix()
}

// UI

func PrintFileEntry(handle *OriFile) {
		  fmt.Println(handle.Path)
		  dateCreate := time.Unix(handle.DateCreated, 0)
		  dateMod := time.Unix(handle.DateMod, 0)
		  fmt.Println("File created:  ", dateCreate)
		  fmt.Println("File modified:  ", dateMod)
		  fmt.Println("Is this file archived:", handle.IsArc)
}

func PrintFileLog(handle *[]OriFile) {
		  length := len(*handle)

		  for i := 0; i < length; i++ {
					 PrintFileEntry(&(*handle)[i])
		  }
}
