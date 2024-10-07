package main


// BIG IMPORTANT NOTE
// CURRENTLY A REPLACED FILE IS REPLACED AT FACE VALUE
// NO METADATA IS BACK SYNCED INTO THE LOG ENTRY WHICH WILL REPLACE IT
// THIS IS GENERALLY AS DESIGNED BUT COULD NEED TO BE AMENDED AS METADATA USE BECOMES MORE ADVANCED

type syncMode byte
const (
	Normal syncMode = iota
	NoAdd
	BackwardsClean
)

// --- File system sync (between mounted drives on a local machine)

// Housing funciton for syncing two orichalcum repos.
// TODO Option to enum
func SyncFSRepos(fromPath string, toPath string, option syncMode) {
	// check that both paths are legitimate oriroots and Load the logs

	var (
		fromLog *OriLog
		toLog *OriLog
	)

	check, _ := IsOriRoot(fromPath)
	if check {
		LoadLog(fromPath, fromLog)
	} else {
		return
	}// TODO else err
	check, _ = IsOriRoot(toPath)
	if check {
		LoadLog(toPath, toLog)
	} else {
		return
	}// TODO else err
	
	// This whoulsn't be needed, as I can set that in the main function... \Put both up to date (self sync)
	
	// loop through the list of files in the from dirrectory and make the necessary comparisons
	switch option {
	case Normal: syncLoop(fromLog, toLog, fromPath, toPath, normalDual, normalForward, normalBack)
	//case NoAdd: syncLoop(fromPath, toPath, NoAdd, NoAddForward, NoAddBack)
	}
	// Write the log TODO - so the long should be passed as apointer to the functions so that it can be nodified by the operation funcitons
}

// --- Funcitons for each mode's operations

// Type signatures
type syncOperator func(*OriFile, *OriFile, string, string) err
type syncSingleOperator func(*OriFile, *OriFile) err

// ----- Normal
// takes the file from from in all cases
// back Populates
// does not back remove


func normalDual(from, to *OriFile, fromDir, toDir string) int {
	// File has the same ID. We know this.

	// File has the same abspath?
	// No -> file was moved. New path is from's.
	moved := 0
	if from.path != to.Path { moved = 1 } 

	// File has the same date updated?
	// File has the same hash?
	// Do nothing (if the file hasn't been moved)
	if moved == 0 {
		copyEntry()
		return 0
	} else {
		moveFile(to, toDir, from.path)
		return 0
	}

	// File has a different hash or date
	// Update new replace file entry and file

	// Has to
}

// If the file only exists in from repo
func normalForward(from, to *OriFile, fromDir, toDir string) int {
	// Create the file in the to directory
	// Copy the log
}
// If the file only exists in to repo
func normalBack(from, to *OriFile, fromDir, toDir string) err {
	// Create the file in the from repo
	copyFile()
	// Copy the log
}

// Master function passing the above fucniton to each
// ForwardOp is if the file only exists in from, back is if it only exists in to
func syncLoop(fromLog, toLog *OriLog, fromRepo, toRepo string, op syncOperator, forwardOp, backOp syncSingleOperator) {
	//Load slices
	// Mapify slices
	fromMap :=  fromLog.Mapify()
	toMap := toLog.Mapify() 
	// Has to be evaluated for both id and path, because same paths will overrite even if Orichalcum can distinguish them
	for k, v := range fromMap {
		if toMap[k] != nil { // Contained in both
			op()
		} else { // File only in source repo
			forwardOp()
		}
	}

	// For all files in the destination repo not found in the home repo
	for k, v := range toMap {
		if fromMap[k] == nil { backOp() } // Matches are ignored as they've already been treated... better to remove them as they're treated the first time.
	}
}

// from exists, to does not
// from does not exist, to does

// Sum type for status between two tracked files
type FileDif byte
const (
	// For a given pair of files on the system, determin the context:
	Identical FileDif = iota// They have the same date and hash
	DifHash // They have the same date but different hash
	FromLastUpdated // They have different hashs and differnet dates, from was changed most recently
	ToLastUpdated // They have different hashs and differnet dates, to was changed most recently
	FromLastUpdatedSameHash // They have the same hash but differnet dates, from was changed most recently
	ToLastUpdatedSameHash // They have the same hash but differnet dates, to was changed most recently
)

func EvalFileSync(from string, to string) FileDif {
	if from.Hash == to.Hash && from.DateUpdated == to.DateUpdated { return Identical } // As noted above, this currently doesn't combine any meta data potentially added by orichalcum.
	if from.Hash != to.Hash && from.DateUpdated > to.DateUpdated { return FromLastUpdated }
	if from.Hash != to.Hash && from.DateUpdated < to.DateUpdated { return ToLastUpdated }
	if from.Hash != to.Hash && from.DateUpdated > to.DateUpdated { return FromLastUpdatedSameHash }
	if from.Hash != to.Hash && from.DateUpdated < to.DateUpdated { return ToLastUpdatedSameHash }
	if from.Hash != to.Hash { return DifHash }
}

// --- Atomic file operations (basic operation of moving files, copying files into a repo or removing files from a repo etc)

// Copy a file from one repo to another, replaces.
// DOES NOT remove a copy of the file in a sperate location.
func copyFile(source *OriFile, dest *OriLog, sourceRootpath, destRootPath string) err {
	// Copy the file
	// Copy the entry
}

// repo path is the path to the ori root from the working directory
// repopath + orifile.path == path to that file
// new path is abspath withing repo
// Move a file within a repo
func moveFile(file *OriFile, repoPath string, newPath string)

// Delete file from repo and untrack
func deleteFile
