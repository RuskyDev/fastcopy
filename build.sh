#!/bin/bash

platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

rm -rf bin/*

for platform in "${platforms[@]}"; do
  GOOS=$(echo $platform | cut -d'/' -f1)
  GOARCH=$(echo $platform | cut -d'/' -f2)

  if [ "$GOOS" == "windows" ]; then
    output="bin/windows/FastCopy.exe"
  elif [ "$GOOS" == "linux" ]; then
    output="bin/linux/FastCopy"
  elif [ "$GOOS" == "darwin" ]; then
    output="bin/macos/FastCopy"
  fi

  mkdir -p "$(dirname "$output")"

  GOOS=$GOOS GOARCH=$GOARCH go build -o "$output" src/main.go src/args.go src/chunk.go
done
