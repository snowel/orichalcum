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
		  DefaultVaultPath string // Relative to the oriRoot
		  IsStatic bool
		  ID string // Identifes, a repo. This ID should be the same for a backup of the repo. 

		  secret string// Key which identifies an ori-repo. When syncing orichalcum will check for matching repos with a pulbic key and use this secret to decrypt files sent
}


// Infomation about a given file, as tracked by Orichalcum.
type OriFile struct {
		  Path string // Path relative to the root of the ori repo.
		  Hash [sha512.Size]byte // SHA512 hash sum
		  DateMod int64
		  DateCreated int64
		  Size int64 // Size of the file in bytes

		  IsArc bool // File has auto archive option enabled or not.
		  Mode ArcMode
		  SizeChangeThresh int64 // every time the file is changed a copy is saved to a vault if the difference in size is greater than or equal to this threshold.
		  InteriorVaultDir string // Directory within the ori-repo which hosts the file's archives.
		  ExteriorVaultDirs []string // Directories outside of the ori-repo where backups are stored.

		  IsRED bool // A single copy is saved in the vault upon each file change which is RED protected.
		  REDFactor uint //How heavily is the files bits redundacified? (Factors of 8. I.e. if the factor is 10, the file is 80x the size)
}

// Informaiton about an archive. Used for setting the archive option on an OriFile entry.
type OriArc struct {
		  Mode ArcMode
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
		  Monthly
		  SizeChange// for a size change threshold - Size change is in both directions.
)

// Time periods in seconds (for unix time)
type UnixPeriod int64
const (
		  Day UnixPeriod = 86400 // TODO Days should behave based on date. One backup per calender day.
		  Week UnixPeriod = 604800
		  Month UnixPeriod = 2628000
)

