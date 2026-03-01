#!/bin/bash
GO_CHANGING=false
for file in $(git diff --name-only $3..$2); do 
    if [[ $file == *.go && $GO_CHANGING == "false" ]]; then 
        GO_CHANGING=true
        break 
    fi
done

if [[ "$GO_CHANGING" == "true" ]]; 
    then echo "Go test are starting"; 
fi