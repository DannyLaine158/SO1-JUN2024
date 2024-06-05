# Módulos de Kernel

## Requisitos:
### Compilador de C
- gcc-12

### Creación de módulo
- touch file.c
- touch Makefile

### Compilación de módulo
- make

### Instalación de módulo en /proc
- sudo insmod file.ko

### Eliminación de módulo de /proc
- sudo rmmod file.ko
