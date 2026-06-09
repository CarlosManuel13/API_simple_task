# 🚀 Go Task Manager API

A minimalist RESTful API built with **Go** (Golang) for task management (To-Do list). This project uses the standard `net/http` package to demonstrate route handling, HTTP methods, and in-memory data persistence (slices).

## 📋 Features

- **GET `/tasks/`**: Retrieve the full list of tasks.
- **GET `/tasks/{id}`**: Search for a specific task by ID.
- **POST `/tasks/`**: Create a new task by sending JSON.
- **DELETE `/tasks/{id}`**: Remove a task from the system.

## 🛠️ Technologies & Structure

* **Language:** Go (Golang)
* **Architecture:** Based on a single controller using a `switch` statement for HTTP methods.
* **Data Format:** JSON (with custom tags).

## 🚀 How to Run the Server

1. Make sure Go is installed (`go version`).
2. Save the code in a file named `main.go`.
3. Run the following command in your terminal:
   ```bash
   go run main.go

## 🧪 Complete Guide to API Calls (cURL)

Use these commands in your terminal to test each server feature:

### 🔹 1. List all tasks (GET)
Obtains the complete array of JSON objects.
```bash
curl -i http://localhost:8000/tasks/
```

### 🔹 2. List a specific task (GET)
Lists a single item by filtering by ID
```bash
curl -i http://localhost:8000/tasks/1
```

### 🔹 3. Add a task (POST)
Add a task to the API using a request
```bash
curl -i -X POST http://localhost:8000/tasks/ \
-H "Content-Type: application/json" \
-d '{"id": "3", "title": "Stretch abdomen and correct posture", "completed": false}'
```

### 🔹 4. Delete a task (DELETE)
Delete a task by filtering by ID
```bash
curl -i -X DELETE http://localhost:8000/tasks/1
```

# 🇪🇸 <img width="256" height="256" alt="image" src="https://github.com/user-attachments/assets/5221a5d5-c703-42b1-a6e3-ec6cc7ac6437" />


# 🚀 Go Task Manager API

Una API RESTful minimalista construida con **Go** (Golang) para la gestión de tareas (To-Do list). Este proyecto utiliza el paquete estándar `net/http` para demostrar el manejo de rutas, métodos HTTP y persistencia de datos en memoria (slices).

## 📋 Características

- **GET `/tasks/`**: Obtiene el listado completo de tareas.
- **GET `/tasks/{id}`**: Busca una tarea específica por su ID.
- **POST `/tasks/`**: Crea una nueva tarea enviando un JSON.
- **DELETE `/tasks/{id}`**: Elimina una tarea del sistema.

## 🛠️ Tecnologías y Estructura

* **Lenguaje:** Go (Golang)
* **Arquitectura:** Basada en un único controlador con `switch` para métodos HTTP.
* **Formato de datos:** JSON (con etiquetas personalizadas).



## 🚀 Cómo ejecutar el servidor

1. Asegúrate de tener Go instalado (`go version`).
2. Guarda el código en un archivo llamado `main.go`.
3. Ejecuta el siguiente comando en tu terminal:
   ```bash
   go run main.go
   ```

## 🧪 Guía Completa de Llamadas API (cURL)

Usa estos comandos en tu terminal para probar cada funcionalidad del servidor:

### 🔹 1. Listar todas las tareas (GET)
Obtiene el array completo de objetos JSON.
```bash
curl -i http://localhost:8000/tasks/1
```

### 🔹 2. Listar una tarea en especifico (GET)
Lista un unico item filtrando por el ID
```bash
curl -i http://localhost:8000/tasks/
```

### 🔹 3. Anadir una tarea (POST)
Anades una tarea a la api mediante un peticion
```bash
curl -i -X POST http://localhost:8000/tasks/ \
-H "Content-Type: application/json" \
-d '{"id": "3", "title": "Estirar abdomen y corregir postura", "completed": false}'
```

### 🔹 4. Eliminar una tarea (DELETE)
Eliminas una tarea filtrando por el ID
```bash
curl -i -X DELETE http://localhost:8000/tasks/1
```

