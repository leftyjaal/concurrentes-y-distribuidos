package main

import (
  "fmt"
  "os"
  "sort"
)

func main(){
  file, err := os.Create("asecendente.txt")
  file2, err2 := os.Create("descendente.txt")
  if err != nil{
     fmt.Println(err)
     return
  }
  if err2 != nil {
     fmt.Println(err2)
     return
  }
  defer file.Close()
  defer file2.Close()

  s := make([]string, 0, 1000)
  var lenght int
  var cache string

  fmt.Println("Ingrese el tamaño del Slice: ")
  fmt.Scan(&lenght)
  fmt.Println(lenght)

  for i := 0 ; i < lenght ; i++ {
     fmt.Println("Ingrese la cadena: ")
     fmt.Scan(&cache)
     s = append(s, cache)
  }

  fmt.Println("Escribiendo...")
  //sort
  sort.Strings(s)
  for j := 0; j < lenght; j++{
     file.WriteString(s[j])
     file.WriteString("\n")
  }
  //rerversa
  sort.Sort(sort.Reverse(sort.StringSlice(s)))
  for k := 0; k < lenght; k++{
     file2.WriteString(s[k])
     file2.WriteString("\n")
  }
}