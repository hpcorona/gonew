#!/bin/sh

export GOOS=windows
export GOARCH=386

8g -o _go_.8 gonew.go
8l -o gonew.exe _go_.8
