package main

import (
    "fmt"
    "os"
    "runtime"
    "sync"
)

func main() {
    src, dst, quiet := parseArgs()

    sFile, err := os.Open(src)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: opening source file: %v\n", err)
        return
    }
    defer sFile.Close()

    info, _ := sFile.Stat()
    fileSize := info.Size()

    dFile, err := os.Create(dst)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: creating destination file: %v\n", err)
        return
    }
    defer dFile.Close()

    if err := dFile.Truncate(fileSize); err != nil {
        fmt.Fprintf(os.Stderr, "Error: allocating space in destination: %v\n", err)
        return
    }

    numWorkers := runtime.NumCPU()
    chunkChan := make(chan int64)

    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            buf := make([]byte, chunkSize)
            for pos := range chunkChan {
                readChunk(sFile, dFile, buf, pos, fileSize, quiet)
            }
        }()
    }

    for pos := int64(0); pos < fileSize; pos += chunkSize {
        chunkChan <- pos
    }
    close(chunkChan)

    wg.Wait()
    dFile.Sync()

    if !quiet {
        fmt.Println("\nFile copied successfully.")
    }
}
