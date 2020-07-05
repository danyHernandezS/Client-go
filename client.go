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

//Person definition
type Person struct {
	Nombre       string `json:"nombre,omitempty"`
	Departamento string `json:"departamento,omitempty"`
	Edad         int    `json:"edad,omitempty"`
	Estado       string `json:"estado,omitempty"`
	Contagio     string `json:"forma de contagio,omitempty"`
}

//arreglo con el listado de personas
var people []Person

//These function will receive work on the jobs channel and send the corresponding results on results.
func enviarData(id int, strURL string, jobs <-chan int, results chan int) {

	for i := range jobs {
		if len(people) > 0 {
			fmt.Println("worker", id, "started  job", i)
			//read the data in the first position
			jsonReq, err := json.Marshal(people[0])

			//make a slice to do a dequeue kind of operation
			people = people[1:]

			//send the data taked and POST it to the url API
			resp, err := http.Post(strURL, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
			if err != nil {
				log.Fatalln(err)
				return
			}

			//close connection
			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)

			// Convert response body to string
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
			fmt.Println("worker", id, "finished  job", i)
		}
		results <- i * 2
	}
}

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

	// Open our jsonFile
	jsonFile, err := os.Open(strArchivo)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &people)

	// make a chanel to indicate the number of jobs to run
	jobs := make(chan int, intCantidad)
	// collect all the results of the work.
	results := make(chan int, intCantidad)

	// starts up "w" workers, initially blocked because there are no jobs yet.
	for w := 1; w <= intHilos; w++ {

		go enviarData(w, strURL, jobs, results)
	}
	//send  jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= intCantidad; j++ {
		jobs <- j
	}
	close(jobs)
	//collect all the results of the work. This also ensures that the worker goroutines have finished
	for a := 1; a <= intCantidad; a++ {
		<-results
	}
}
