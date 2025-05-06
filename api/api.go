package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// Estructura para mapear la respuesta JSON
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
    url := "https://jsonplaceholder.typicode.com/posts/1"

    // Hacer la petición GET
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error en la petición:", err)
        return
    }
    defer resp.Body.Close()

    // Leer el cuerpo de la respuesta
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error al leer la respuesta:", err)
        return
    }

    // Convertir JSON a estructura Post
    var post Post
    if err := json.Unmarshal(body, &post); err != nil {
        fmt.Println("Error al parsear JSON:", err)
        return
    }

    // Mostrar resultado
    fmt.Printf("Título: %s\nCuerpo: %s\n", post.Title, post.Body)
}
