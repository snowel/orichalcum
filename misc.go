package main

import (
	"github.com/oklog/ulid/v2"
)

// Mapify  a slice of OriFiles using absPaths as keys
//func PMapify(logList *[]OriFile) map[string]*OriFile {
//}


// Mapify a slice of OriFiles using ID as keys
func (log OriLog) Mapify() map[string]*OriFile {
	out := make(map[ulid.ULID]*Orifile)

	for _, v := range log.FileEntries {
		out[v.ID] = v
	}

	return out
}


func LocateFile(id ulid.ULID, repo *OriLog) *OriFile {
	for _, v := range repo.FileEntries {
		if v.id == id { return id }
	}

	return nil
}
