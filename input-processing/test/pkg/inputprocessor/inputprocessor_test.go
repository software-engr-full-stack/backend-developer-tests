package main

import (
    "testing"
    "bufio"
    "os"
    "bytes"
    "io"
    "strings"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/input-processing/pkg/inputprocessor"
)

func TestRun(t *testing.T) {
    type testType struct {
        inputFile string
        expectedFile string
    }

    tests := []testType{
        testType{inputFile: "./fixtures/input-one.txt", expectedFile: "./fixtures/expected-one.txt"},

        // // Hangs after 10 minutes
        // testType{inputFile: "./fixtures/input-two.txt", expectedFile: "./fixtures/expected-two.txt"},
    }

    for _, test := range tests {
        file, err := os.Open(test.inputFile)
        title := "Input file read error"
        if actual, expected := err, error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        defer file.Close() //nolint:gocritic,deferInLoop

        reader := bufio.NewReader(file)

        originalStdout := os.Stdout
        r, w, err := os.Pipe()
        title = "Pipe creation error"
        if actual, expected := err, error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        os.Stdout = w

        title = "error"
        if actual, expected := inputprocessor.Run(reader), error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        outChan := make(chan string)
        go func() {
            var buf bytes.Buffer
            _, err = io.Copy(&buf, r)
            if err != nil {
                panic(err)
            }
            outChan <- buf.String()
        }()

        w.Close()
        os.Stdout = originalStdout
        out := <-outChan

        expFile, err := os.Open(test.expectedFile)
        title = "Expected file read error"
        if actual, expected := err, error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        defer expFile.Close() //nolint:gocritic,deferInLoop

        scanner := bufio.NewScanner(expFile)
        actualLines := strings.Split(out, "\n")
        ix := 0
        for scanner.Scan() {
            title = fmt.Sprintf("Line %d", ix + 1)
            if actual, expected := actualLines[ix], scanner.Text(); actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            ix++
        }

        title = "Line count"
        if actual, expected := len(actualLines), ix + 1; actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        title = "scanner error"
        if actual, expected := scanner.Err(), error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
    }
}
