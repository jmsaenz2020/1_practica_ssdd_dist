package main

import (

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


func main(){

}
