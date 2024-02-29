# Usa una imagen de Python liviana como base
FROM python:3.9-slim

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el código Python al contenedor
COPY . .

# Instala las dependencias necesarias, incluyendo aiohttp
RUN pip install --no-cache-dir aiohttp

# Establece la variable de entorno SERVER_URL
ENV SERVER_URL=http://localhost

# Expone el puerto 8080 para que pueda ser accesible desde otro contenedor
EXPOSE 8080

# Ejecuta el script Python
CMD ["python", "client.py"]

#docker build -t python-client .
#docker network create my-network
#docker network connect my-network go-server
#docker run --name python-client --network my-network -p 8080:80 -e SERVER_URL=http://go-server:80 python-client



