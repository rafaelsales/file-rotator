# Street sweeper

Street sweeper is a program that monitors a directory and removes old files whenever
the storage use meets certain condition.

# Usage and examples

* All options:

   `streetsweeper --help`

* When the directory `/etc/lib/motion` has used more than 4GB,
delete oldest video files until it only has 3.5GB used.

  `streetsweeper --max-size 4000MB --target-size 3500MB /etc/lib/motion`
