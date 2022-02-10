package main

import (
  "encoding/gob"
  "fmt"
  "net"
  "strconv"
  "time"
)

var s = make([]*Proceso, 0)

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

     p.ik++

     if !p.aumentar {
        break
     }
     fmt.Println(strconv.Itoa(p.id) + " : " + strconv.FormatInt(p.ik, 10))
     time.Sleep(time.Millisecond * 500)
  }

}

func server() {
  s, err := net.Listen("tcp", ":9999")
  if err != nil {
     fmt.Println(err)
     return
  }
  for {
     c, err := s.Accept()
     if err != nil {

        fmt.Println(err)
        continue
     }
     go handleClient(c)

  }
}

func handleClient(c net.Conn) {
  var m Mensaje
  err := gob.NewDecoder(c).Decode(&m)
  if err != nil {
     fmt.Println(err)

  } else {
     if m.Msg == "nuevo" {
        if len(s) > 0 {

           id, ultimo := getLast()
           a := make([]*Proceso, 0)
           for _, i := range s {
              if i.id != id {
                 a = append(a, i)

              } else {

                 i.aumentar = false
                 aux := Mensaje{
                    Msg: "Respuesta",
                    Id:  id,
                    Ik:  s[ultimo].ik,
                 }
                 err := gob.NewEncoder(c).Encode(aux)
                 if err != nil {
                    fmt.Println(err)
                 }
                 c.Close()
              }

           }
           s = a
        } else {
           fmt.Println("Error no hay mas procesos Disponibles")
           aux := Mensaje{
              Msg: "ERROR",
              Id:  0,
              Ik:  0,
           }
           err := gob.NewEncoder(c).Encode(aux)
           if err != nil {
              fmt.Println(err)
           }

        }
     } else {

        aux := Proceso{ik: m.Ik, id: m.Id, aumentar: true}
        s = append(s, &aux)
        go aux.run()

     }

  }

}

func getLast() (int, int) {
  tam := len(s)
  id := s[tam-1].id
  return id, tam - 1
}

func main() {
  for i := 0; i < 5; i++ {
     aux := Proceso{
        ik:       0,
        id:       i,
        aumentar: true}
     s = append(s, &aux)
     go aux.run()
  }


  go server()
  var input string
  fmt.Scanln(&input)

}
