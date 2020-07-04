package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Estructura de persona
type Person struct {
	Nombre       string `json:"nombre,omitempty"`
	Departamento string `json:"departamento,omitempty"`
	Edad         int    `json:"edad,omitempty"`
	Estado       string `json:"estado,omitempty"`
	Contagio     string `json:"forma de contagio,omitempty"`
}

//arreglo con el listado de personas
var people []Person

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("1. Ingrese la URL: ")
	scanner.Scan()
	strURL := scanner.Text()
	fmt.Printf("2. Ingrese Cantidad de hilos para enviar: ")
	scanner.Scan()
	intHilos, _ := strconv.Atoi(scanner.Text())
	fmt.Printf("3. Ingrese la Cantidad de datos: ")
	scanner.Scan()
	intCantidad, _ := strconv.Atoi(scanner.Text())
	fmt.Printf("4. Ingrese la ruta del archivo a cargar: ")
	scanner.Scan()
	strArchivo := scanner.Text()

	fmt.Printf("URL: %s HILOS: %d CANTIDAD: %d ARCHIVO: %s", strURL, intHilos, intCantidad, strArchivo)

	// Open our jsonFile
	jsonFile, err := os.Open(strArchivo)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Archivo leido satisfactoriamente")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &people)

	for i := 0; i < len(people); i++ {
		//fmt.Println("Nombre : " + people[i].Nombre)
		//fmt.Println("Estado: " + people[i].Estado)
		//fmt.Println("Edad: " + strconv.Itoa(people[i].Edad))
		//fmt.Println("Departamento: " + people[i].Departamento)
		//fmt.Println("TipoContagio: " + people[i].Contagio)
		jsonReq, err := json.Marshal(people[i])
		resp, err := http.Post(strURL, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		// Convert response body to string
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
