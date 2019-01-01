## 0.3.8

FIXES:

* minor update to travis.yml to fix build pipeline.

## 0.3.0

IMPROVEMENTS:

* `resource/host`: Added CHAP parameters and host personality
* Added Partial States to all resources
* Updated .travis.yml to build cross-platform binaries

## 0.2.1

NOTES:

* Updated go-purestorage library to v0.1.2
* Updated documentation

## 0.2.0 (December 28, 2018)

NOTES:

* Create hostgroups with attached hosts
* Attach volumes to hostgroups using connected_volumes parameter
* Added protection group resource
* Added Importers to all resources
* Added ability to update all resource parameters
* Added some random number to names in tests, so they can be run multiple times without cleaning up
* Fixed Protection Groups not getting Hosts or Hostgroups attached upon creation.(#2)

## 0.1.0 (December 2, 2018)

NOTES:

* Initial Release. Limited functionality. Create/Update/Delete Volumes, Hosts, Hostgroups
* Update of names only at this point
* Delete of volumes does *NOT* eradicate the volume
