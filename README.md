# Students Create Service

Este servicio es parte del sistema de gestión de estudiantes y se encarga de la creación de nuevos registros de estudiantes.

## Estructura del Servicio

```
students-create-service/
├── controllers/     # Controladores REST
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuraciones de Kubernetes
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
└── test/          # Scripts de prueba
    └── test-create.sh
```

## Endpoints

### POST /create
Crea un nuevo estudiante en el sistema.

**Request Body:**
```json
{
    "name": "string",
    "age": number,
    "email": "string"
}
```

**Response (200 OK):**
```json
{
    "id": "string",
    "name": "string",
    "age": number,
    "email": "string"
}
```

## Configuración Kubernetes

### Deployment
El servicio se despliega con las siguientes especificaciones:
- Replicas: 1
- Puerto: 8080
- Imagen: students-create-service:latest

### Service
- Tipo: NodePort
- Puerto: 8080
- NodePort: 30081

### Ingress
- Path: /create
- Servicio: students-create-service
- Puerto: 8080

## Despliegue en Kubernetes

### 1. Aplicar configuraciones
```bash
# Crear el deployment
kubectl apply -f k8s/deployment.yaml

# Crear el service
kubectl apply -f k8s/service.yaml

# Crear el ingress
kubectl apply -f k8s/ingress.yaml
```

### 2. Verificar el despliegue
```bash
# Verificar el deployment
kubectl get deployment students-create-deployment
kubectl describe deployment students-create-deployment

# Verificar los pods
kubectl get pods -l app=students-create
kubectl describe pod -l app=students-create

# Verificar el service
kubectl get svc students-create-service
kubectl describe svc students-create-service

# Verificar el ingress
kubectl get ingress students-create-ingress
kubectl describe ingress students-create-ingress
```

### 3. Verificar logs
```bash
# Ver logs de los pods
kubectl logs -l app=students-create
```

### 4. Escalar el servicio
```bash
# Escalar a más réplicas si es necesario
kubectl scale deployment students-create-deployment --replicas=3
```

### 5. Actualizar el servicio
```bash
# Actualizar la imagen del servicio
kubectl set image deployment/students-create-deployment students-create=students-create-service:nueva-version
```

### 6. Eliminar recursos
```bash
# Si necesitas eliminar los recursos
kubectl delete -f k8s/ingress.yaml
kubectl delete -f k8s/service.yaml
kubectl delete -f k8s/deployment.yaml
```

## Pruebas

El servicio incluye un script de pruebas automatizadas (`test/test-create.sh`) que verifica:

1. Creación exitosa de un estudiante
2. Manejo de datos inválidos
3. Manejo de campos faltantes

Para ejecutar las pruebas:
```bash
./test/test-create.sh
```

También se puede ejecutar como parte de la suite completa de pruebas:
```bash
./test-all-services.sh
```

### Casos de Prueba

1. **Test 1:** Crear un estudiante válido
   - Envía datos completos y válidos
   - Verifica la respuesta exitosa y el ID generado

2. **Test 2:** Intentar crear con datos inválidos
   - Envía datos con formato incorrecto
   - Verifica el manejo apropiado de errores

3. **Test 3:** Crear con campos faltantes
   - Omite campos requeridos
   - Verifica la validación de campos obligatorios

## Variables de Entorno

- `MONGODB_URI`: URI de conexión a MongoDB (default: "mongodb://mongo-service:27017")
- `DATABASE_NAME`: Nombre de la base de datos (default: "studentsdb")
- `COLLECTION_NAME`: Nombre de la colección (default: "students")

## Dependencias

- Go 1.19+
- MongoDB
- Kubernetes 1.19+
- Ingress NGINX Controller

## Consideraciones de Seguridad

1. Validación de entrada de datos
2. Sanitización de datos
3. Manejo seguro de errores
4. Límites de tamaño en las solicitudes

## Monitoreo y Logs

- Endpoint de health check: `/health`
- Logs en formato JSON
- Métricas básicas de rendimiento

## Solución de Problemas

1. Verificar la conexión con MongoDB
2. Comprobar los logs del pod
3. Validar la configuración del Ingress
4. Verificar el estado del servicio en Kubernetes 