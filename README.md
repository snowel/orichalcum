# Orichalcum

![alt text](./assets/logo.png "Orichalcum")

*Minimally Aggressive/Obtrusive File Automation and Organisation*

Welcome to Orichalcum! Orichalcum is a simple tool made to "manually" track files. It simplifies archving copies of files, keeping backups of repos and tracking files. 

## Quick-star Installation

*Don't.* Orichalcum is a simple tool and will be availble soon, but is not nearly production ready. But... if that won't stop you, a quick `git clone && cd orichalcum/ && go install` should get you where you're going!

## Feature tracker
NOTE: Orichalcum is still very immature. Many features are implemented and fully functional but will likely be redone to match the overall codebase as the project grows.

* File Tracking
  * Normal
  * Static
  * Backup
* Archive
* Repo sync
  * Sync
    * Normal/Static/backup
    * Index
  * OverLAN
* Redundant Error-resistance

### Normal Repos

The normal repo is, as it sounds, the default tool of Orichalcum, it will track files and keep informaiotn about itself and its files.

### Static Repos

A static repo will behave largely like a normal repo, but it will throw a warning about modified files.

### Backup Repos

Backup repos are a little different, they act as containers which hold a copy of another repo, or multiple other repos, as well as their own meta-data. A Backup repo cna only be synced _to_, making it easier to

## Archives

While backups are "remote" copies of repos, archvies are copies of files, snapshots at certain points in time.

Archvies are automatically remaned to the filename prepended by the the date in unix time. For manual archives, it is possible to anotate it with a filename appropriate string, which will be between the date and the original filename.

## File Syncinc

One core feature of Orichalcum is to synchronize different copies of the same repo across the same or multiple file-systems, without the need for a WAN connection.

### Syncing bewteen same vs different repos

Syncing two repos genereally requires the repos to be two copies of the same repo (i.e. have the same UUID). The _only_ exceptions to this are the index and backup repo types, which will have a different ID and a non-null auxiliary ID, htis aux ID msut match the ID of the normal repo being .

An important concept is that in every sync operation there is a *to* and a *from* repo.

No matter what, the follwoing holds true:

* If a File is different in each repo, the file in the _to_ repo is replaced by the file in the _from_ repo.
* If a file is present in the _from_ directory but not the _to_ repo, it will be copied to the _to_ repo.

But the following is configurable:

* If a file is present in the _to_ repo but not the _from_ repo it can either:
  * left as is
  * deleted from the _to_ repo
  * left and synced back into the _from_ repo (this is still in deliberation, as it might lead to some confusing results,)

### Over LAN

Ori repos can talk to each other over local area networks.

#### Listening Repo

A repo can be listening, this is to say, a port on the system is being monitored. When you set a repo to listening, the port it monitors will be displayed. When you want a repo to talk to another you need to specify which port (thoughtI'd like not to, to be able to identiy all listening repos and diferentiate based on hostnames)

## Archive and Back-Up

Important to note, there are 2 kinks of "back-ups": file archives and repo back-ups.

### File Archive

A file archive is a simply a copy of a file made at specific times. This files can be copied into the archive's dir, they can be coppied to specific directories created by the user or they can be external from the ori repo.

### Repo Back-ups

A repo back-up is simply a syncing done with similar automation to the file archiving. Multiplce historical versiosn of the repo can't be saved, but if individual files ar archvied, the information is preserved. It is recomended to use a Back-up repo for this (though I'm begging to struggle to see the advantage of the back up repo... static couls still be useful though... this could all be done in the repo config) 

#### Back-up endpoints

To make back-up easy, the location of back-up repos can be included as endpoints, these can be on the local system, on external drives or network storage, as long as your systems sees it as a directory.



## Repo Types

### Default

A default ori repo is, as it sounds, the symplest system.

### Backup

A back-up directory is mean to have no changes made to it outside of syncing to another ori repo. A back-up repo can be listed as an back-up target for any repo and can automatically 

### Static

A static repo is eman for files which can be added moved or removed, but aren't emant to be modified. Things liek music collections, wallpaper folders and vacation photos make great candidates for static repos.

### Index

Likely dep. The index repo can be well served by an orichalcum function... Though partial syncing might eventually make it back into the system
An index repo behave significantly differently than the others, by ddefault every file in the repo will be replaced by a text file containing the ori entry of that file from the source repo. 


## R.E.D or RED

Setting a file to "RED" means a copy of that file will be made to the vaulr dir which has ben redundancy hardened against corruption. This essentially equates to saving multiple copies of the file in memory. While not enterily useful if a back-up is in place.
