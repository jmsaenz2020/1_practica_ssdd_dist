package main

import (
  "fmt"
)

const PLAZAS_MECANICO = 2
const BOLD = "\033[1;37m"
const RED = "\033[1;31m"
const YELLOW = "\033[1;33m"
const BLUE = "\033[1;34m"
const END = "\033[0m"

type Taller struct{
  Clientes []Cliente
  Plazas *[]Vehiculo
  Mecanicos []Mecanico
  UltimoId int
}

func (t *Taller)Menu(){
  menu := []string{
    "Menu del taller",
    "Listar incidencias",
    "Listar vehiculos",
    "Listar incidencias de mecánico",
    "Mecánicos disponibles"}

  for{
    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          // Listar Incidencias
        case 2:
          // Listar Vehiculos
        case 3:
          // Listar por mecánico
        case 4:
          t.MecanicosDisponibles()
        default:
          continue
      }
    } else if status == 2{
      break
    }
  }
}

func (t *Taller)MenuMecanicos(){
  var menu []string

  if len(t.Mecanicos) > 0{
    for{
      menu = []string{"Selecciona un mecánico"}
      for _, m := range t.Mecanicos{
        menu = append(menu, m.Info())
      }

      opt, status := menuFunc(menu)
      
      if status == 0{
        if opt > 0 && opt <= len(t.Mecanicos) - 1{
          t.Mecanicos[opt - 1].Menu()
        }
      } else if status == 2{
        break
      }
    }
  } else {
    warningMsg("No hay mecánicos en el taller")
  }
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

func (t *Taller)EliminarMecanico(m Mecanico){
  
  indice := t.ObtenerIndiceMecanico(m)
    
  if indice >= 0{ // Eliminar
    lista := t.Mecanicos
    lista[indice] = lista[len(lista) - 1]
    t.Mecanicos = lista[:len(lista) - 1]
  } else {
    errorMsg("No se pudo eliminar al mecánico")
  }
}

func (t Taller)ObtenerIndiceMecanico(m_in Mecanico) (int){
  var res int = -1

  for i, m := range t.Mecanicos{
    if m.Igual(m_in){
      res = i
    }
  }

  return res
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

func (t Taller)ObtenerMecanicosDisponibles() ([]Mecanico){
  var mecanicos []Mecanico  

  for _, m := range t.Mecanicos{
    if m.Alta{
      mecanicos = append(mecanicos, m)
    }
  }

  return mecanicos
}

func (t Taller)MecanicosDisponibles(){
  for _, m := range t.Mecanicos{
    if m.Alta{
      fmt.Println(m.Info())
    }
  }
}

func (t *Taller)ModificarMecanico(modif Mecanico){
  for i, m := range t.Mecanicos{
    if m.Igual(modif){
      t.Mecanicos[i] = modif
    }
  }
}


type Cliente struct{

}


type Vehiculo struct{

}


type Incidencia struct{

}


type Mecanico struct{
  Id int
  Nombre string
  Especialidad int // Mecanica, Electrica, Carroceria
  Experiencia int
  Alta bool
}

func (m *Mecanico)Menu(){
  
}

func (m Mecanico)Info() (string){
  return fmt.Sprintf("%s (%03d)", m.Nombre, m.Id)
}

func (m Mecanico)Visualizar(){
  fmt.Printf("%sID: %s%03d\n", BOLD, END, m.Id)
  fmt.Printf("%sNombre: %s%s\n", BOLD, END, m.Nombre)
  fmt.Printf("%sEspecialidad: %s%s\n", BOLD, END, m.ObtenerEspecialidad())
  fmt.Printf("%sExperiencia: %s%d años\n", BOLD, END, m.Experiencia)
  fmt.Printf("%s¿Está de alta? %s%t\n", BOLD, END, m.Alta)
}

func (m1 Mecanico)Igual(m2 Mecanico) (bool){
  return m1.Id == m2.Id
}

func (m *Mecanico)Modificar(){
  menu := []string{
    "Modificar datos de cliente",
    "Nombre",
    "Especialidad",
    "Experiencia",
    "Dar de baja"}
  var buf string
  var num int

  for{
    if !m.Alta{
      menu[len(menu) - 1] = "Dar de alta"
    } else {
      menu[len(menu) - 1] = "Dar de baja"
    }
    menu[0] = fmt.Sprintf("Modificar datos de %s", m.Nombre)
    opt, status := menuFunc(menu)
    if status == 0{
      switch opt{
        case 1:
          leerStr(&buf)
          m.Nombre = buf
          infoMsg("Nombre modificado")
        case 2:
          menu_esp := []string{
            "Selecciona especialidad",
            "Mecánica",
            "Electrónica",
            "Carrocería"}
          opt, status = menuFunc(menu_esp)
          if status == 0{
            esp := m.ObtenerEspecialidad()
            m.Especialidad = opt - 1
            msg := fmt.Sprintf("Especialidad modificada: %s->%s", esp, m.ObtenerEspecialidad())
            infoMsg(msg)
          }
        case 3:
          leerInt(&num)
          m.Experiencia = num
          infoMsg("Experiencia modificada")
        case 4:
          m.Alta = !m.Alta
          infoMsg("Estado modificado")
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
  fmt.Printf("%s%s%s\n", RED, msg, END)
}

func warningMsg(msg string){
  fmt.Printf("%s%s%s\n", YELLOW, msg, END)
}

func infoMsg(msg string){
  fmt.Printf("%s%s%s\n", BLUE, msg, END)
}

func leerInt(i *int){
  for{
    fmt.Print("> ")
    fmt.Scanf("%d", i)
    if *i >= 0{
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
    if len(*str) > 0{
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
  
  menu := []string{
    "Menu principal",
    "Taller",
    "Clientes",
    "Vehiculos",
    "Incidencias",
    "Mecánicos"}

  t.CrearMecanico("Pepe", 0, 0)

  for{
    opt, status := menuFunc(menu)
    
    if status == 0{
      switch opt{
        case 1:
          t.Menu()
        case 2:
          // Clientes
        case 3:
          // Vehiculos
        case 4:
          // Incidencias
        case 5:
          t.MenuMecanicos()
      }
    } else if status == 2{
      break
    }
  }
  
}
