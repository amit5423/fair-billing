# `fair-billing`

## Problem Statement
You work for a hosted application provider which charges for the use of its application by the duration of sessions. There is a charge per second of usage. The usage data comes from a log file that lists the time at which a session starts or stops (in the format HH:MM:SS), the name of the user (which is a 
alphanumeric string, of arbitrary length)

Unfortunately, the developer of the application omitted some vital information from the log file. There is no indicator of which start and end lines are paired together. Even more, unfortunately, the log files are re-written on a regular basis, so sessions may overlap the time boundaries of the log file. In other words, there may be “End” entries for sessions that were already in progress when the log file started, which will have no preceding `Start`. Similarly, when the log files are retrieved, there may be sessions still in progress that have a `Start` but no `End`. 

A user can also have more than one session active concurrently:
* Where there is an `End` with no possible matching start, the start time should be assumed to be the earliest time of any record in the file.
* Where there is a `Start` with no possible matching “End”, the end time should be assumed to be the latest time of any record in the file. 

## Task
The task is to take the log file and print out the user's session report. 

### Input
The program takes a single command line parameter, which will be the path to the data file to read. We can assume that the data in the input will be correctly ordered chronologically and that all records in the file will be from within a single day (i.e. they will not span midnight). 

The usage data comes from a log file that lists the time at which a session starts or stops (in the format HH:MM:SS), the name of the user (which is a single alphanumeric string, of arbitrary length), and whether this is the start or end of the session and whether this is the start or end of the session.

### Expected Output
This is a command-line application that takes a log file as input and prints report of the users, the number of sessions, and the minimum possible total duration of their sessions in seconds.

## Usage  

#### Clone the repository
```
git clone https://github.com/amit5423/fair-billing.git
```
Run application, providing log/data files as an argument 
```
go run main.go logfiles/logs0
go run main.go logfiles/logs1
```
#### Unit testing

Run following command where unit test files are present i.e., `fair-billing/main_test.go` 

``` go test ```
### More Details

* All Golang's standard libraries are used in the solution.
* Any invalid or irrelevant data (any lines that do not contain a valid time-stamp, username, and a `Start` or `End` marker) are silently ignored and not included in any calculations while processing the log files