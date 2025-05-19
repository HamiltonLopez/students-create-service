package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/students-create-service/models"
)

// ---------------------- MOCK DEL SERVICIO ----------------------

type MockStudentService struct {
    mock.Mock
}




func (m *MockStudentService) AddStudent(student models.Student) (*models.Student, error) {
	args := m.Called(student)
	return args.Get(0).(*models.Student), args.Error(1)
}


// ---------------------- TESTS ----------------------

func TestCreateStudent_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    student := models.Student{
        Name:  "Juan",
        Age:   22,
        Email: "juan@test.com",
    }

    body, _ := json.Marshal(student)
    req := httptest.NewRequest("POST", "/students", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    mockService.On("AddStudent", mock.MatchedBy(func(s models.Student) bool {
        return s.Name == student.Name && s.Age == student.Age && s.Email == student.Email
    })).Return(&models.Student{
        ID:    primitive.NewObjectID(),
        Name:  student.Name,
        Age:   student.Age,
        Email: student.Email,
    }, nil)
    
    controller.CreateStudent(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
    mockService.AssertExpectations(t)
}

func TestCreateStudent_InvalidData(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    req := httptest.NewRequest("POST", "/students", bytes.NewReader([]byte("invalid json")))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    controller.CreateStudent(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}



