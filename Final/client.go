package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)


func client() {
	c, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1.- Registrar calificacion")
		fmt.Println("2.- Mostrar promedio de un alumno")
		fmt.Println("3.- Mostrar promedio general")
		fmt.Println("4.- Mostrar promedio de una materia")
		fmt.Println("5.- Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name string
			var grade string
			var subject string
			var result string
			fmt.Print("Alumno: ")
			scanner.Scan()
			name = scanner.Text()
			fmt.Print("Calificación: ")
			scanner.Scan()
			grade = scanner.Text()
			fmt.Print("Materia: ")
			scanner.Scan()
			subject = scanner.Text()
			slice := []string{name, subject, grade}
			err = c.Call("Server.Registrar", slice, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case 2:
			var name string
			fmt.Print("Alumno: ")
			scanner.Scan()
			name = scanner.Text()

			var result float64
			err = c.Call("Server.PromedioAlumno", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de ", name, " es: ", result)
			}
			scanner.Scan()
		case 3:
			var result float64
			err = c.Call("Server.PromedioGeneral", "Promedio General", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es: ", result)
			}
			scanner.Scan()
		case 4:
			var name string
			fmt.Print("Nombre de la materia: ")
			scanner.Scan()
			name = scanner.Text()

			var result float64
			err = c.Call("Server.PromedioMateria", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de la materia ", name, "es: ", result)
			}
			scanner.Scan()
		case 5:
			return
		}
	}
}

func main() {
	client()
}