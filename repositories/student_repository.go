package repositories

import (
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "example.com/students-create-service/models"
    "log"
    "os"  
)

type StudentRepository struct {
    collection *mongo.Collection
}

func NewStudentRepository() *StudentRepository {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("school").Collection("students")
    return &StudentRepository{collection}
}

func (repo *StudentRepository) CreateStudent(student models.Student) (*models.Student, error) {
    student.ID = primitive.NewObjectID()
    _, err := repo.collection.InsertOne(context.TODO(), student)
    if err != nil {
        return nil, err
    }
    return &student, nil
}

