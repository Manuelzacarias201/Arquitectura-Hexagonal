-- Script para agregar solo la tabla users a una base de datos existente
-- Ejecuta este script si ya tienes las tablas alumns y teachers creadas

USE school;

-- Tabla de Usuarios (para autenticación)
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Índice para mejorar el rendimiento en búsquedas por email
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
