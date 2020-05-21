# apclog
Apache log generator using golang

# Prerequisites
#  go  - Go Language installed and the binary location is included in the PATH variable
#  git - Required for downloading dependent repositories


# Installation
# Once the repository is downloaded to the system, please run the following commands
$ cd apclog
$ go build
#
# Verify by running
$ ./apclog -h
#

# Generate the apache logs
# Run the following command to generate the apache logs
$ ./gen.sh
Usage: ./gen.sh gb_per_day days workers dest

# gb_per_day - Log file size in GB per day 
# days - Number of days the log file should be generated
# workers - Parallel threads/workers to speed up the data generation per day
# dest - Destination directory where the log files will be created

# For example, if you want to generate 1TB/day for 7 days the following command
# will create 7 x 1TB files, one for every day from (current_day - 7) days
# using 12 workers and create them under /logs/apache directory

$ ./gen.sh 1024 7 12 /logs/apache/

# The log files generated will look the following
apache_05_08_2020.log  apache_05_09_2020.log  apache_05_10_2020.log  apache_05_11_2020.log  apache_05_12_2020.log
apache_05_13_2020.log  apache_05_14_2020.log
