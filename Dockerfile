# Utilizamos una imagen liviana de Go como base
FROM golang:alpine

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el código fuente al directorio de trabajo
COPY . .

# Compilamos el código Go
RUN go build -o server .

# Exponemos el puerto 80 en el contenedor
EXPOSE 80

# Comando para ejecutar el servidor
CMD ["./server"]
