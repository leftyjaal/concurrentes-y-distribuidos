package main

import "fmt"

func Burbuja(unSlice []int){
    sliceLenght := len(unSlice)

    for i := sliceLenght ; i > 0 ; i--{
        for j := 1 ; j < i ; j++ {
            if unSlice[j-1] > unSlice[j]{
                slot := unSlice[j]
                unSlice[j] = unSlice[j-1]
                unSlice[j-1] = slot

            }//end if
        }// end for j
    }//end for i

    fmt.Println(unSlice)
}

func Fibonacci(n uint) uint {
    if n <= 1 {
        return n
    } else {
        return Fibonacci(n - 1) + Fibonacci(n - 2)
    }
}

func Greater(args... int) int{
    higher := 0
    for _, i := range args {
        if i > higher {
            higher = i
        }
    }
    return higher
}

func generadorInPares() func() uint {
    i := uint(0)
    return func() uint {
        var non = i
        if i % 2 == 0 {
            i += 1
        } else {
            i += 2
        }
        return non
    }
}

func intercambia(a, b *int){
    temp := *a
    *a = *b
    *b = temp
    fmt.Println(*a, *b)
}

func main(){
    elSlice := []int{100, 3, 8, 45, 1, 3,56, 24, 87, 77, 34, 42, 41, 40, 4, 6, 12, 9, 20, 1, 2, 101, 102, 300, 4}
    nextNon := generadorInPares()
    a, b := 10, 15

    defer Burbuja(elSlice)

    fmt.Println("Fibonacci(10)")
    fmt.Println(Fibonacci(10))

    fmt.Println("Greater()")
    fmt.Println(Greater(elSlice...))

    fmt.Println("generadorInPares()")
    fmt.Println(nextNon())
    fmt.Println(nextNon())
    fmt.Println(nextNon())
    fmt.Println(nextNon())
    fmt.Println(nextNon())

    fmt.Println("intercambia(10, 15)")
    intercambia(&a, &b)

    fmt.Println("Burbuja()")
}
