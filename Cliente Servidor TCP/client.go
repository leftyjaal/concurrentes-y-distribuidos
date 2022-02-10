package main

import (
  "encoding/gob"
  "fmt"
  "net"
  "strconv"
  "time"
)

var mostrar = true
var servidorOcupado = false
var proceso = Proceso{}

type Proceso struct {
  ik       int64
  id       int
  aumentar bool
}
type Mensaje struct {
  Msg string
  Id  int
  Ik  int64
}

func (p *Proceso) run() {

  for {
     time.Sleep(time.Millisecond * 500)
     p.ik++
     if !p.aumentar {
        break
     }
     if mostrar {
        fmt.Println(strconv.Itoa(p.id) + " : " + strconv.FormatInt(p.ik, 10))
     }
  }
}

func cliente(m Mensaje) {
  c, err := net.Dial("tcp", ":9999")
  if err != nil {
     fmt.Println(err)
     return
  }
  err = gob.NewEncoder(c).Encode(m)
  if err != nil {
     fmt.Println(err)
  }
  go handleCliente(c)

}

func handleCliente(c net.Conn) {
  var mensaje Mensaje
  err := gob.NewDecoder(c).Decode(&mensaje)
  if err != nil {
     fmt.Println(err)
  } else {
     aux := Proceso{
        ik:       mensaje.Ik,
        id:       mensaje.Id,
        aumentar: true,
     }
     proceso = aux
     if mensaje.Msg == "ERROR" {
        println("no process avaible")
        servidorOcupado = true
     } else {
        go proceso.run()

     }
  }
  c.Close()

}

func main() {
  m := Mensaje{
     Msg: "new",
     Id:  0,
     Ik:  0,
  }
  go cliente(m)
  var input string
  fmt.Scanln(&input)
  proceso.aumentar = false
  n := Mensaje{
     Msg: "out",
     Id:  proceso.id,
     Ik:  proceso.ik,
  }
  if !servidorOcupado {
     c, err := net.Dial("tcp", ":9999")
     if err != nil {
        fmt.Println(err)
        return
     }

     err = gob.NewEncoder(c).Encode(n)
     if err != nil {
        fmt.Println(err)
     }
  }

}
