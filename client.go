package main

import (
	"fmt"
	"net/rpc"
	"strconv"
)

type Alumno struct {
	Nombre       string
	Calificacion float64
	Materia      string
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var op int64
	for {
		fmt.Println("1) Agregar calificación de un alumno por materia")
		fmt.Println("2) Obtener el promedio del alumno")
		fmt.Println("3) Obtener el promedio de todos los alumnos")
		fmt.Println("4) Obtener el promedio por materia")
		fmt.Println("0) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name string
			fmt.Print("Nombre: ")
			fmt.Scanln(&name)

			var ca string
			fmt.Print("Calificación: ")
			fmt.Scanln(&ca)

			var m string
			fmt.Print("Materia: ")
			fmt.Scanln(&m)

			al := Alumno{}
			al.Nombre = name
			i, e := strconv.ParseFloat(ca, 64)
			if e != nil {
				al.Calificacion = 0
			} else {
				al.Calificacion = i
			}

			al.Materia = m

			var result string
			err = c.Call("Server.AgregarCalificacion", al, &result)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Server.AgregarCalificacion =", result)
			}
		case 2:
			var name string
			fmt.Print("Nombre: ")
			fmt.Scanln(&name)

			var result string
			err = c.Call("Server.PromedioAlumno", name, &result)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Server.PromedioAlumno =", result)
			}
		case 3:
			var result string
			err = c.Call("Server.PromedioGeneral", "", &result)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Server.PromedioGeneral=", result)
			}
		case 4:
			var name string
			fmt.Print("Materia: ")
			fmt.Scanln(&name)

			var result string
			err = c.Call("Server.PromedioMateria", name, &result)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Server.PromedioMateria =", result)
			}
		case 0:
			return
		}
	}

	var result string
	al := Alumno{}
	al.Nombre = "Erick"
	al.Calificacion = 100
	al.Materia = "sepa"
	err = c.Call("Server.AgregarCalificacion", al, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Hello =", result)
	}
}

func main() {
	client()
}
