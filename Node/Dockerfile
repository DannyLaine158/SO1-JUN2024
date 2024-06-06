# Desde que lenguaje o framework queremos trabajar
FROM node:20-alpine

# Crear un directorio de trabajo
WORKDIR /api

# Copiar los archivos necesarios para el código
COPY package.json ./
COPY package-lock.json ./

# Ejecutar el comando para obtener las librerías
RUN npm install

# Copiar el resto del código
COPY . ./

# Ejecutar el comando para lanzar la aplicación
CMD ["node", "index.js"]
