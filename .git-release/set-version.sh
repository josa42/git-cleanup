#!/bin/bash

sed -E -e "s/\"Git Cleanup [^\"]+\"/\"Git Cleanup $1\"/" -i.bak main.go || exit 1
rm main.go.bak || exit 1
