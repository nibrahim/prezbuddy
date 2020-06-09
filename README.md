# Prezbuddy

Simple program that displays the time required by each section of a presentation using `libxosd`. Requires an input file which gives duration of each section as well as the name of the section. An example has been provided in this repo as `sample_input.txt`

## Installation

To just get the program and run it `go get github.com/nibrahim/prezbuddy` and then `./bin/prezbuddy src/github.com/nibrahim/prezbuddy/sample_input.txt`

To get the source and everything else, clone this repository and run `go build .` from the repo root to get the `prezbuddy` executable.

## Running

    ./prezbuddy <sample_input.txt>
    
## Format of input file

     # Lines that start with # are ignored
     # You can use this to add comments and 
     # metadata
     # Each line is of the form 
     #     mm:ss Section name
     # The space between the duration (mm:ss) and the Section name is
     # complusory
     
     0:5 Introduction
     1:00 The problem
     4:30 Our solution
     5:15 Implementation
     5:00 Questions
     1:30 Thanks


## Limitations

Relies on xosd to get the job done and hence needs `libxosd` installed. Works only on Gnu/Linux.
