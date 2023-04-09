package main

import (
		  "fmt"
		  "os"
		  "strings"
		  "time"
)


// Determins if the given, working directory is the root of an ori repo.
// Returns the bool and all other directories in ori root (for downwards search)
func IsOriRoot(working string) (bool, []string) {
		  pathFiles, _ := os.ReadDir(working)

		  length := len(pathFiles)
		  var subDirs []string
		  ans := false

		  for i := 0; i < length; i++ {
				   if pathFiles[i].Name() == ".orichalcum" && pathFiles[i].IsDir() {
								// One way switch
								ans = true
					 }
					 // collect all sub directories
					 if pathFiles[i].IsDir() {
								subDirs = append(subDirs, pathFiles[i].Name())
					 }
		  }
			return ans, subDirs
}

// Check if a directory is part of an ori repo.
func IsOriRepo(working string) bool {
		  current, other := IsOriRoot(working)

		  if current {
					 return true
		  } else if Contains(other, "home") { // Stops searching wehn reaching the root dir TODO cleaner
					 return false
		  }

		  newPath := strings.Join([]string{working, ".."}, "/")
		  return IsOriRepo(newPath)
}

// Returns the path to the root of the Orichalcum repo from the working dir.
// Previously WhereIsOriRepo
func WhereIsOriRoot(working string) string {
		  if ans, _ := IsOriRoot(working); ans {
					 return working
		  }

		  newPath := strings.Join([]string{working, ".."}, "/")
		  return WhereIsOriRoot(newPath) 
}


// Determins if the any of the subdirectories of the working directory are the roots of an ori repo. 
func ContainsOriRepo(working ...string) bool {
		  var subDirs []string
		  for _, dir := range working {

					 oriFound, partialSubDirs := IsOriRoot(dir)
					 if oriFound {return oriFound}
					 
					 subDirs = append(subDirs, partialSubDirs...)
		  }
					 if len(subDirs) == 0 {
								return false
					 } else {
								return ContainsOriRepo(subDirs...)
					 }
}

// If this is not a .orichalcum/ directory anywhere aboce the curent directory, init one with the current dir as ori-root.
func InitOriDir() {
		  // Check that the current working-dir is not part of an ori-repo.
		  if IsOriRepo(".") {
					 fmt.Println("This is already an orichalcum directory.")
					 return
		  }

		  // Make the hidden folder.
		  ok := os.Mkdir(".orichalcum", 0777)
		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }
		  // Make the Vault folder
		  ok = os.Mkdir(".orichalcum/vault", 0777)
		  if ok != nil {
					 fmt.Println("Somethign went wrong...")
		  }

		  // Make the RED Folder
		  ok = os.Mkdir(".orichalcum/red", 0777)
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
					 UpdateFileEntry(logHandle, files[i])
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

func UpdateTrackedFile(entry *OriFile, orilog *OriLog, path PathPair) {
	 newHash :=HashFile(path.path)  
	 oldSize := entry.Size
	 oldEdit := entry.DateMod
	 rel := WhereIsOriRoot(".")// TODO might want to add a "vitual working dir" for the case where I'll be syncing etc
	if entry.Hash != newHash {
				 entry.DateMod = time.Now().Unix()
				 entry.DateChanged = FileChanged(path.path).Unix() 
				 entry.Hash = HashFile(path.path)
				 entry.Size = FileSize(path.path) 
		  }
		  
		  if entry.IsArc == true {
			  AutoArchive(entry, orilog, oldSize, oldEdit, rel)
		  }
}

func UpdateFileEntry(orilog *OriLog, path PathPair) {
	 fSlice := &(orilog.FileEntries)
	 length := len(*fSlice)

	 for i := 0; i < length; i++ {
	 if (*fSlice)[i].Path == path.absPath {
			UpdateTrackedFile(&(*fSlice)[i],orilog, path)
			return
	 }
}

		  TrackFile(fSlice, path)
}

// Remove a file entry from the log.
/*
func UntrackDeletedFiles(fileList []string, handle *OriLog) {

	 entry := &handle.FileEntries

	 length := len(*entry)
	 for i := 0; i < length; i++ {
				if !Contains(fileList, (*entry)[i].Path) {
						 *entry = append((*entry)[:i], (*entry)[i+1:]...)
				}
	 }
}

*/
