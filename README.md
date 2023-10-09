![alt text](./assets/logo.png "Orichalcum")
# Orichalcum - The Meta-Filesystem
*Minimally Obtrusive File Organisation*

An axiomatic way to reason about manual file organization.

Welcome to Orichalcum! Orichalcum is a simple tool made to "manually" track and organize files. It simplifies archving copies of files, keeping backups of repos and tracking metadata across filesystems. 

## Quick-star Installation

*Don't.* Orichalcum is a simple tool and will be availble soon, but is not nearly production ready. But... if that won't stop you, a quick `git clone && cd orichalcum/ && go install` should get you where you're going!

## Feature tracker

* File Tracking
  * Normal
  * Static
  * Backup
* Meta
  * Tags
    * ID tags
    * Directory tags
* Functionality
  * Duplicate search
  * Symbolic links
* Archive
  * Generations
  * In repo
  * Non default
* Repo sync
  * FS Sync
    * Normal
    * Index
    * Endpoint
  * OverLAN
  * Over Network
    * Single Repo listen
    * Orichalcum as a Net Service
* Redundant Error-resistance
* BLOOB
  * Create/Restore
  * Encrypted BLOOB
  * BLOOBBEs

## File Tracking

Please note that when we update a repo, which is to say update the infromatoin orichalcum logs with the file system information, we refer to this as a *self sync*. Where as when we synchronize two copies of a same repo together, we call this a sync.

### Normal Repos

The normal repo is, as it sounds, the default tool of Orichalcum, it will track files and keep informaiotn about itself and its files.

### Static Repos

A repo can be set to static. Static repos does not *allow* modificaiton to files, which is to say, that on repo update, it will track new files, untrack deleted files, but if a previously tracked file has a different hashsum it will give a warning.

NOTE: Static repos where originally a different sort of object. Now, all repos will be the same (as far as internal logic is concerned), with local configs. A static repo is jsut a normal repo with the config to set all newly tracked files to static.

#### Backup

In the future, there is consideratoin for creating what we call a backup repo, which is a copy of a non static repo which ehave like a static destination for syncing to.

At this point, for the sake of symplicity, this will remain an idea and nothing more.


## Meta

Part of tracking files involves tracking metadata. This includes tracking the metadata of the filesystem, as well as custom orichalcum metadata

### Tags

Tags are the general purpose way to strucutre any infromation about the file. This makes metadata extensible, without needing to accomodate new fields in the struct one could add any number of medata fields.

Tags are organized as nested slices of strings. The first element is always a meta-meta-tag, showing what that sub slice's elements represent. The following are established meta-tags:

|Meta-Tag|Description|
|-|-|
|id|Identiy tags: distinguis files from duplicates and track them as they move around the file systems|
|dir|Direcroties where a file is located, absolute from ori root|
|opt|Options regarding how the file is treated by orichalcum, backs the local config.|


#### Identity

In orichalcum, when a file is tracked it is given an ID, a random UUID which will give that file the notion of "uniqueness". The main use for this ID is to keep the same tracking entry regardless of if the file is moved: Originally the traching ID was the path name, in which case a renamed or moved file would be untracked and retracked, this way it is not.

#### Directories

Directories are added to the tags. This allows searching by directory or having operations over directories, etc.

i.e. [oriroot]/ideas/programming/awesomeapp.md gets tracted and taged with "/", "/ideas/", "/ideas/programming/"

## Functionality

### Symbolic link

Orichalculm has a built in form of symbolic link. Each file has a slice of alternative paths, these can be actuated in different ways:

* Populate : The file is copied to all aternative directories
* Populate placeholders : empty files with the names or the alternate paths are placed across the repo.
* Populate directory : for a given directoy, expand the symlinked content.. ie. copy the files that refer to that directory
* Depopulate working directory
* Pull duplicates : for a given file, pull all duplicates to being symlinks of the current file.

Reverse

NOTE:: On sync, if there is a real file with the name of a symbolic link, a warning should be set, therefor population whould need to function acordingly. Or perhapse a seperate list of "populated" placeholders

## Archives

While backups are "remote" copies of repos, archvies are copies of files, snapshots at certain points in time.

Archvies are automatically remaned to the filename prepended by the the date in unix time. For manual archives, it is possible to anotate it with a filename appropriate string, which will be between the date and the original filename.

### Generations

