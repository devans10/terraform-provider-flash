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
