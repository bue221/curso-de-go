-- Usar la base de datos (ya creada por docker-compose)
USE db_contact;

-- Crear la tabla contacts
CREATE TABLE IF NOT EXISTS contacts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insertar algunos datos de ejemplo
INSERT INTO contacts (name, email, phone) VALUES 
('Juan Pérez', 'juan@email.com', '123-456-7890'),
('María García', 'maria@email.com', '098-765-4321'),
('Carlos López', 'carlos@email.com', '555-123-4567');
