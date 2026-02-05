# Guía de Autenticación - Login y Register

Esta guía contiene ejemplos de peticiones HTTP para el módulo de autenticación usando Insomnia.

**URL Base**: `http://localhost:8080`

---

## 🔐 Módulo de Autenticación (`/auth`)

### 1. Registrar un Nuevo Usuario
**POST** `http://localhost:8080/auth/register`

**Headers:**
```
Content-Type: application/json
```

**Body (JSON):**
```json
{
  "email": "usuario@ejemplo.com",
  "password": "password123",
  "name": "Juan Pérez"
}
```

**Respuesta Exitosa (201 Created):**
```json
{
  "message": "Usuario registrado exitosamente"
}
```

**Respuestas de Error:**

**Email ya registrado (400 Bad Request):**
```json
{
  "error": "el email ya está registrado"
}
```

**Contraseña muy corta (400 Bad Request):**
```json
{
  "error": "la contraseña debe tener al menos 6 caracteres"
}
```

**Campos faltantes (400 Bad Request):**
```json
{
  "error": "Datos inválidos: Key: 'Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

---

### 2. Iniciar Sesión (Login)
**POST** `http://localhost:8080/auth/login`

**Headers:**
```
Content-Type: application/json
```

**Body (JSON):**
```json
{
  "email": "usuario@ejemplo.com",
  "password": "password123"
}
```

**Respuesta Exitosa (200 OK):**
```json
{
  "message": "Login exitoso",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InVzdWFyaW9AZWplbXBsby5jb20iLCJleHAiOjE3MDcwMDAwMDAsImlhdCI6MTcwNzAwMDAwMCwiaXNzIjoiYXBpIn0.abc123...",
  "user": {
    "id": 1,
    "email": "usuario@ejemplo.com",
    "name": "Juan Pérez"
  }
}
```

**Respuestas de Error:**

**Credenciales inválidas (401 Unauthorized):**
```json
{
  "error": "credenciales inválidas"
}
```

**Campos faltantes (400 Bad Request):**
```json
{
  "error": "Datos inválidos: Key: 'Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

---

## 📋 Flujo Completo de Autenticación

### Paso 1: Registrar un Usuario
```
POST http://localhost:8080/auth/register
Body: {
  "email": "nuevo@usuario.com",
  "password": "miPassword123",
  "name": "Nuevo Usuario"
}
```

### Paso 2: Iniciar Sesión
```
POST http://localhost:8080/auth/login
Body: {
  "email": "nuevo@usuario.com",
  "password": "miPassword123"
}
```

**Respuesta:**
- Guarda el `token` recibido
- El token es válido por 24 horas

### Paso 3: Usar el Token (Futuro)
Para proteger otros endpoints, puedes usar el token en el header:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## 🔒 Seguridad

### Características Implementadas:

1. **Encriptación de Contraseñas**: 
   - Las contraseñas se encriptan con bcrypt antes de guardarse
   - Nunca se retorna la contraseña en las respuestas

2. **Tokens JWT**:
   - Tokens firmados con algoritmo HS256
   - Válidos por 24 horas
   - Contienen: user_id, email, exp, iat, iss

3. **Validaciones**:
   - Email único (no se pueden registrar dos usuarios con el mismo email)
   - Contraseña mínima de 6 caracteres
   - Validación de campos requeridos

---

## ⚠️ Validaciones y Reglas

### Register:
- ✅ Email es requerido y debe ser único
- ✅ Password es requerido (mínimo 6 caracteres)
- ✅ Name es requerido

### Login:
- ✅ Email es requerido
- ✅ Password es requerido
- ✅ Las credenciales deben ser correctas

---

## 🐛 Errores Comunes

### Error: "el email ya está registrado"
**Solución**: Usa un email diferente o intenta hacer login si ya tienes cuenta.

### Error: "credenciales inválidas"
**Solución**: Verifica que el email y contraseña sean correctos. Recuerda que las contraseñas son case-sensitive.

### Error: "la contraseña debe tener al menos 6 caracteres"
**Solución**: Usa una contraseña de al menos 6 caracteres.

### Error: "Error al cargar el archivo .env"
**Solución**: Asegúrate de que el archivo `.env` existe y contiene `JWT_SECRET`.

---

## 📝 Ejemplos de Prueba en Insomnia

### Colección de Ejemplos:

**1. Registrar Usuario de Prueba:**
```json
POST http://localhost:8080/auth/register
{
  "email": "test@test.com",
  "password": "test123",
  "name": "Usuario de Prueba"
}
```

**2. Login con Usuario de Prueba:**
```json
POST http://localhost:8080/auth/login
{
  "email": "test@test.com",
  "password": "test123"
}
```

**3. Intentar Registrar Email Duplicado:**
```json
POST http://localhost:8080/auth/register
{
  "email": "test@test.com",
  "password": "otraPassword",
  "name": "Otro Usuario"
}
```
*Debería retornar error: "el email ya está registrado"*

**4. Login con Credenciales Incorrectas:**
```json
POST http://localhost:8080/auth/login
{
  "email": "test@test.com",
  "password": "passwordIncorrecta"
}
```
*Debería retornar error: "credenciales inválidas"*

---

## 🔧 Configuración Requerida

### Variables de Entorno (.env):
```env
JWT_SECRET=tu-clave-secreta-super-segura-cambiar-en-produccion-2026
```

**Importante**: Cambia `JWT_SECRET` por una clave segura en producción.

### Base de Datos:
Asegúrate de ejecutar el script `database.sql` que incluye la tabla `users`:
```sql
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

## 📊 Estructura de Respuestas

### Register Response:
```json
{
  "message": "Usuario registrado exitosamente"
}
```

### Login Response:
```json
{
  "message": "Login exitoso",
  "token": "jwt_token_aqui",
  "user": {
    "id": 1,
    "email": "usuario@ejemplo.com",
    "name": "Juan Pérez"
  }
}
```

---

## 🎯 Próximos Pasos

Para usar el token en otros endpoints protegidos, necesitarás:

1. Crear un middleware de autenticación
2. Validar el token JWT en cada petición protegida
3. Extraer el user_id del token para identificar al usuario

---

*Última actualización: 4 de febrero de 2026*
