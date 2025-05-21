# Students Create Service

Servicio responsable de la creación de nuevos estudiantes en el sistema.

## Funcionalidad

Este servicio expone un endpoint POST que permite crear nuevos registros de estudiantes en la base de datos MongoDB.

## Especificaciones Técnicas

- **Puerto**: 8081 (interno), 30081 (NodePort)
- **Endpoint**: POST `/students`
- **Runtime**: Go
- **Base de Datos**: MongoDB

## Estructura del Servicio

```
students-create-service/
├── k8s/
│   ├── deployment.yaml
│   └── service.yaml
├── src/
│   ├── main.go
│   ├── handlers/
│   ├── models/
│   └── config/
├── Dockerfile
└── README.md
```

## API Endpoint

### POST /students

Crea un nuevo estudiante en el sistema.

#### Request Body
```json
{
    "name": "string",
    "age": number,
    "email": "string"
}
```

#### Response
```json
{
    "id": "string",
    "name": "string",
    "age": number,
    "email": "string",
    "created_at": "timestamp"
}
```

## Configuración Kubernetes

### Deployment
- **Replicas**: 3
- **Imagen**: hamiltonlg/students-create-service:latest
- **Variables de Entorno**:
  - MONGO_URI: mongodb://mongo-service:27017

### Service
- **Tipo**: NodePort
- **Puerto**: 8081 -> 30081

## Despliegue

```bash
kubectl apply -f k8s/
```

## Verificación

1. Verificar el deployment:
```bash
kubectl get deployment students-create-deployment
```

2. Verificar los pods:
```bash
kubectl get pods -l app=students-create
```

3. Verificar el servicio:
```bash
kubectl get svc students-create-service
```

## Pruebas

### Crear un nuevo estudiante
```bash
curl -X POST http://localhost:30081/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "age": 20,
    "email": "juan@example.com"
  }'
```

## Logs

Ver logs de un pod específico:
```bash
kubectl logs -f <pod-name>
```

## Monitoreo

### Métricas Importantes
- Tiempo de respuesta del endpoint
- Tasa de éxito/error en creación
- Uso de recursos (CPU/Memoria)

## Solución de Problemas

1. **Error de Conexión a MongoDB**:
   - Verificar la variable MONGO_URI
   - Comprobar conectividad con mongo-service
   - Revisar logs de MongoDB

2. **Pod en CrashLoopBackOff**:
   - Verificar logs del pod
   - Comprobar recursos asignados
   - Verificar configuración del deployment

3. **Servicio no accesible**:
   - Verificar el estado del service
   - Comprobar la configuración de NodePort
   - Verificar reglas de firewall 