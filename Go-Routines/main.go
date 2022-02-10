package main

import (
  "fmt"
  "time"
)

var cerrarProceso uint64
var vanderaImprimirValores bool


func proceso(id uint64, imprimirValores chan bool, cerrarProcesoID chan uint64) {
  i := 0
  for {
     msgID := <-cerrarProcesoID
     msgImprimirValor := <-imprimirValores
     if msgImprimirValor {
        fmt.Printf("Proceso : %d  Valor:%d\n", id, i)
     }
     time.Sleep(time.Millisecond * 500)
     if id == msgID {
        break
     }
     i++
  }
}

func menu (){
  fmt.Println("1. Agregar Proceso")
  fmt.Println("2. Imprimir Proceso")
  fmt.Println("3. Eliminar Proceso")
  fmt.Println("4. Salir")
}

func imprimirProcesos(canalImprimirValor chan bool) {
  for {
     canalImprimirValor <- vanderaImprimirValores
  }
}

func terminarProceso(idCerrarProceso chan uint64) {
  for {
     idCerrarProceso <- cerrarProceso
  }
}


func main() {

  var cerrarProcesoID chan uint64 = make(chan uint64)
  var canalImprimirValor chan bool = make(chan bool)
  var id uint64
  var opcion int8

  go imprimirProcesos(canalImprimirValor)
  go terminarProceso(cerrarProcesoID)

  for opcion != 4{
     menu()
     fmt.Scanln(&opcion)

     switch opcion {
     case 1:
        id++
        go proceso(id, canalImprimirValor, cerrarProcesoID)
        fmt.Printf("Proceso %d agregado \n", id)

     case 2:
        vanderaImprimirValores = true
        fmt.Scanln()
        vanderaImprimirValores = false

     case 3:
        fmt.Print("Digite el proceso que desea eliminar: ")
        fmt.Scanln(&cerrarProceso)
        fmt.Printf("Proceso %d Eliminado \n", cerrarProceso)

     case 4:
        fmt.Printf("...")

     default:
        fmt.Println("Opcion Incorrecta")
     }
  }

}
