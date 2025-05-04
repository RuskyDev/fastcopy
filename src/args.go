package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func parseArgs() (src, dst string, quiet bool) {
    if len(os.Args) < 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s <source> <destination> [--quiet]\n", os.Args[0])
        os.Exit(1)
    }
    src, dst = os.Args[1], os.Args[2]
    quiet = len(os.Args) > 3 && os.Args[3] == "--quiet"

    info, _ := os.Stat(dst)
    if info != nil && info.IsDir() {
        _, file := filepath.Split(src)
        dst = filepath.Join(dst, file)
    }
    return
}