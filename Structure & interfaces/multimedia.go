package multimedia

import "fmt"

type Multimedia interface {
  Mostrar()
}

type ContenidoWeb struct {
  SMultimedia [] Multimedia
}

type Imagen struct {
  Titulo string
  Formato string
  Canales string
}

func (i *Imagen) Mostrar() {
  var str string
  str += i.Titulo
  str += "\n"
  str += i.Formato
  str += "\n"
  str += i.Canales
  fmt.Println(str)
}

type Audio struct {
  Titulo string
  Formato string
  Duracion string
}

func (a *Audio) Mostrar()  {
  var str string
  str += a.Titulo
  str += "\n"
  str += a.Formato
  str += "\n"
  str += a.Duracion
  fmt.Println(str)
}

type Video struct {
  Titulo string
  Formato string
  Frames string
}

func (v *Video) Mostrar() {
  var Str string
  Str += v.Titulo
  Str += "\n"
  Str += v.Formato
  Str += "\n"
  Str += v.Frames
  fmt.Println(Str)
}
