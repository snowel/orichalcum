package main

import (
		  "fmt"
		  "strings"
		  "time"
)

// UI

// Orichalcum options
// none
//   self syncs repo, can be called anywhere inside the ori dir
//   if this isn't a ori repo, it will init an ori root at working dir
// sync + 1 = sync to path
// sync + 2 =  sync from path to path
//    By defualt it won't sync if both paths aren't oriroots
// todo flags


// ---- Basic, full funciton main calls ----

// no arg
//Current;y testing
// Future:
// self sync repo
// if repo is not ori, ask for confimation then init
func SelfSync(){
		  sessionLog := OriLog{} // need this no matter what
		  var toOriRoot string // path from working dir to repo root

		  // find ori repo or create it
		  if IsOriRepo(".") {
					 toOriRoot = WhereIsOriRoot(".")
					 LoadLog(toOriRoot, &sessionLog)
		  } else {
					 InitOriDir()
		  }

		  fmt.Println(toOriRoot)

		  fmt.Println("Before update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&sessionLog.FileEntries)
		  fmt.Println(toOriRoot)
		  UpdateOriDir(toOriRoot, &sessionLog)
		  fmt.Println("After update  --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		  PrintFileLog(&sessionLog.FileEntries)

		  WriteLog(&sessionLog, toOriRoot)

		  //fmt.Println("Orichalcum.") the original, total functionality of this program's first install 
}

// Help menus :

func PrintGeHelp(){
	fmt.Println("Welcome to the general orichalcum help menu.")
	fmt.Println("")
	fmt.Println("The following option are available:")
	fmt.Println("init")
	fmt.Println("self")
	fmt.Println("sync")
	fmt.Println("help")
	fmt.Println("")
}
// Debug tools
func PrintTrackedFiles(ori *OriLog) {
		  length := len(ori.FileEntries)

		  for i := 0; i < length; i++ {
					 fmt.Println(ori.FileEntries[i].Path)
		  }
}


// Print the information tracked about one file.
func PrintFileEntry(handle *OriFile) {
		  fmt.Println(handle.Path)
		  dateCreate := time.Unix(handle.DateCreated, 0)
		  dateMod := time.Unix(handle.DateMod, 0)
		  fmt.Println("File created:  ", dateCreate)
		  fmt.Println("File modified:  ", dateMod)
		  fmt.Println("Is this file archived:", handle.IsArc, handle.ArchiveMode)
		  fmt.Println("File RED:  ", handle.IsRED)
}

func PrintFileLog(handle *[]OriFile) {
		  length := len(*handle)
		  for i := 0; i < length; i++ {
					 PrintFileEntry(&(*handle)[i])
		  }
}


// Display infromation about two (of the same) files side by side. Used for sync compares.
func PrintVersionCompaire(from *OriFile, to *OriFile){
	fmt.Println(" VS ")
}

// Option alias check -- chekcs if a given string arg is equivalent to a slice of 
func IsAlias(arg string, aliases []string) bool {

	command := strings.ToLower(arg)

	for _, v := range aliases {
		if command == v {return true}
	}

	return false
}

// *** Semantic aliases for positional args ***

var (
	yes = []string{"yes", "y", "yeah", "duh", "obviously", "true", "1", "affirmative"}
	no = []string{"no", "n", "nah", "nope", "never", "false", "0", "negative"}
	sync = []string{"sync", "synchronize", "s", "makesame"}
)

// *** Inputs ***

//Archive
/*
// Takes an absolute path... not ideal
func SetArch(file string, oriroot string, ori *OriLog) {
		  index := !IsTracked(file)4
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
*/
