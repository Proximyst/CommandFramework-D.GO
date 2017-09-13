#!/bin/bash

function commit {
    ./commit.sh work.
}
while [ : ]
do
    commit
    sleep 5
done