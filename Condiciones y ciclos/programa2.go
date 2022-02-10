package main

import (
    "fmt"
)

func main() {
    var e float64 = 0
    var sum float64 = 1
    var limit float64
    var i float64
    i = 1

    fmt.Println("# de ciclos:")
    fmt.Scan(&limit)

    for i < limit {
        e += sum
        sum /= i
        i++
    }

    fmt.Println(e)
}
