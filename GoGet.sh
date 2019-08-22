#!/bin/sh
if [ $# -eq 1 ]; then
        GOPATH=$1
fi

cd peekaboo/cmd
go get -v 
