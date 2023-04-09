package main


import (
		"os"
		"log"
		"time"
)
// Archiving

// Create an archive copy of a file
// DIr must be absolute path from Oriroot
// rel is the path from working to oriroot
func WriteArchiveCopy(entry *OriFile, orilog *OriLog, rel string){
	// Get filename
		  file, ok := os.Stat(rel + entry.Path)
		  if ok != nil {
					 log.Fatal(ok)
		  }
		  filename := file.Name()
	// Open file
		f, ok := os.ReadFile(filename)
		  if ok != nil {
					 log.Fatal(ok)
		  }
	// Write the file to to the Dir
	
	//  Get the dir, whether custom or nah

	var dir string
	if entry.CustomVaultDir != "" {
		dir = entry.CustomVaultDir
	} else {
		dir = orilog.Meta.DefaultVaultDir
		// TODO add check to see tha tthe dir is actually there
	}
	archiveDir := dir + time.Unix(entry.DateMod, 0).Format("2006-01-02_15-04-05") + "_" + filename 
	//TODO time should be last modified
		  ok = os.WriteFile(rel + archiveDir, f, 0666)
		  if ok != nil {
					 log.Fatal(ok)
		  }
	// track file with the copied entry log struct + add orichalcum_archive_file tag
	arcLog := entry // this doesn't work as the metatags are a slice and need deep copy
	arcLog.MetaTags = arcLog.MetaTags + " orichalcum_archive_file"
	// TODO temporarily using only literals in the struct to be able to do it the lazy way...
	orilog.FileEntries = append(orilog.FileEntries, *arcLog)
}
// namimg will prepend yyyymmdd_hhmmss
//that could actually, temporarily serve to assure it's not overwiting another copy


// Manual write an archive

// Check auto archiving properties and creates an archive copy if one of the qualities match.
// TODO Auto archiving only supports default archive location for now
// TODO Auto archiving only supports one archive location, probalby long term
func AutoArchive(entry *OriFile, orilog *OriLog, oldSize int64, previousEdit int64, rel string) {
		  switch entry.ArchiveMode {
		  
					 case None: return

					 case Daily: {
						 if i, _ := changeCompaire(entry.DateMod, previousEdit, int64(Day)); i {
										  WriteArchiveCopy(entry, orilog, rel)
								}
					 }
					 case Weekly: { 
						 if i, _ := changeCompaire(entry.DateMod, previousEdit, int64(Day)); i {
										  WriteArchiveCopy(entry, orilog, rel)
								}
					 }
					 case Monthly: {
						 if i, _ := changeCompaire(entry.DateMod, previousEdit, int64(Day)); i {
										  WriteArchiveCopy(entry, orilog, rel)
								}
					 }

					 case SizeChange: {
						 if i, _ := changeCompaire(entry.Size, oldSize, int64(entry.SizeChangeThresh)); i {
										  WriteArchiveCopy(entry, orilog, rel)
								}
					 }


		  }
		  
}

// Reutrns true if the difference between the two is creater than the tresh.
// As time and size are both measured as int64, this works for both.
// returns a bool of if the difference is greater and the difference if it's possitive the file grew by that amount, else degreses
func changeCompaire(current int64, prev int64, tresh int64) (bool, int64) { 
		  dif := current - prev
		  
		  if dif >= tresh || dif <= -tresh {
					 return true, dif
		  } else {
					 return false, dif
		  }

}
/*
// Get the archiving mode from the user, manually.
func QueryArc(handle *OriArc, filelog *OriFile) {
		handle.IsArc
		handle.ArchiveMode
		handle.SizeChangeThresh
		 //CustomVaultDirs
}

// Set the archiving mode for a file
func SetArc(entry *OriFile, settings *OriArc) {

		  entry.ArchiveMode = settings.Mode
		  entry.ExteriorVaultDirs = settings.VaultDirs
		  entry.SizeChangeThresh = settings.SizeChangeThresh
}

// Manual set ori arc
func MSetArc(entry *OriFile) {
	settings := new(OriArc)
	QueryArc(settings)
	SetArc(entry, settings)
}
*/
