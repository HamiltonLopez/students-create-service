package controllers

import (
    "encoding/json"
    "net/http"
    "example.com/students-create-service/models"
    "example.com/students-create-service/services"
)

type StudentController struct {
    Service *services.StudentService
}

func NewStudentController(service *services.StudentService) *StudentController {
    return &StudentController{
        Service: service,
    }
}

func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    err := json.NewDecoder(r.Body).Decode(&student)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    newStudent, err := c.Service.AddStudent(student)
    if err != nil {
        http.Error(w, "Error al insertar estudiante", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newStudent)
}

