package main

import (
  "./multimedia"
  "fmt"
)


func main(){
  dataMultimedia := multimedia.ContenidoWeb{}
  var opcion int
  var (
     nombre string
     formato string
     comodin string
  )

  fmt.Println("Bienvenido!")
  fmt.Println("1.- Agregar Imagen")
  fmt.Println("2.- Agregar Audio")
  fmt.Println("3.- Agregar Video")
  fmt.Println("4.- Mostrar")
  fmt.Scanln(&opcion)

  switch opcion {

  case 1:

     fmt.Println("...Imagen...")

     fmt.Println("Nombre: ")
     fmt.Scanln(&nombre)

     fmt.Println("Formato: ")
     fmt.Scanln(&formato)

     fmt.Println("# canales: ")
     fmt.Scanln(&comodin)

     img := multimedia.Imagen{Titulo: nombre, Formato: formato, Canales: comodin}
     dataMultimedia = append(img)

  case 2:

     fmt.Println("...Audio...")

     fmt.Println("Nombre: ")
     fmt.Scanln(&nombre)

     fmt.Println("Formato: ")
     fmt.Scanln(&formato)

     fmt.Println("duracion: ")
     fmt.Scanln(&comodin)

     aud := multimedia.Audio{Titulo: nombre, Formato: formato, Duracion: comodin}
     dataMultimedia = append(aud)

  case 3:

     fmt.Println("...Video...")

     fmt.Println("Nombre: ")
     fmt.Scanln(&nombre)
     fmt.Println("Formato: ")
     fmt.Scanln(&formato)
     fmt.Println("# de frames: ")
     fmt.Scanln(&comodin)

     vid := multimedia.Video{Titulo: nombre, Formato: formato, Frames: comodin}
     dataMultimedia = append(vid)

  case 4:

     for _, contenido := range dataMultimedia.SMultimedia{
        contenido.Mostrar()
     }

  }


}