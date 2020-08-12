#!/bin/bash

WDIR=$PWD

for f in *; do
    if [ -d "$f" ]; then
        cd $WDIR/"$f"
        go test -v -race ./...
        cd $WDIR
    fi
done