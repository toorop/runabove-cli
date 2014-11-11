runabove-cli
============

Command-line tool to interact with Runabove services.

The aim is not to provide another openstack cli client but an easy to use tool to interact with Runabove services. You don't have to read tons of docs and understand how openstack works to use this tool. You'll just have to use your familiar commands like cp, ls, or rm. 

![Demo](http://dl.toorop.fr/softs/runabove-cli/pict/runabove-cli-demo.gif)

## Todo


- [x] storage
- [ ] compute
- [ ] image
- [ ] network
- [ ] me
- [ ] docker
- [ ] ...


## How to use

#### 1 - Download executable

* [Windows](http://dl.toorop.fr/softs/runabove-cli/windows/ra.exe)
* [MacOS](http://dl.toorop.fr/softs/runabove-cli/macos/ra)
* [Linux](http://dl.toorop.fr/softs/runabove-cli/linux/ra)


#### 2 - Create a Runabove account

If you don't have a Runabove account, create one using this link:

[Create Runabove account](http://runabove.me/N5SJ) (you will get 10$ free credit)

#### 3 -  Get a consumer key
In order to access to your account, the app need your authorization. You have to get "a consumer key" and record it as an environement variable.

"Keep calm and carry on", the app will do the job for you, just run it and follow instructions :

On Linux and MacOs run app with :

	./ra
	
On windows with :

	./ra.exe

## Avalaible commands

runabove-cli have an embeded help, at any moment get help by adding the --help flag.

Examples :

	$ ./ra --help
	
	NAME:
 	  ra - runabove-cli (aka ra) brings Runabove services to the command line.

	USAGE:
   	ra [section] [command] [arguments]

	SECTIONS:
   	compute	Manage runabove instances (Nova)
   	image	Manage images (Glance)
   	network	Manage networks (Neutron)
   	storage	Manage objects storage (Swift)
   	me		Manage Runabove user account
   	help, h

	OPTIONS:
   	--help, -h		show help
   	--version, -v	print the version
   	
Or:
   
	$ ./ra storage  --help
	NAME:
   	ra storage - Manage objects storage (Swift)

	USAGE:
   	ra storage command [command options] [arguments...]

	COMMANDS:
   	list, ls	Lists a path
   	copy, cp	Copies SRC to DEST. SRC & DEST can be a local path or a remote path. Remote Path (on Runabove) must start with the storage region (ex /SBG-1/images/oles-naked.jpg).
    remove, rm	Remove PATH (container,object or folder). PATH must start with the storage region (ex /SBG-1/images/oles-naked.jpg).
    help, h

	OPTIONS:
   	--help, -h	show help  
  	
 For now only storage part is done, you can :
 
### list	(ls)
Lists remote path via the "list" command 

Example:
	
	$ ./ra storage ls SBG-1/projects/runabove-cli
	0    application/octet-stream   6      2014-11-11T15:11:26Z   194577a7e20bdcc7afbb718f502c134c   .DS_Store
	0    vfolder                    17     2014-11-11T15:11:26Z                                      .git
	0    application/octet-stream   0      2014-11-11T15:11:36Z   a7fb3e6ac6b725d9577e4d3dbd9b0ab7   .gitignore
	0    application/octet-stream   1      2014-11-11T15:11:37Z   a0f6e89bc2edc26fe185fb73ec240454   LICENSE
	0    application/octet-stream   0      2014-11-11T15:11:37Z   0a74b02de55f30d801c4a7d916213d2e   README.md
	0    application/octet-stream   0      2014-11-11T15:11:49Z   c906295b5d885999f1a6a5882f6b2b13   compute.go
	0    application/octet-stream   0      2014-11-11T15:11:38Z   73207c35e42df11fe6673b2c1e024e51   const.go
	0    vfolder                    10     2014-11-11T15:11:39Z                                      dev
	0    application/octet-stream   0      2014-11-11T15:11:41Z   18aba86e6b7e4fefb18ac764db958c85   errors.go
	0    application/octet-stream   5      2014-11-11T15:11:41Z   0bba6b48160fb27a9e8f2361e9afe247   objectStorage.go
	0    application/octet-stream   7261   2014-11-11T15:11:42Z   6bad42ca722c71d1b716b32b011cf3e7   ra
	0    application/octet-stream   4      2014-11-11T15:11:44Z   b5772e0ea1afce8ae695f4e1d707cf0a   ra.go
	0    application/octet-stream   0      2014-11-11T15:11:44Z   fde7e701b0b2573d3966551c175870c3   util.go

 
### copy (cp)
With the "copy" command you can recursively download path form Runabove storage to your disk, or upload from your local storage to Runabove.

Example for downloading:

	./ra storage cp /SBG-1/projects/runabove-cli /home/toorop/
	
Or for uploading:

	./ra storage cp /home/toorop/project/ /SBG-1/projects/	

### remove (rm)
Recursively remove path and file from Runabove storage

Example:
	
	$ ./ra storage rm /SBG-1/projects/
 	
	 	



## Need a specific Runabove tool/developpement ?
Please feel free to contact me: toorop@gmail.com

## Donate
 
If you find this helpful, i encourage you to
donate what you believe would have been a fair price :

BTC Address: 18tkF3NLDt64uCSkaqr655ANJLZL1gESWw

![Donation QR](http://dl.toorop.fr/pics/btc-address-github.png)