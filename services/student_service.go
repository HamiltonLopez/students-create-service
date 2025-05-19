package services

import (
    "example.com/students-create-service/models"
    "example.com/students-create-service/repositories"
)

type StudentServiceInterface interface {
    AddStudent(student models.Student) (*models.Student, error)
}

type StudentService struct {
    repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
    return &StudentService{repo}
}

func (s *StudentService) AddStudent(student models.Student) (*models.Student, error) {
    return s.repo.CreateStudent(student)
}

