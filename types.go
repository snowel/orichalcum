package main

import (
		  "crypto/sha512"
)

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
		  DefaultVaultDir string // Relative to the oriRoot
		  IsStatic bool
		  Endpoints string//Also only as a space seperated string, for now -- []string // List of directories to look at when syncing. If a matching Oriroot is found in any of those dirs, the repos are sync in safe mode
		  EndMode SyncMode // Mode to use with enpoint bacpups
		  // EndMode will probably not be used as we want to keep it simple to treat backups...
		  ID string // Identifes, a repo. This ID should be the same for a backup of the repo. 


		  secret string// Key which identifies an ori-repo. When syncing orichalcum will check for matching repos with a pulbic key and use this secret to decrypt files sent
}

// Infomation about a given file, as tracked by Orichalcum.
type OriFile struct {
		  Path string // Path relative to the root of the ori repo.
		  Filename string // temp
		  Hash [sha512.Size]byte // SHA512 hash sum
		  DateMod int64 // Date of last sync where the file was differnt
		  DateChanged int64 // Date of las change, according to OS FS .  currently redundant but temporarily both are used
		  DateTracked int64 // Date when the file is added to the repo
		  DateCreated int64// Date of the file creation according to OS FS 
		  Size int64 // Size of the file in bytes

		  // Meta fields are essetial information fields that are exclusive to Orichicalcum.
		  MetaTitle string // A file name but only as far as Orichalcum is concerned. Multiple files can have the same meta title. Does not suport nuewlines.
		  MetaNote string // Arbitrary info about the file, supports newliens.
		  MetaTags string//[]string // Tags used for file organizations what go over the heads of the dir structure and orichalcums base features.
		//TODO due to neededing deeep copy in Archiving we're currently seperating tages and external dirs withing a single string

		  IsStatic bool // 

		  SyncSafe bool // Each time the file would be overwritten in a sync the version compare is shown and manual confimation is required
		  AutoBacksync bool // Even when not backsyncing, a file with auto backsync will be saved in the to repo. TEMP.
		  // AutoBacksync will probably not be used.

		  IsArc bool // File has auto archive option enabled or not.
		  ArchiveMode ArcMode
		  SizeChangeThresh int64 // every time the file is changed a copy is saved to a vault if the difference in size is greater than or equal to this threshold.
		  //Size chance can double as time change as well
		  //TODO rename
		  CustomVaultDir string

		  IsRED bool // A single RED copy of each tracked file is saved in the vault upon modification.
		  REDFactor uint // How heavily is the files bits redundacified? (Factors of 8. I.e. if the factor is 10, the file is 80x the size)
}

// Informaiton about an archive. Used for setting the archive option on an OriFile entry.
type OriArc struct {
		  Mode ArcMode
		  SizeChangeThresh uint // every time the file is changed a copy is saved to a vault
		  VaultDirs []string
}
//TODO future pot: add RED arhcives

// Enum for setting the type of auto-archiving.
type ArcMode int
const (
		  None ArcMode = iota
		  Total // Every time the file is changed a copy is saved to the vault
		  Daily
		  Weekly
		  Monthly
		  SizeChange// for a size change threshold - Size change is in both directions.
)


// Enum of sync modes.
type SyncMode byte
const (
	Safe SyncMode = iota
	// from files replace to files if different AND change date of from file is more recent than to file
	// 	If to file is more recent, manul confirmation is asked
	// from files are added if not found in to
	// to files that are not found in from are left untouched
	
	Standard
	// Like safe, but doesn't check date changed


	cleaning
	// from files replace to files is different
	// from files are added if not found in to
	// to files that are not found in from are deleted and untracked

	backsync
	// from files replace to files is different
	// from files are added if not found in to
	// to files that are not found in from are added to the from repo
	
)

// TODO, add file property of safety: each time that file would be synced, manual confirmaiotn is asked

// Time periods in seconds (for unix time)
type UnixPeriod int64
const (
		  Day UnixPeriod = 86400 // TODO Days should behave based on date. One backup per calender day.
		  Week UnixPeriod = 604800
		  Month UnixPeriod = 2628000
)

