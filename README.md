# CRUD de Usuarios en Go 1.23
![Go Logo](https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

Este proyecto implementa un sistema básico de CRUD (Crear, Leer, Actualizar, Eliminar) para gestionar usuarios utilizando el lenguaje de programación Go. El propósito principal es demostrar cómo se puede estructurar un proyecto en Go para manejar operaciones básicas sobre una base de datos.

## Características

- Crear usuarios
- Leer información de usuarios
- Actualizar datos de los usuarios
- Eliminar usuarios

El proyecto está diseñado utilizando el enfoque MVC (Modelo, Vista, Controlador) para mantener una separación clara de responsabilidades. Además, se aprovechan plantillas HTML para las vistas y se hace uso de PostgreSQL como base de datos para persistencia de los datos.

## Estructura del Proyecto

- **/plantillas**: Contiene las plantillas HTML para las vistas.
- **/routes**: Define las rutas de la aplicación.
- **/services**: Contiene la lógica de negocio, incluida la gestión del CRUD de usuarios.
- **main.go**: Punto de entrada principal de la aplicación.

## Requisitos

- Go 1.23 o superior
- PostgreSQL

## Instrucciones de Instalación

1. Clona este repositorio.
    ```bash
    git clone https://github.com/AndresLondon/Crud-Go.git
    ```
2. Configura las credenciales de la base de datos en .env
3. Ejecuta los siguientes comandos.
    ```bash
    >>> go mod init main
    >>> go mod tidy 
    ```
4. Ejecuta el servidor Go.
    ```bash
    go run main.go
    ```


## Contribuciones

Las contribuciones son bienvenidas. Si deseas mejorar alguna funcionalidad o corregir errores, no dudes en enviar un mensaje al siguiente correo: moralondono3@gmail.com

## Licencia

Este proyecto está licenciado bajo la Licencia MIT.


## Derechos Reservados

Esta app, fue diseñada y desarrollada por el Porgramador Full Stack: Andres Cmailo Londoño Mora
