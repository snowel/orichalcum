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
		  var rootDir string
		  if IsOriDir(".") {
					 rootDir = WhereIsOriDir(".")
					 LoadLog(rootDir, &thisLog)
		  } else {
					 InitOriDir()
		  }

		  fmt.Println(rootDir)

		  fmt.Println("Before update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&thisLog.FileEntries)
		  fmt.Println(rootDir)
		  UpdateOriDir(rootDir, &thisLog)
		  fmt.Println("After update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&thisLog.FileEntries)

		  WriteLog(&thisLog, ".")

		  files := RecursiveWalk(rootDir, len(rootDir))
		  fmt.Println(files)
		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}


// JSON read
func LoadLog(oriroot string, logHandle *OriLog) {
		  fileBytes := OpenFile(oriroot + "/.orichalcum/log")
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

// Check if a directory is aprt of an ori repo.
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
func WhereIsOriDir(path string) string {
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

// Recursive walk returning every file under the ori-root in pairs.
// Both the path relative to the working dir and the absolute path from the working dir "."
func RecursiveWalk(rootdir string, pathDif int) []PathPair {
		  filesAndDirs, ok := os.ReadDir(rootdir)
		  if ok != nil {
					 fmt.Println("Something went wrong...")
					 return []PathPair{}
		  }

		  var files []PathPair
		  length := len(filesAndDirs)
		  for i := 0; i < length; i++ {
					 if filesAndDirs[i].IsDir() {
								newRootDir := []string{rootdir, filesAndDirs[i].Name()}
								files = append(files, RecursiveWalk(strings.Join(newRootDir, "/"), pathDif)...)
					 } else {
								path := strings.Join([]string{rootdir, filesAndDirs[i].Name()}, "/")
								absPath := "." + path[pathDif:]
								pathPair := PathPair{path: path, absPath: absPath}
								files = append(files, pathPair)
					 }
		  }
		  return files
}

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

func UpdateOriDir(dir string, logHandle *OriLog) {
		  files := RecursiveWalk(dir, len(dir))
		  length := len(files)
		  for i := 0; i < length; i++ {
					 UpdateFileEntry(&(logHandle.FileEntries), files[i])
		  }
}

// create/update file entries
// The path is the actual path from the working directory to the file
// the absPath is the absolute path from the oriroot to the file, the one used as a name in the log
func TrackFile(fSlice *[]OriFile, path PathPair) {//TODO error
		  newEntry := OriFile{
								Path: path.absPath,
								Hash: HashFile(path.path),
								DateCreated: time.Now().Unix(),
								DateMod: time.Now().Unix(),
								}

		  *fSlice = append(*fSlice, newEntry)
}

func UpdateTrackedFile(entry *OriFile, path PathPair) {
		  if entry.Hash != HashFile(path.path) { 
					 entry.DateMod = time.Now().Unix()
					 entry.Hash = HashFile(path.path)
		  }
		  /*
		  if *entry.IsArc == true {
					 ArcFile(entry.Path, entry.Archive)
		  }
*/
}

func UpdateFileEntry(fSlice *[]OriFile, path PathPair) {
		  length := len(*fSlice)

		  for i := 0; i < length; i++ {
					 if (*fSlice)[i].Path == path.absPath {
								UpdateTrackedFile(&(*fSlice)[i], path)
								return
					 }
		  }

		  TrackFile(fSlice, path)
}

// Remove a file entry from the log.
func UntrackDeletedFiles(fileList []string, handle *OriLog) {

		  entry := &handle.FileEntries

		  length := len(*entry)
		  for i := 0; i < length; i++ {
					 if !Contains(fileList, (*entry)[i].Path) {
								*entry = append((*entry)[:i], (*entry)[i+1:]...)
					 }
		  }
}

// file structs

type PathPair struct {
		  path string // The relative path from the working directory. Needed for accessing the file (i.e. to hash)
		  absPath string // The absolute path from the oriroot. Used to name the file in the log.
}

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
		  Type OriKind
		  ID string // Identifes, a repo. This ID should be the same for a backup of the repo. 

		  secret string// Key which identifies an ori-repo. When syncing orichalcum will check for matching repos with a pulbic key and use this secret to decrypt files sent

}

type OriKind int // The kind of orichalcum dir.
const (
		  Normal OriKind = iota
		  Static // A Directory of files which don't or aren't supposed to change.
		  Backup // A Backup ori repo is not meant to have any edits made within it, but only accept edits from syncing another oridir.
)

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
		  Manual
)

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
