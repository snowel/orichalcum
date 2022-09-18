package main

import (
		  //"fmt"
		  //"os"
		  "log"
		  "strings"
		  "time"
		  "crypto/sha512"
)

// Archiving

func WriteArchiveCopy(entry *OriFile)// namimg will automate to ad _yyyymmdd_hhmmss, that could actually, temporarily serve to assure it's not overwiting another copy

// Check auto archiving properties and creates an archive copy if one of the qualities match.
func AutoArchive(entry *OriFile, oldSize int64, previousEdit int64) {
		  switch entry.Mode {
		  
					 case None: return

					 case Daily: {
								if changeCompaire(entry.DateMod, previousEdit, int64(Day)) {
										  WriteArchiveCopy(entry)
								}
					 }
					 case Weekly: { 
								if changeCompaire(entry.DateMod, previousEdit, int64(Week)) {
										  WriteArchiveCopy(entry)
								}
					 }
					 case Monthly: {
								if changeCompaire(entry.DateMod, previousEdit, int64(Month)) {
										  WriteArchiveCopy(entry)
								}
					 }

					 case SizeChange: {
								if changeCompaire(entry.Size, oldSize, entry.SizeChangeThresh) {
										  WriteArchiveCopy(entry)
								}
					 }


		  }
		  
}

// Reutrns true if the difference between the two is creater than the tresh.
// As time and size are both measured as int64, this works for both.
func changeCompaire(current int64, prev int64, tresh int64) bool {
		  dif := current - prev
		  if dif < 0 {dif *= -1}
		  
		  if dif >= tersh {
					 return true
		  } else {
					 return false
		  }

}

func SetArc(entry *OriFile, settings *OriArc) {

		  entry.Mode = settings.Mode
		  entry.IsRED = settings.IsRED
		  entry.ExteriorVaultDirs = settings.VaultDirs
		  entry.SizeChangeThresh = settings.SizeChangeThresh
}
