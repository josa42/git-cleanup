#!/bin/bash

sed -E -e "s/\"git-cleanup [^\"]+\"/\"git-cleanup $1\"/" -i.bak main.go || exit 1
rm main.go.bak || exit 1
