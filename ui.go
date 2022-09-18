package main

import (
		  "fmt"
		  "os"
		  "log"
		  "time"
)

// UI

func PrintTrackedFiles(ori *OriLog) {
		  length := len(ori.FileEntries)

		  for i := 0; i < length; i++ {
					 fmt.Println(oir.FileEntries[i].Path)
		  }
}


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

// *** Inputs ***

//Archive
/*
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
*/
