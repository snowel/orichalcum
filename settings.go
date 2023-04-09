package main

import (
	"fmt"
)

// Set Repo ID
func NameRepo(repo *OriMeta, name string){
	//TODO add check before ovewriiting name... though not sure it's necc as there will be chekc in ui
	repo.ID = name
}

// Set repo as static.
func MakeRepoStatic(handle *OriMeta){
	if handle.IsStatic == true {
		fmt.Println("This repository is already a static repository.")
		return
	}
	
	handle.IsStatic = true
	fmt.Println("The repository: ", handle.ID, " has been set to static.")
}
