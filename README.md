![alt text](./assets/logo.png "Orichalcum")
# Orichalcum
*Minimally Aggressive or Obtrusive File Automation and Organisation*

An axiomatic way to reason about manual file organization.

Welcome to Orichalcum! Orichalcum is a simple tool made to "manually" track files. It simplifies archving copies of files, keeping backups of repos and tracking files. 
Oricalcum is something like a VCS, but, rather than controlling versions, it controls state, making it possible to treat a given directory structure of files as one porta le unit.

## Quick-star Installation

*Don't.* Orichalcum is a simple tool and will be availble soon, but is not nearly production ready. But... if that won't stop you, a quick `git clone && cd orichalcum/ && go install` should get you where you're going!

## Feature tracker

* File Tracking
  * Normal
  * Static
  * Backup
* Archive
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

#### Static Repos

A repo can be set to static. Static repos does not *allow* modificaiton to files, which is to say, that on repo update, it will track new files, untrack deleted files, but if a previously tracked file has a different hashsum it will throw an error

Additionally, within static repos, on sync, a moved

#### Backup

In the future, there is consideratoin for creating what we call a backup repo, which is a copy of a non static repo which ehave like a static destination for syncing to.

At this point, for the sake of symplicity, this will remain an idea and nothing more.

## Archives

While backups are "remote" copies of repos, archvies are copies of files, snapshots at certain points in time.

Archvies are automatically remaned to the filename prepended by the the date in unix time. For manual archives, it is possible to anotate it with a filename appropriate string, which will be between the date and the original filename.

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

A Binary Large Orichalcum OBject, of BLOOB, is a full copy of an ori repo, including the binary data of all the files, all ori logs, all vaulted and RED files, etc, as a single binary file.

### Encyrpted BLOOB

Perhapse the main reason BLOOBs could be usefull would be for small files which we want to keep backups of automatically. havig these backups pre encryted makes it a nice little package which you can drop on any cloud storage platform without much worry. _Note: Orichalcum is not, primarily, a security application, and is not recomended for security dependent uses._

### BLOOBBEs

Binary Large Orichalcum OBject Backup Endpoins are just automated backups on selfsync. Combined with encrypted BLOOBs, this could simplify cloud backups for certain applications.

