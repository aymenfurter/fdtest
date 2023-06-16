package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "os"
)

func main() {
        var i uint64
        var files []*os.File

        for {
                fileName := fmt.Sprintf("/tmp/dummyFile%d.txt", i)
                file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
                if err != nil {
                        log.Fatalf("Failed opening file: %s", err)
                }

                files = append(files, file)

                if i%10000 == 0 {
                        pid := os.Getpid()
                        fdDir := fmt.Sprintf("/proc/%d/fd", pid)
                        fds, err := ioutil.ReadDir(fdDir)
                        if err != nil {
                                log.Fatalf("Failed reading dir: %s", err)
                        }

                        log.Printf("Currently %d file descriptors are being used", len(fds))
                }
                i++
        }
}
