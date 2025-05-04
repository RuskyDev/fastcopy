#!/bin/bash

[ -f "bin/FastCopy.exe" ] && rm "bin/FastCopy.exe"
go build -o "bin/FastCopy.exe" src/main.go src/args.go src/chunk.go