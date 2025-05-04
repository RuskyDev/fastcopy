package main

import (
    "fmt"
    "io"
    "os"
    "sync"
    "time"
)

const chunkSize = 8 * 1024 * 1024

var (
    totalCopied   int64
    progressMutex sync.Mutex
    startGlobal   = time.Now()
)

func readChunk(src *os.File, dst *os.File, buf []byte, pos int64, totalSize int64, quiet bool) {
    end := pos + chunkSize
    if end > totalSize {
        end = totalSize
    }
    n, err := src.ReadAt(buf, pos)
    if err != nil && err != io.EOF {
        return
    }
    _, err = dst.WriteAt(buf[:n], pos)
    if err != nil {
        return
    }
    if !quiet {
        updateProgress(int64(n), totalSize, end == totalSize)
    }
}

func updateProgress(n int64, total int64, done bool) {
    progressMutex.Lock()
    defer progressMutex.Unlock()

    totalCopied += n
    currMB := float64(totalCopied) / 1e6
    totMB := float64(total) / 1e6
    percent := float64(totalCopied) / float64(total)

    elapsed := time.Since(startGlobal).Seconds()
    speed := currMB / elapsed
    estimated := time.Duration(float64(startGlobal.Sub(time.Now()))/percent) * -1

    fmt.Printf("\rCopying file to %s [%.1f/%.1f MB] [Speed: %.2f MB/s] [ETA: %s]    ",
        os.Args[2],
        currMB,
        totMB,
        speed,
        estimated.Round(time.Second),
    )

    if done {
        fmt.Println()
    }
}
