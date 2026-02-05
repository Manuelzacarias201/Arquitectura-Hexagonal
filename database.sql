-- Script de creación de base de datos para la API
-- Base de datos: school

-- Crear la base de datos (si no existe)
CREATE DATABASE IF NOT EXISTS school;
USE school;

-- Tabla de Alumnos
CREATE TABLE IF NOT EXISTS alumns (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    matricula VARCHAR(100)
);

-- Tabla de Profesores
CREATE TABLE IF NOT EXISTS teachers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    asignature VARCHAR(50)
);

-- Tabla de Usuarios (para autenticación)
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Índices para mejorar el rendimiento
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
