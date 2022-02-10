package main

import "fmt"

func main() {
    var (
        day   int
        month int
    )

    fmt.Println("Ingrese su dia y mes de nacimiento:")
    fmt.Scan(&day)
    fmt.Scan(&month)

    if month == 2 && day > 29 {
        fmt.Println("Febrero no tiene esos dias")
    } else if day <= 31 && day >= 1 && month >= 1 && month <= 12 {
        switch month {
        case 1:
            if day < 20 {
                fmt.Println("Capricornio")
            } else {
                fmt.Println("Acuario")
            }
        case 2:
            if day < 19 {
                fmt.Println("Acuario")
            } else {
                fmt.Println("Piscis")
            }
        case 3:
            if day < 21 {
                fmt.Println("Piscis")
            } else {
                fmt.Println("Aries")
            }
        case 4:
            if day < 20 {
                fmt.Println("Aries")
            } else {
                fmt.Println("Tauro")
            }
        case 5:
            if day < 21 {
                fmt.Println("Tauro")
            } else {
                fmt.Println("Geminis")
            }
        case 6:
            if day < 21 {
                fmt.Println("Geminis")
            } else {
                fmt.Println("Cancer")
            }
        case 7:
            if day < 23 {
                fmt.Println("Cancer")
            } else {
                fmt.Println("Leo")
            }
        case 8:
            if day < 23 {
                fmt.Println("Leo")
            } else {
                fmt.Println("Virgo")
            }
        case 9:
            if day < 23 {
                fmt.Println("Virgo")
            } else {
                fmt.Println("Libra")
            }
        case 10:
            if day < 23 {
                fmt.Println("Libra")
            } else {
                fmt.Println("Escorpio")
            }
        case 11:
            if day < 22 {
                fmt.Println("Escorpio")
            } else {
                fmt.Println("Sagitario")
            }
        case 12:
            if day < 22 {
                fmt.Println("Sagitario")
            } else {
                fmt.Println("Capricornio")
            }
        default:
            fmt.Println("")
        }
    } else {
        fmt.Println("Verifique sus datos")
    }
}