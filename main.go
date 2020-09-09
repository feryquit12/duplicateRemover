package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

func removeDuplicate(lines []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range lines {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func main() {
    lines, err := readLines("tmp.txt")
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    lines = removeDuplicate(lines)
    //fmt.Println(lines)

    if err := writeLines(lines, "tmp-out.txt"); err != nil {
        log.Fatalf("writeLines: %s", err)
    }
}