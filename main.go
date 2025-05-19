package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "example.com/students-create-service/controllers"
    "example.com/students-create-service/services"
    "example.com/students-create-service/repositories"
)

func init() {
    godotenv.Load()
}

func main() {
    repo := repositories.NewStudentRepository()
    service := services.NewStudentService(repo)
    controller := controllers.NewStudentController(service)
    
    r := mux.NewRouter()

    r.HandleFunc("/students", controller.CreateStudent).Methods("POST")

    fmt.Println("Servicio CREATE Students escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

