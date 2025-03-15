#!/bin/sh
clear
echo "$1"

if [ "$1" = "build" ]
then
    echo -e "Building app..."
else
    echo -e Running app...
    go run .
fi