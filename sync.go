package main


// Housing funciton for syncing two orichalcum repos.
// TODO Option to enum
func SyncFSRepos(fromPath string, toPath string, option byte) {
	// check that both paths are legitimate oriroots

	// Put both up to date (self sync)
	// load both logs
	// loop through the list of files in the from dirrectory and make the necessary comparisons
		//for each file, based on option, copy, replace, etc
}

// For a given pair of files on the system, determin the context:
// from exists, to does not
// from does not exist, to does
// from and to exist, they have the same date and hash
// from and to exist, they have the same date but different hash
// from and to exist, they have the same hash but differnet dates, from was changed most recently
// from and to exist, they have the same hash but differnet dates, to was changed most recently
// TODO return enum
func EvalFileSync(from string, to string) {
return
}
