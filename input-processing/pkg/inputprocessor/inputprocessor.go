package inputprocessor

import (
    "bufio"
    "io"
    "strings"
    "fmt"
)

func Run(reader *bufio.Reader) error {
    lineNo := 1
    for {
        text, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        fields := strings.Fields(text)

        for _, field := range fields {
            // We could use "strings.ToLower(field)" to make the search case-insensitive.
            // But the requirements didn't say so.
            // Something to keep in mind.
            if field == "error" {
                fmt.Printf("%d: %s", lineNo, text)
                continue
            }
        }
        lineNo++
    }
}
