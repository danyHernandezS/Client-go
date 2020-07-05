# Client-go

Este programa lleva a cabo la lectura de los siguientes parámetros de entrada:

1. URL a la cual se enviarán las peticiones

2. Cantidad de hilos que se abriran para trabajar con concurrencia

3. Cantidad de información que se enviará

4. Nombre del archivo que tiene toda la data

Este programa utilza el objeto *scanner* de la libreria **"bufio"** para tomar cada parámetro de entrada de la consola y almacenarla.

Depues con la libreria **"os"** se lee el archivo de entrada especificado en los parámetros y se parsea su contenido para después convertir su contenido en un array de jsons con la libreria **"enconding/json"**.

Se hizo uso de una arquitectura "Working Pool" para realizar la concurrencia en el envío de la data a la REST API que se especifica.

Esta arquitectura hace uso de *Gorrutines* la cual abre un número de workers de acuerdo a la cantidad de hilos especificados en el parámetro de entrada y va tomando el primer dato del array con la data que actua como una cola FIFO y realiza el envío a la API viía POST.
