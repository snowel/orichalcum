package main

import (
		  "fmt"
		  "os"
		  "log"
		  "strings"
		  "time"
)

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
		  }

		  newPath := strings.Join([]string{path, ".."}, "/")
		  return WhereIsOriDir(newPath) 
}


// Determins if the given, working directory is the root of an ori repo.
func IsOriRoot(working string) bool, []string {
		  pathFiles, _ := os.ReadDir(wokring)

		  length := len(pathFiles)
		  var subDirs []string

		  for i := 0; i < length; i++ {
					 if pathFiles[i].Name() == ".orichalcum" && pathFiles[i].IsDir() {
								return true, subDirs
					 }
					 // collect all sub directories
					 if pathFiles[i].IsDir() {
								subDirs = append(subdirs, pathFiles[i].Name())
					 }
		  }
		  return false, subDirs
}

// Determins if the any of the subdirectories of the working directory are the roots of an ori repo. 
func ContainsOriRoot(working ...string) bool {
		  var subDirs []string
		  for _, dir := range working {

					 oriFound, partialSubDirs := IsOriRoot(dir)
					 if oriFount {return oriFound}
					 
					 subDirs = append(subDirs, partialSubDirs...)
		  }
					 if len(subDirs) == 0 {
								return false
					 } else {
								return ContainsOriRoot(subDirs...)
					 }
}

// If this is not a .orichalcum/ directory anywhere aboce the curent directory, init one with the current dir as ori-root.
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

