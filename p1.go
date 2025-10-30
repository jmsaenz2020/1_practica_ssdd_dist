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
  Clientes []Cliente // Plazas
  Mecanicos []Mecanico
}

func (t Taller)MenuPrincipal(){
  menu := []string{
    "Opciones del taller",
    "Lista de Incidencias",
    "Lista de Incidencias de Mecánico",
    "Lista de vehículos"}
  var exit bool = false

  for{
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
        default:
          exit = true
      }
    }
    if exit{
      break
    }
  }
}

func (t Taller)MenuClientes() (Taller){
  var menu []string  
  var exit bool = false
  var cliente Cliente

  for ;;{
    menu = []string{"Menú de clientes", "Crear cliente"}
    for _, c := range t.Clientes{
      menu = append(menu, c.Nombre)
    }
    
    opt, status := menuFunc(menu)

    opt--
    if (status == 0 && opt >= 1 && opt <= len(t.Clientes)){
      opt--
      cliente = t.Clientes[opt]
      cliente = cliente.MenuCliente(t)
      t.Clientes[opt] = cliente
    } else if (status == 0 && opt == 0){ // Crear cliente
      cliente = crearCliente(t)
      t.Clientes = append(t.Clientes, cliente)
    } else if (status == 0 && opt == len(t.Clientes) + 1){
      exit = true
    }
    if (exit){
      break
    }
  }

  return t
}

func (t Taller)MenuVehiculos(){
  vehiculos := t.ObtenerVehiculos()  
  var exit bool = false

  if len(vehiculos) > 0{
    for {
      menu := []string{"Menú de vehículos"}
      for _, v := range vehiculos{
        menu = append(menu, v.Info())
      }
      opt, status := menuFunc(menu)
      if status == 0{
        if opt == len(vehiculos) + 1{
          exit = true
        } else {
          vehiculos[opt - 1] = vehiculos[opt - 1].MenuVehiculo()
        }
      }
      if exit{
        break
      }
    }
  } else {
    warningMsg("No hay vehículos en el taller")
  }
}

