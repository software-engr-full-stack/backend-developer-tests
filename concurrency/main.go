package main

import (
    "time"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/concurrency/pkg/concurrency"
)

func main() {
    delays := []int{3, 1, 2, 5, 6, 10, 4}
    funList := []func(){}
    var totalDelay int
    for _, d := range delays {
        funList = append(funList, funFactory(d))
        totalDelay += d
    }

    if false {
        runWithoutConcurrency(funList, totalDelay)
    }

    const maxConcurrent = 4
    simplePool := concurrency.NewSimplePool(maxConcurrent)
    for _, fun := range funList {
        simplePool.Submit(fun)
    }
    simplePool.Wait()
}

func funFactory(d int) func() {
    return func() {
        fmt.Printf("Function with %d-second delay starting\n", d)
        time.Sleep(time.Duration(d) * time.Second)
        fmt.Printf("Function with %d-second delay terminating normally\n", d)
    }
}

func runWithoutConcurrency(funList []func(), totalDelay int) {
    fmt.Println("No concurrency...")
    for _, fun := range funList {
        fun()
    }
    fmt.Printf("Total delay: %d seconds\n", totalDelay)
}
