package main

import (
  "fmt"
)

const PLAZAS_MECANICO = 2
const BOLD = "\033[1;37m"
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const END = "\033[0m"

type Taller struct{
  Mecanicos []Mecanico
  UltimoId int
}

func (t *Taller)CrearMecanico(nombre string, especialidad int, experiencia int){
  var err bool = false
  var m Mecanico

  if len(nombre) > 0{
    m.Nombre = nombre
  } else {
    err = true 
    errorMsg("Nombre inválido")
  }

  if especialidad >= 0 && especialidad <= 2{
    m.Especialidad = especialidad
  } else {
    err = true 
    errorMsg("Especialidad no reconocida")
  }

  if experiencia >= 0{
    m.Experiencia = experiencia
  } else {
    err = true 
    errorMsg("Valor de experiencia inválido")
  }

  if !err{
    t.UltimoId++
    m.Id = t.UltimoId
    m.Alta = true
    t.Mecanicos = append(t.Mecanicos, m)
  }
}

func (t Taller)ObtenerMecanicoPorId(id int) (Mecanico){
  var res Mecanico

  for i, m := range t.Mecanicos{
    if m.Id == id{
      res = t.Mecanicos[i]
    }
  }

  return res
}


type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecanica, Electrica, Carroceria
  Experiencia int
  Alta bool
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%s ID: %s%03d\n", BOLD, END, m.Id)
  fmt.Printf("%s Nombre: %s%s\n", BOLD, END, m.Nombre)
  fmt.Printf("%s Especialidad: %s%s\n", BOLD, END, m.ObtenerEspecialidad())
  fmt.Printf("%s Experiencia: %s%d años\n", BOLD, END, m.Experiencia)
  fmt.Printf("%s ¿Está de alta? %s%s\n", BOLD, END, m.Alta)
}

func (m *Mecanico)Modificar(){
  titulo := fmt.Sprintf("Modificar datos de %s", m.Nombre)
  menu := []string{
    titulo,
    "Nombre",
    "Especialidad",
    "Experiencia"}

  for{
    opt, status := menuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          fmt.Println(m.Nombre)
        case 2:
          fmt.Println(m.ObtenerEspecialidad())
        case 3:
          fmt.Println(m.Experiencia)
      }
    } else if status == 2{
      break
    }
  }
}

func (m Mecanico)ObtenerEspecialidad() (string){
  switch m.Especialidad{
    case 0:
      return "Mecánica"
    case 1:
      return "Electrónica"
    case 2:
      return "Carrocería"
    default:
      return "Sin especialidad"
  }
}


func errorMsg(msg string){
  fmt.Printf("%s%s%s", RED, msg, END)
}

func warningMsg(msg string){
  fmt.Printf("%s%s%s", YELLOW, msg, END)
}

func leerInt(i *int){
  for{
    fmt.Print("> ")
    fmt.Scanf("%d", i)
    if (*i >= 0){
      break
    } else {
      warningMsg("Valor entero inválido")
    }
  }
}

func leerStr(str *string){
  for{
    fmt.Print("> ")
    fmt.Scanf("%s", str)
    if (len(*str) > 0){
      break
    } else {
      warningMsg("Cadena de texto inválida")
    }
  }
}

func menuFunc(menu []string) (int, int){
  var opt int

  menu = append(menu, "Salir")
  fmt.Printf("%s%s%s\n", BOLD, menu[0], END) // Menu title

  for i:= 1; i < len(menu); i++{
    fmt.Printf("%d.- %s\n", i, menu[i])
  }

  leerInt(&opt)

  if opt > 0 && opt < len(menu) - 1{
    return opt, 0
  } else if opt == len(menu) - 1{
    return opt, 2
  }
  return 0, 1
}


func main(){
  var t Taller
  var m Mecanico

  t.CrearMecanico("Pepe", 0, 0)
  t.CrearMecanico("Pepe", 0, 0)
  m = t.ObtenerMecanicoPorId(1)
  m.Visualizar()
  m.Modificar()
}
