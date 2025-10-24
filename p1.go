package main

import (
  "fmt"
)

const MAX_PLAZAS = 2 // plazas por mecanico
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const GREEN = "\033[1;32m"
const BLUE = "\033[1;34m"
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
    switch(opt){
      case 1:
        t.ListarIncidencias()
      case 2:
        var m Mecanico
        t.ListarIncidenciasMecanico(m)
      case 3:
        t.ListarVehiculos()
    }
  }
}

func (t Taller)MenuClientes(){
  if (len(t.Clientes) > 0){
    return
  } else {
    warningMsg("No hay clientes en el taller")
  }
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

func (t Taller)ObtenerIncidencias() ([]Incidencia){
  var incidencias []Incidencia
  
  infoMsg("Obteniendo incidencias")
  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      for _, i := range v.Incidencias{
        incidencias = append(incidencias, i)
      }
    }
  }

  return incidencias
}

func (t Taller)ObtenerVehiculos() ([]Vehiculo){
  var vehiculos []Vehiculo

  infoMsg("Obteniendo vehículos")
  for _, c := range t.Clientes{
    for _, v := range c.Vehiculos{
      vehiculos = append(vehiculos, v)
    }
  }

  return vehiculos
}

func (t Taller)ListarIncidencias(){
  incidencias := t.ObtenerIncidencias()
  
  if (len(incidencias) > 0){
    for _, i := range incidencias{
      i.Visualizar()
    }
  } else {
    warningMsg("No hay incidencias en el taller")
  }
}

func (t Taller)ListarIncidenciasMecanico(m Mecanico){
  incidencias := t.ObtenerIncidencias()

  if (len(incidencias) > 0){
    for _, i := range incidencias{
      for _, m_aux := range i.Mecanicos{
        if (m.Id == m_aux.Id){ // Añadir función comparar Mecánicos
          i.Visualizar()
        }
      }
    }
  } else {
    warningMsg("No hay incidencias en el taller")
  }
}

func (t Taller)ListarVehiculos(){
  vehiculos := t.ObtenerVehiculos()

  if (len(vehiculos) > 0){
    fmt.Printf("\t%s·%s", BOLD, END)
    for _, v := range vehiculos{
      v.VisualizarMinimo()
    }
  } else {
    warningMsg("No hay vehículos en el taller")
  }
}


type Cliente struct{
  Id int
  Nombre string
  Telefono int
  Email string
  Vehiculos []Vehiculo
}

func (c Cliente)Visualizar(){
  fmt.Printf("%sID:%s %s\n", BOLD, END, c.Id)
  fmt.Printf("%sNombre:%s %s\n", BOLD, END, c.Nombre)
  fmt.Printf("%sTeléfono:%s %s\n", BOLD, END, c.Telefono)
  fmt.Printf("%sEmail:%s %s\n", BOLD, END, c.Email)
  if (len(c.Vehiculos) > 0){
    for _, v := range c.Vehiculos{
      fmt.Printf("\t%s·%s", BOLD, END)
      v.VisualizarMinimo()
    }
  } else {
    fmt.Println(BOLD, "SIN VEHÍCULOS", END)
  }
}

func (c Cliente)Modificar(){
  
}

func (c Cliente)ListarVehículos(){  

  if (len(c.Vehiculos) > 0){
    for _, v := range c.Vehiculos{
      fmt.Println("-------------------")
      v.Visualizar()
      fmt.Println("-------------------")
    }
  } else {
    warningMsg("El cliente no tiene vehículos")
  }
}


type Vehiculo struct{
  Matricula string // 1324ACB
  Marca string
  Modelo string
  FechaEntrada string
  FechaSalida string
  Incidencias []Incidencia
}

func (v Vehiculo)VisualizarMinimo(){
  fmt.Println(v.Marca, v.Modelo, " (", v.Matricula, ")")
}

func (v Vehiculo)Visualizar(){
  fmt.Printf("%sMatrícula:%s %s\n", BOLD, END, v.Matricula)
  fmt.Printf("%sMarca:%s %s\n", BOLD, END, v.Marca)
  fmt.Printf("%sModelo:%s %s\n", BOLD, END, v.Modelo)
  fmt.Printf("%sFecha de entrada:%s %s\n", BOLD, END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de salida:%s %s\n", BOLD, END, v.FechaSalida)
  // Visualizar incidencias (mínimo)
}

func (v Vehiculo)Modificar(){

}

func (v Vehiculo)Asignar(){

}

func (v Vehiculo)ListarIncidencias(){
  if (len(v.Incidencias) > 0){
    for _, i := range v.Incidencias{
      fmt.Println("-------------------")
      i.Visualizar()
      fmt.Println("-------------------")
    }
  } else {
    warningMsg("El cliente no tiene vehículos")
  }
}


type Incidencia struct{
  Id int
  Mecanicos []Mecanico
  Tipo int // Mecánica, eléctrica, carroceria (0, 1, 2)
  Prioridad int // 1 a 3 (alta a baja)
  Descripción string
  Estado int // (0) Cerrada, (1) Abierta, (2) En proceso
}

func (i Incidencia)Visualizar(){
  fmt.Printf("%sId:%s %s\n", BOLD, END, i.Id)
  //fmt.Printf("%sMecánicos:%s %s\n", BOLD, END, i.Mecanicos)
  fmt.Printf("%sTipo de incidencia:%s %s\n", BOLD, END, i.Tipo)
  fmt.Printf("%sPrioridad:%s %s\n", BOLD, END, i.Prioridad)
  fmt.Printf("%sDescripción:%s %s\n", BOLD, END, i.Descripción)
  fmt.Printf("%sEstado:%s %s\n", BOLD, END, i.Estado)
}


type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecánica, eléctrica, carroceria (0, 1, 2)
  Experiencia int
  Alta bool
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%sId:%s %s\n", BOLD, END, m.Id)
  fmt.Printf("%sNombre:%s %s\n", BOLD, END, m.Nombre)
  fmt.Printf("%sEspecialidad:%s %s\n", BOLD, END, m.Especialidad)
  fmt.Printf("%sExperiencia:%s %s\n", BOLD, END, m.Experiencia)
  fmt.Printf("%s¿En alta?:%s %s\n", BOLD, END, m.Alta)
}

func warningMsg(msg string){
  fmt.Printf("%s%s%s\n", YELLOW, msg, END)
}

func infoMsg(msg string){
  fmt.Printf("%s%s%s\n", BLUE, msg, END)
}

func menuFunc(menu []string) (int, int){
  var opt int

  menu = append(menu, "Salir")
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
    "Mecánicos"}

  var t Taller
  var exit bool = false

  for{
    opt, status := menuFunc(mainMenu)
    if status == 0 { // 0: Ok, 1: Error
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
