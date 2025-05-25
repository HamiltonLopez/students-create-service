#!/bin/bash

# Colores para la salida
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# URL del servicio
SERVICE_URL="http://${KUBE_IP}:30081"

echo "Probando Students Create Service..."
echo "=================================="

# Test 1: Crear un estudiante válido
echo -e "\nTest 1: Crear un estudiante válido"
response=$(curl -s -X POST \
  "${SERVICE_URL}/students" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "age": 20,
    "email": "juan@example.com"
  }')

if [[ $response == *"id"* ]]; then
    echo -e "${GREEN}✓ Test 1 exitoso: Estudiante creado correctamente${NC}"
    # Extraer el ID para usarlo en otras pruebas
    STUDENT_ID=$(echo $response | grep -o '"id":"[^"]*' | cut -d'"' -f4)
    echo "ID del estudiante creado: $STUDENT_ID"
else
    echo -e "${RED}✗ Test 1 fallido: No se pudo crear el estudiante${NC}"
    echo "Respuesta: $response"
fi

# Test 2: Intentar crear un estudiante con datos inválidos
echo -e "\nTest 2: Crear un estudiante con datos inválidos"
response=$(curl -s -X POST \
  "${SERVICE_URL}/students" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "age": -1,
    "email": "invalid-email"
  }')

if [[ $response == *"error"* ]]; then
    echo -e "${GREEN}✓ Test 2 exitoso: El servicio rechazó correctamente los datos inválidos${NC}"
else
    echo -e "${RED}✗ Test 2 fallido: El servicio aceptó datos inválidos${NC}"
fi

# Test 3: Crear un estudiante con campos faltantes
echo -e "\nTest 3: Crear un estudiante con campos faltantes"
response=$(curl -s -X POST \
  "${SERVICE_URL}/students" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez"
  }')

if [[ $response == *"error"* ]]; then
    echo -e "${GREEN}✓ Test 3 exitoso: El servicio detectó campos faltantes${NC}"
else
    echo -e "${RED}✗ Test 3 fallido: El servicio aceptó datos incompletos${NC}"
fi

echo -e "\nPruebas completadas!" 