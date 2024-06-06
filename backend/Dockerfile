# Usar una imagen base de Go oficial
FROM golang:1.22

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el módulo go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente de la aplicación al contenedor
COPY . .

# Construir la aplicación Go
RUN go build -o main ./cmd/main.go

# Definir el puerto en el que la aplicación escucha (opcional)
EXPOSE 8080

# Definir el comando para ejecutar la aplicación
CMD ["./main"]
