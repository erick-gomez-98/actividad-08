package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {
	materias map[string]map[string]float64
	alumnos  map[string]map[string]float64
}

type Alumno struct {
	Nombre       string
	Calificacion float64
	Materia      string
}

func (this *Server) AgregarCalificacion(al Alumno, reply *string) error {
	*reply = "Calificación agregada."

	if a, ok := this.alumnos[al.Nombre]; ok {
		// Ya existe, buscar la materia
		if _, o := a[al.Materia]; o {
			// Ya existe la materia, entonces no hay que permitir guardar
			return errors.New("Ya existe esa calificación para ese alumno y esa materia")
		}

		// No existe esa materia en ese alumno, agregarla
		this.alumnos[al.Nombre][al.Materia] = al.Calificacion
	} else {
		// No existe el alumno, generarlo
		materia := make(map[string]float64)
		materia[al.Materia] = al.Calificacion
		this.alumnos[al.Nombre] = materia
	}

	if _, ok := this.materias[al.Materia]; ok {
		this.materias[al.Materia][al.Nombre] = al.Calificacion
	} else {
		alumno := make(map[string]float64)
		alumno[al.Nombre] = al.Calificacion
		this.materias[al.Materia] = alumno
	}
	return nil
}

func (this *Server) PromedioAlumno(nombre string, reply *string) error {
	if al, ok := this.alumnos[nombre]; ok {
		// Si existe
		var sum float64
		for _, calificacion := range al {
			sum += calificacion
		}
		*reply = fmt.Sprint("El promedio de "+nombre+" es ", (sum / float64(len(al))))
		return nil
	}
	return errors.New("No existe ningun alumno con ese nombre")
}

func (this *Server) PromedioGeneral(_ string, reply *string) error {
	var sum float64
	for _, materias := range this.alumnos {
		var promedioAlumno float64
		for _, calificacion := range materias {
			promedioAlumno += calificacion
		}
		promedioAlumno /= float64(len(materias))
		sum += promedioAlumno
	}
	*reply = fmt.Sprint("El promedio general es ", (sum / float64(len(this.alumnos))))
	return nil
}

func (this *Server) PromedioMateria(materia string, reply *string) error {
	if al, ok := this.materias[materia]; ok {
		// Si existe
		var sum float64
		for _, calificacion := range al {
			sum += calificacion
		}
		*reply = fmt.Sprint("El promedio de la materia de "+materia+" es ", (sum / float64(len(al))))
		return nil
	}
	return errors.New("No existe ninguna materia con ese nombre")
}

func server() {
	s := new(Server)
	s.materias = make(map[string]map[string]float64)
	s.alumnos = make(map[string]map[string]float64)
	rpc.Register(s)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
