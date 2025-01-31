# icalgen

It is a simple script to generate ical files from a text file.

## Build

```
go build
```

## Usage

#### Prepare an input text file in format
```
title: NAME OF EVENT
MONTH1 DAY1 DAY2 DAY3 ...

title: Another event
01 10
02 11
03 13 15
...

```

#### Call the script
```
cat example.txt | ./icalgen > output.ical
```
to get a `.ical` file which can be imported to calendar apps.

Optionally you can provide year for events using the switch '-year'. Alarms for events are hardcoded.


