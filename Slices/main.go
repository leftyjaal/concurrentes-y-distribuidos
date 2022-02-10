package main

import "fmt"

func main() {
    var (
        lenght uint
        cache  uint
        total  uint
        i      uint
        j      uint
    )

    s := make([]uint, 0, 1000)

    fmt.Println("Ingrese el tamaño del Slice: ")
    fmt.Scan(&lenght)

    for i = 0; i < lenght; i++ {
        fmt.Scan(&cache)
        s = append(s, uint(cache))
    }

    for j = 0; j < lenght; j++ {
        total = total + s[j]
    }
    fmt.Println(total)
}