func (t Taller)MenuIncidencias(){
  incidencias := t.ObtenerIncidencias()

  if (len(incidencias) > 0){
    for {
      menu := []string{"Menú de Incidencias"}
      for _, i := range incidencias{
        menu = append(menu, string(i.Id))
      }

    }
  } else {
    warningMsg("No hay incidencias en el taller")
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

func (t Taller)ObtenerVehiculo() (Vehiculo, int){
  var v Vehiculo
  vehiculos := t.ObtenerVehiculos()
  menu := []string{"Selecciona un vehículo"}
  var exit bool = false  
  
  if len(vehiculos) > 0{
    for _, aux := range vehiculos{
      for _, c := range t.Clientes{
        if c.TieneVehiculo(aux){
          break
        }
      }
      fmt.Println("Nuevo coche")
      menu = append(menu, aux.Info())
    }

    for{
      opt, status := menuFunc(menu)
      if status == 0{
        if opt <= len(vehiculos){
          v = vehiculos[opt - 1]
        } else if opt == len(vehiculos) + 1{
          exit = true
        }
      }
      if exit{
        break
      }
    }
    return v, 0
  }
  return v, 1
  
}

func (t Taller)EstaLleno() (bool){
  vehiculos := t.ObtenerVehiculos()

  return !(len(vehiculos) < len(t.Mecanicos)*MAX_PLAZAS) 
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
    fmt.Printf("  %s·%s", BOLD, END)
    for _, v := range vehiculos{
      fmt.Println(v.Info())
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

func (c Cliente)MenuCliente(t Taller) (Cliente){
  menu := []string{
    "Opciones de " + c.Nombre,
    "Visualizar",
    "Modificar datos",
    "Listar vehículos",
  }
  var exit bool = false


  for{
    menu[0] = "Opciones de " + c.Nombre
    opt, status := menuFunc(menu)
    if (status == 0){
      switch opt{
        case 1:
          c.Visualizar()
        case 2:
          c = c.Modificar(t)
        case 3:
          c.ListarVehiculos()
        default:
          exit = true
      }
    }
    if (exit){
      break
    }
  }

  return c
}

func (c Cliente)MenuVehiculos(t Taller) (Cliente){
  menu := []string{
    "Opciones de modificación de vehículos",
    "Asignar",
    "Modificar",
    "Eliminar"}
  var exit bool = false
  //var v Vehiculo

  for{
    opt, status := menuFunc(menu)
    if (status == 0){
      switch opt{
        case 1:
          //v, status = t.ObtenerVehiculo()
          if (status == 0 && !t.EstaLleno()){
            //v.Asignar(&c)
            //c.Vehiculos = append(c.Vehiculos, v)
          } else {
            warningMsg("El taller está lleno")
          }
        case 2:
          if (len(c.Vehiculos) > 0){
            // Modificar Vehiculos
          } else {
            warningMsg("El cliente no tiene vehículos")
          }
        case 3:
          if (len(c.Vehiculos) > 0){
            // Eliminar Vehiculos
          } else {
            warningMsg("El cliente no tiene vehículos")
          }
        default:
          exit = true
      }
    }
    if exit{
      break
    }
  }
  
  return c
}

func (c Cliente)Visualizar(){
  fmt.Printf("%sID:%s %d\n", BOLD, END, c.Id)
  fmt.Printf("%sNombre:%s %s\n", BOLD, END, c.Nombre)
  fmt.Printf("%sTeléfono:%s %d\n", BOLD, END, c.Telefono)
  fmt.Printf("%sEmail:%s %s\n", BOLD, END, c.Email)
  if (len(c.Vehiculos) > 0){
    for _, v := range c.Vehiculos{
      fmt.Printf("  %s·%s", BOLD, END)
      fmt.Println(v.Info())
    }
  } else {
    fmt.Println(BOLD, "SIN VEHÍCULOS", END)
  }
}

func (c Cliente)Modificar(t Taller) (Cliente){
  menu := []string{
    "Modificar datos de " + c.Nombre,
    "Id",
    "Nombre",
    "Teléfono",
    "Email",
    "Vehículos"}

  var aux_int int
  var aux_str string
  var exit bool = false

  for{
    menu[0] = "Modificar datos de " + c.Nombre
    opt, status := menuFunc(menu)
    if (status == 0){
      switch opt{
        case 1:
          leerInt(&aux_int)
          if (aux_int > 0){
            c.Id = aux_int
            infoMsg("Id actualizado")
          }
        case 2:
          leerStr(&aux_str)
          if (len(aux_str) > 0){
            c.Nombre = aux_str
            infoMsg("Nombre actualizado")
          }
        case 3:
          leerInt(&aux_int)
          if (aux_int > 0){
            c.Telefono = aux_int
            infoMsg("Telefono actualizado")
          }
        case 4:
          leerStr(&aux_str)
          if (len(aux_str) > 0){
            c.Email = aux_str
            infoMsg("Email actualizado")
          }
        case 5:
          c = c.MenuVehiculos(t)
        default:
          exit = true
      }
      
    }
    if exit{
      break
    }
  }

  return c
}

func (c Cliente)ListarVehiculos(){  

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

func (c Cliente)TieneVehiculo(v Vehiculo) (bool){
  if (len(c.Vehiculos) > 0){
    for _, vc := range c.Vehiculos{
      if vc.Comparar(v){
        return true
      }
    }
  }
  
  return false
}

type Vehiculo struct{
  Matricula string // 1324ACB
  Marca string
  Modelo string
  FechaEntrada string
  FechaSalida string
  Incidencias []Incidencia
}

func (v Vehiculo)MenuVehiculo() (Vehiculo){
  var exit bool = false

  for{
    titulo := "Menu del " + v.Marca + " " + v.Modelo
    menu := []string{titulo, "Visualizar", "Modificar", "Listar Incidencias"}
    opt, status := menuFunc(menu)

    if status == 0{
      switch opt{
        case 1:
          v.Visualizar()
        case 2:
          v = v.Modificar()
        case 3:
          v.ListarIncidencias()
        default:
          exit = true
      }
    }
    if exit{
      break
    }
  }

  return v
}

func (v Vehiculo)Info() (string){
  return v.Marca + " " + v.Modelo + " (" + v.Matricula + ")"
}

func (v Vehiculo)Visualizar(){
  fmt.Printf("%sMatrícula:%s %s\n", BOLD, END, v.Matricula)
  fmt.Printf("%sMarca:%s %s\n", BOLD, END, v.Marca)
  fmt.Printf("%sModelo:%s %s\n", BOLD, END, v.Modelo)
  fmt.Printf("%sFecha de entrada:%s %s\n", BOLD, END, v.FechaEntrada)
  fmt.Printf("%sFecha estimada de salida:%s %s\n", BOLD, END, v.FechaSalida)
  fmt.Printf("%sIncidencias%s\n", BOLD, END)
  if len(v.Incidencias) > 0{ 
    for _, i := range v.Incidencias{
      fmt.Printf("  %s·%s %s", BOLD, END, i.Descripcion)
    }
  } else {
    fmt.Println(BOLD, "SIN INCIDENCIAS", END)
  }
  fmt.Println()
}

func (v Vehiculo)Modificar() (Vehiculo){
  menu := []string{
    "Modificar datos",
    "Matricula",
    "Marca y modelo",
    "Fecha de entrada",
    "Fecha de salida",
    "Incidencias"}
  var exit bool = false
  var aux string

  for{
    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          leerStr(&aux)
          if len(aux) > 0{
            v.Matricula = aux
            infoMsg("Matrícula actualizada")
          }
        case 2:
          fmt.Println("Marca")
          leerStr(&aux)
          if len(aux) > 0{
            marca := aux
            fmt.Println("Modelo")
            leerStr(&aux)
            if len(aux) > 0{
              v.Marca = marca
              v.Modelo = aux
              infoMsg("Marca y modelo actualizados")
            }
          }
        case 3:
          leerFecha(&aux)
          if len(aux) > 0{
            v.FechaEntrada = aux
            infoMsg("Fecha de entrada actualizada")
          }
        case 4:
          leerFecha(&aux)
          if len(aux) > 0{
            v.FechaSalida = aux
            infoMsg("Fecha de salida actualizada")
          }
        case 5:
          // Incidencias
        default:
          exit = true
      }
    }
    if exit{
      break
    }
  }

  return v
}

func (v Vehiculo)Comparar(aux Vehiculo) (bool){
  return v.Matricula == aux.Matricula
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
  Descripcion string
  Estado int // (0) Cerrada, (1) Abierta, (2) En proceso
}

func (i Incidencia)Visualizar(){
  fmt.Printf("%sId:%s %d\n", BOLD, END, i.Id)
  //fmt.Printf("%sMecánicos:%s %s\n", BOLD, END, i.Mecanicos)
  fmt.Printf("%sTipo de incidencia:%s %s\n", BOLD, END, i.ObtenerTipo())
  fmt.Printf("%sPrioridad:%s %s\n", BOLD, END, i.ObtenerPrioridad())
  fmt.Printf("%sDescripción:%s %s\n", BOLD, END, i.Descripcion)
  fmt.Printf("%sEstado:%s %s\n", BOLD, END, i.ObtenerEstado())
}

func (i Incidencia)ObtenerTipo() (string){
  switch i.Tipo{
    case 0:
      return "Mecánica"
    case 1:
      return "Eléctrica"
    case 2:
      return "Carrocería"
    default:
      return "Otro"
  }
}

func (i Incidencia)ObtenerPrioridad() (string){
  switch i.Prioridad{
    case 1:
      return "1 (Alta)"
    case 2:
      return "2 (Media)"
    case 3:
      return "3 (Baja)"
    default:
      return "Sin prioridad"
  }
}

func (i Incidencia)ObtenerEstado() (string){
  switch i.Estado{
    case 0:
      return "Cerrado"
    case 1:
      return "Abierto"
    case 2:
      return "En proceso"
    default:
      return "Desconocido"
  }
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

func leerFecha(aux *string){
  var dia int
  var mes int
  var anyo int

  for{
    fmt.Println("Día")
    leerInt(&dia)
    fmt.Println("Mes")
    leerInt(&mes)
    fmt.Println("Año")
    leerInt(&anyo)
    
    if (dia > 0 && dia <= 31 && mes > 0 && mes <= 12 && anyo > 0){
      *aux = fmt.Sprintf("%d-%d-%d", dia, mes, anyo)
      return
    } else if (dia == 0 && mes == 0 && anyo == 0){
      return
    }
  }
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

func crearCliente(t Taller) (Cliente){
  var c Cliente
  var exit bool = false

  fmt.Printf("%sID%s\n", BOLD, END)
  leerInt(&c.Id)
  fmt.Printf("%sNombre%s\n", BOLD, END)
  leerStr(&c.Nombre)
  fmt.Printf("%sTeléfono%s\n", BOLD, END)
  leerInt(&c.Telefono)
  fmt.Printf("%sEmail%s\b\n", BOLD, END)
  leerStr(&c.Email)

  menu := []string{"Vehículos", "Nuevo vehículo"}
  for{
    opt, status := menuFunc(menu)

    if status == 0{
      if opt == 1 && !t.EstaLleno(){
        v := crearVehiculo(t)
        c.Vehiculos = append(c.Vehiculos, v)
      } else {
        exit = true
      }
    }
    if exit{
      break
    }
  }

  return c
}

func crearVehiculo(t Taller) (Vehiculo){
  var v Vehiculo
  var exit bool = false

  fmt.Printf("%sMatrícula%s\n", BOLD, END)
  leerStr(&v.Matricula)
  fmt.Printf("%sMarca%s\n", BOLD, END)
  leerStr(&v.Marca)
  fmt.Printf("%sModelo%s\n", BOLD, END)
  leerStr(&v.Modelo)
  fmt.Printf("%sFecha de entrada%s\n", BOLD, END)
  leerStr(&v.FechaEntrada)
  fmt.Printf("%sFecha esperada de salida%s\n", BOLD, END)
  leerStr(&v.FechaSalida)

  menu := []string{"Incidencias", "Nueva Incidencia"}
  for{
    opt, status := menuFunc(menu)

    if status == 0{
      if opt == 1{
        i := crearIncidencia(t)
        v.Incidencias = append(v.Incidencias, i)
      } else {
        exit = true
      }
    }
    if exit{
      break
    }
  }

  return v
}

func crearIncidencia(t Taller) (Incidencia){
  var i Incidencia
  menu_tipo := []string{"Tipo", "Mecánica", "Eléctrica", "Carroceria"}
  menu_prioridad := []string{"Prioridad", "Baja", "Media", "Alta"}

  fmt.Printf("%sID%s\n", BOLD, END)
  leerInt(&i.Id)
  
  for{
    opt, status := menuFunc(menu_tipo)
    if status == 0 && opt < 4{
      i.Tipo = opt - 1
      break
    }
  }

  for{
    opt, status := menuFunc(menu_prioridad)
    if status == 0 && opt < 4{
      i.Prioridad = opt
      break
    }
  }

  fmt.Printf("%sDescripción%s\n", BOLD, END)
  leerStr(&i.Descripcion)
  i.Estado = 1 // Abierta

  return i
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

  leerInt(&opt)

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

  // INICIALIZAR
  v := Vehiculo{
    Matricula:"1234ABC",
    Marca: "Dacia",
    Modelo: "Sandero",
    FechaEntrada: "14-10-2025",
    FechaSalida: "09-11-2025"}
  c := Cliente{
    Id: 12345678,
    Nombre: "Pepito",
    Telefono: 12345678,
    Email: "pepito@webmail.org"}
  m := Mecanico{
    Id: 12345678,
    Nombre: "Luis Martinez",
    Especialidad: 1,
    Experiencia: 4,
    Alta: true}
  i := Incidencia{
    Id: 1,
    Tipo: 0,
    Prioridad: 2,
    Descripcion: "Cambio de embrague",
    Estado: 1}

  i.Mecanicos = append(i.Mecanicos, m)
  v.Incidencias = append(v.Incidencias, i)
  c.Vehiculos = append(c.Vehiculos, v)
  t.Clientes = append(t.Clientes, c)
  t.Mecanicos = append(t.Mecanicos, m)
  // FIN INICIALIZAR

  var exit bool = false

  for{
    opt, status := menuFunc(mainMenu)
    if status == 0 { // 0: Ok, 1: Error
      switch opt{
        case 1:
          t.MenuPrincipal()
        case 2:
          t = t.MenuClientes()
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
