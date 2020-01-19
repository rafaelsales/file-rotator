# File Rotator (old name Street-Sweeper)

File Rotator is a program that monitors a directory and removes old files whenever the storage use meets certain condition.

# Usage and examples

* When the directory `/var/lib/motion` has used more than 1GB, delete oldest video files until it only has 500MB used.

  ```
  pi@raspberrypi:/var/lib/motion $ ~/streetsweeper 1GB 500MB /var/lib/motion
  2017/08/25 08:12:57 Initializing. Path: /var/lib/motion | Max size: 1073741824 | Target size: 524288000
  2017/08/25 08:12:57 Path: /var/lib/motion | Current size: 1179781545
  2017/08/25 08:12:57 Path: /var/lib/motion | Over max size - Current size: 524288425 | Max size: 1073741824
  2017/08/25 08:15:35 Write detected: /var/lib/motion/01-20170825083251.avi
  2017/08/25 08:15:35 Deleting: /var/lib/motion/01-20170807042950-01.jpg | Modified at: 2017-08-07 04:31:14.154876893 +0000 UTC | Size: 60372
  2017/08/25 08:15:35 Deleting: /var/lib/motion/02-20170807043341-01.jpg | Modified at: 2017-08-07 04:34:57.175492342 +0000 UTC | Size: 57898
  2017/08/25 08:15:35 Deleting: /var/lib/motion/03-20170807045241-00.jpg | Modified at: 2017-08-07 04:53:46.317611474 +0000 UTC | Size: 62855
  ...
  2017/08/25 08:15:35 Deleting: /var/lib/motion/334-20170815154704-06.jpg | Modified at: 2017-08-15 15:48:04.145820372 +0000 UTC | Size: 83145
  2017/08/25 08:15:35 Deleting: /var/lib/motion/335-20170815154849.avi | Modified at: 2017-08-15 15:49:49.075832118 +0000 UTC | Size: 756396
  2017/08/25 08:15:35 Deleted a total of 65575992
  ```
