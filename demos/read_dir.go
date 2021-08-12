package main

import (
    "log"
    "fmt"
    "io/fs"
    "os"
)


func ExampleReadDir(path string) {
    files, err := os.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
        fmt.Printf("\t%t\n", file)
        // fmt.Println(file.Mode())
        fmt.Println()
    }
}

func ExampleFileMode(path string) {
    fi, err := os.Lstat(path)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0777, etc.
    switch mode := fi.Mode(); {
    case mode.IsRegular():
        fmt.Println("regular file")
        fmt.Println(mode.IsRegular())
    case mode.IsDir():
        fmt.Println("directory")
        fmt.Println(mode.IsDir())
    case mode&fs.ModeSymlink != 0:
        fmt.Println("symbolic link")
    case mode&fs.ModeNamedPipe != 0:
        fmt.Println("named pipe")
    }
}


func main() {
    ExampleReadDir(".")
    ExampleFileMode("xml_parse.go")
    ExampleFileMode("wailsdemo")
}