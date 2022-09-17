package main

import (
		  "fmt"
		  "os"
		  "log"
		  "strings"
		  "time"
		  "math"
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

// Check if a directory is aprt of an ori repo.
func IsOriDir(path string) bool {
		  pathFiles, _ := os.ReadDir(path)

		  length := len(pathFiles)

		  for i := 0; i < length; i++ {
					 if pathFiles[i].Name() == ".orichalcum" && pathFiles[i].IsDir() {
								return true
					 }
					 //TODO not sure how to best stop searhcing up, will probably be a config
					 if pathFiles[i].Name() == "home" && pathFiles[i].IsDir() {
								return false
					 }
		  }

		  newPath := strings.Join([]string{path, ".."}, "/")
		  return IsOriDir(newPath)
}

// Returns the path to the root of the Orichalcum directory.
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
		  // Check that the current working-dir is not part of an ori-repo.
		  if IsOriDir(".") {
					 fmt.Println("This is already an orichalcum directory.")
					 return
		  }

		  // Make the hidden folder.
		  ok := os.Mkdir(".orichalcum", 0777)// permissions are messed up 
		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }
		  // Make the Vault folder
		  ok = os.Mkdir(".orichalcum/vault", 0777)
		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }

		  // Make the RED Folder
		  ok = os.Mkdir(".orichalcum/red", 0777)// permissions are messed up 
		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }


		  // Create the initial log with metadata on the repo.
		  var Initlog OriLog

		  Initlog.Meta.DateInit = time.Now().Unix()
		  Initlog.Meta.DefaultVaultDir = ".orichalcum/vault/" //TODO config file

		  // Write the initial log
		  WriteLog(&Initlog, ".")

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


func IsTracked(absolutePath string, ori *OriLog) int {
		  length := len(ori.FileEntries)

		  for i := 0; i < length; i++ {
					 if ori.FileEntries[i].Path == absolutePath {
								return i
					 }
		  }

		  return -1
}

func PrintTrackedFiles(ori *OriLog) {
		  length := len(ori.FileEntries)

		  for i := 0; i < length; i++ {
					 fmt.Println(oir.FileEntries[i].Path)
		  }
}

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
					 oldSize := entry.Size
					 oldEdit := entry.DateMod
					 entry.DateMod = time.Now().Unix()
					 entry.Hash = HashFile(path.path)
					 entry.Size = FileSize(path.path) // TODO When i implement arc, this has to be updated later, as the size is important to size chance arhcive 
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
		  DefaultVaultPath string // Relative to the oriRoot
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

// Infomation about a given file, as tracked by Orichalcum.
type OriFile struct {
		  Path string // Path relative to the root of the ori repo.
		  Hash [sha512.Size]byte // SHA512 hash sum
		  DateMod int64
		  DateCreated int64
		  Size int64 // Size of the file in bytes

		  IsArc bool // File has auto archive option enabled or not.
		  Mode ArcMode
		  IsRED bool // A single copy is saved in the vault upon each file change which is RED protected.
		  REDFactor uint //How heavily is the files bits redundacified? (Factors of 8. I.e. if the factor is 10, the file is 80x the size)
		  SizeChangeThresh uint // every time the file is changed a copy is saved to a vault if the difference in size is greater than or equal to this threshold.
		  InteriorVaultDir string // Directory within the ori-repo which hosts the file's archives.
		  ExteriorVaultDirs []string // Directories outside of the ori-repo where backups are stored.
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
		  Total // Every time the file is changed a copy is saved to the vault
		  Daily
		  Weekly
		  Montly
		  SizeChange// for a size change threshold - Size change is in both directions.
		  Manual
)

type UnixPeriod int64
const (
		  Day UnixPeriod = 86400 // TODO Days should behave based on date. One backup per calender day.
		  Week UnixPeriod = 604800
		  Month UnixPeriod = 2628000
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

// UI

// Print the information tracked about one file.
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

func PathpairContains(slice []Pathpair, elem string, path int) bool {
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

// Archiving

func WriteArchiveCopy(entry *OriFile)// naimg will automate to ad _yyyymmdd_hhmmss, that could actually, temporarily serve to assure it's not overwiting another copy

func ArchiveTracked(entry *OriFile, oldSize int64, previousEdit int64) {
		  switch entry.Mode {
		  
					 case None: return

					 case Daily: {
		  }

					 case SizeChange: {
								if math.Abs(entry.Size - oldSize) >= entry.SizeChangeTresh {
										  WriteArchiveCopy(entry)
								}
					 }


		  }
		  
}

// Takes an absolute path... not ideal
func SetArch(file string, oriroot string, ori *OriLog) {
		  index := !IsTracked(file)
		  var answer string
		  if index == -1 {
					 fmt.Println("This file is not tracked! Would you like to see a list of tracked files? y/N")
					 fmt.Scanln(&answer)
					 answer = strings.ToLower(answer)
					 if  answer == "y" || answer == "yes" {
								PrintTrackedFiles(ori)
					 }
					 return
		  }

		  tracked := &(ori.FileEntries[i])

		  // Confirm action

		  // Sets the mode
		  clear := false
		  for !clear {
					 fmt.Println("Which mode? (hint: 0 = None, 1 = Total, 2 = Daily, 3 = Weekly, 4 = Montly, 5 = SizeChange, 6= Manual )")
					 var numAns int
					 fmt.Scanln(&numAns)

					 if numAns >= 6 && numAns <= 0 {
								clear = true
								continue
					 }

					 fmt.Println("That is not a valid mode.")

		  }
		  tracked.ArcMode = ArcMode(numAns)
		  if numAns == 6 {
					 fmt.Println("Please set the size change threshold (bytes):")
					 fmt.Scanln(&numAns)
					 tracked.SizeChangeThresh = numAns
		  }

		  // Sets the internal vault
		  tracked.InteriorVaultDir = ori.Meta.DefaultVaultPath

		  // Sets external vaults
}
