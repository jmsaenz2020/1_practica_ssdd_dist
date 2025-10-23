package main

import (
  "fmt"
)

const MAX_PLAZAS = 2 // plazas por mecanico
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const GREEN = "\033[1;32m"
const BOLD = "\033[1;37m"
const END = "\033[0m"

type Taller struct{
  Clientes []Cliente
  Mecanicos []Mecanico
}

func (t Taller)MenuPrincipal(){
  menu := []string{
    "Opciones del taller",
    "Lista de Incidencias",
    "Lista de Incidencias de Mecánico",
    "Lista de vehículos"}

  opt, status := menuFunc(menu)

  if (status == 0){
    fmt.Println(menu[opt])
  }
}

func (t Taller)MenuClientes(){
  if (len(t.Clientes) > 0){
    return
  } else {
    warningMsg("No hay clientes en el taller")
  }
}

func (t Taller)ObtenerVehiculos() ([]Vehiculo){
  var vehiculos []Vehiculo

  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      vehiculos = append(vehiculos, v)
    }
  }

  return vehiculos
}

func (t Taller)MenuVehiculos(){
  vehiculos := t.ObtenerVehiculos()  

  if (len(vehiculos) > 0){
    return
  } else {
    warningMsg("No hay vehículos en el taller")
  }
}

func (t Taller)MenuIncidencias(){
  vehiculos := t.ObtenerVehiculos()
  var incidencias []Incidencia

  
  if (len(vehiculos) > 0){
    for _, v := range vehiculos{
      for _, inc := range v.Incidencias{
        incidencias = append(incidencias, inc)
      }
    }
    if (len(incidencias) > 0){
      return
    } else {
      warningMsg("No hay incidencias en el taller")
    }
  } else {
    warningMsg("No hay vehiculos en el taller")
  }
}

func (t Taller)MenuMecanicos(){
  if(len(t.Mecanicos) > 0){
    return
  } else {
    warningMsg("No hay mecánicos en el taller")
  }
}


type Cliente struct{
  Id int
  Nombre string
  Telefono int
  Email string
  Vehiculos []Vehiculo
}

type Vehiculo struct{
  Matricula string // 1324ACB
  Marca string
  Modelo string
  FechaEntrada string
  FechaSalida string
  Incidencias []Incidencia
}

type Incidencia struct{
  Id int
  Mecanicos []Mecanico
  Tipo int // Mecánica, eléctrica, carroceria (0, 1, 2)
  Prioridad int // 1 a 3 (alta a baja)
  Descripción string
  Estado int // (0) Cerrada, (1) Abierta, (2) En proceso
}

type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecánica, eléctrica, carroceria (0, 1, 2)
  Experiencia int
  Alta bool
}


func warningMsg(msg string){
  fmt.Printf("%s%s%s\n", YELLOW, msg, END)
}

func menuFunc(menu []string) (int, int){
  var opt int

  fmt.Printf("%s%s%s\n", BOLD, menu[0], END) // Menu title

  for i:= 1; i < len(menu); i++{
    fmt.Printf("%d.- %s\n", i, menu[i])
  }

  fmt.Print("> ")
  fmt.Scanf("%d", &opt)

  if (opt > 0 && opt < len(menu)){
    return opt, 0
  }
  return 0, 1
}

func main(){
  mainMenu := []string{
    "Menu principal",
    "Taller",
    "Clientes",
    "Vehiculos",
    "Incidencias",
    "Mecánicos",
    "Salir"}

  var t Taller
  var exit bool = false

  for{
    opt, status := menuFunc(mainMenu)
    if status == 0 { // ok
      switch opt{
        case 1:
          t.MenuPrincipal()
        case 2:
          t.MenuClientes()
        case 3:
          t.MenuVehiculos()
        case 4:
          t.MenuIncidencias()
        case 5:
          t.MenuMecanicos()
        case 6:
          exit = true
        default:
          break
      }
    }
    if (exit){
      break
    }
  }
}