Files tracked by orichalcum have a generations number, for that file, each time its orichalcum data is updated a version is caches in the hidden `.orichalcum` directory equal to the number of generations it is set to save.

i.e. If a file (file.txt) has `generation = 2` and is tracked, a file named file.txt.gen1 will be created in the `.orichlacum/generations/<file stucture mirror>` dir. Then, when the file is self-synched again, assuming it has changed, the file `file.txt.gen1` is renamed to `file.txt.gen2` and the current state of the file is coppied to `file.txt.gen1`. From then on, whenever the file is self-synched and has changed, the file `file.txt.gen2` is replaced by `file.txt.gen1` and the current state is copied to `file.txt.gen1`.

## File Syncinc

One core feature of Orichalcum is to synchronize different copies of the same repo across the same or multiple file-systems, without the need for a WAN connection or associated cloud service.

### Syncing bewteen same vs different repos

Syncing two repos genereally requires the repos to be two copies of the same repo (i.e. have the same UUID). The _only_ exceptions to this are the index and backup repo types, which will have a different ID and a non-null auxiliary ID, htis aux ID msut match the ID of the normal repo being .

### Sync Behaviours

An important concept is that in every sync operation there is a *to* and a *from* repo.

No matter what, the following holds true:

* If a File is different in each repo, the file in the _to_ repo is replaced by the file in the _from_ repo.
* If a file is present in the _from_ directory but not the _to_ repo, it will be copied to the _to_ repo.

But the following is configurable:

* If a file is present in the _to_ repo but not the _from_ repo it can either:
  * left as is
  * deleted from the _to_ repo
  * left and synced back into the _from_ repo (this is still in deliberation, as it might lead to some confusing results,)

#### Endpoints

Endpoints are locations in the file system where the OriRepo can automatically back itself up. In practice, this is jsut a sync to a copy of itself, but if it is set up to sync to locations on external storage devices, for exampls, or a NAS folder, when the repo syncs itself, it will also update the backup.

TODO: For now, the back up sync works the same way as a normal, sync, where the to repo will be selfsynced, this is not necessiary for an actual backup system, and so eventually the backup repo type will be put in place.

### Over LAN

Ori repos can talk to each other over local area networks.

#### Listening Repo

A repo can be listening, this is to say, a port on the system is being monitored. When you set a repo to listening, the port it monitors will be displayed. When you want a repo to talk to another you need to specify which port (thoughtI'd like not to, to be able to identiy all listening repos and diferentiate based on hostnames)

## Archives and Backups

Important to note, there are 2 kinks of "back-ups": file archives and repo backups.

### File Archive

A file archive is a simply a copy of a file made at specific times. This files can be copied into the archive's dir, they can be coppied to specific directories created by the user or they can be external from the ori repo.

### Repo Back-ups

A repo back-up is simply a syncing done with similar automation to the file archiving. Multiplce historical versiosn of the repo can't be saved, but if individual files ar archvied, the information is preserved. It is recomended to use a Back-up repo for this (though I'm begging to struggle to see the advantage of the back up repo... static couls still be useful though... this could all be done in the repo config) 

#### Back-up endpoints

To make back-up easy, the location of back-up repos can be included as endpoints, these can be on the local system, on external drives or network storage, as long as your systems sees it as a directory.

## RED

Setting a file to "RED" means a copy of that file will be made to the vault dir which has been redundancy hardened against corruption. This essentially equates to saving multiple copies of the file in memory. While not enterily useful if a back-up is in place it does protect against corruption of a local copy.

Originally, it was inteded that archive files would have the option of being RED copies, but there are few cases where this wouldn't be dramatic overkill. Instead, we limit the option to being avaiable when creating a manual archive, meaning when you create a manual archive you can choose to make it RED.

## BLOOB

A Binary Large Orichalcum OBject, or BLOOB, is a full copy of an ori repo, including the binary data of all the files, all ori logs, all vaulted and RED files, etc, as a single binary file.

### Encyrpted BLOOB

Perhapse the main reason BLOOBs could be usefull would be for small files which we want to keep backups of automatically. Havig these backups pre encryted makes it a nice little package which you can drop on any cloud storage platform without much worry. _Note: Orichalcum is not, primarily, a security application, and is not recomended for security dependent uses._

### BLOOBBEs

Binary Large Orichalcum OBject Backup Endpoins are just automated backups on selfsync. Combined with encrypted BLOOBs, this could simplify cloud backups for certain applications.

