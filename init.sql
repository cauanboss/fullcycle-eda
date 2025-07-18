-- Script de inicialização do banco de dados
-- Verifica e cria as tabelas se não existirem

USE wallet;

-- Criar tabela clients se não existir
CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL
);

-- Criar tabela accounts se não existir
CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    balance DECIMAL(10,2) DEFAULT 0,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- Criar tabela transactions se não existir
CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    account_from_id VARCHAR(255) NOT NULL,
    account_to_id VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (account_from_id) REFERENCES accounts(id),
    FOREIGN KEY (account_to_id) REFERENCES accounts(id)
);

-- Criar tabela balances se não existir
CREATE TABLE IF NOT EXISTS balances (
    id VARCHAR(255) PRIMARY KEY,
    account_id VARCHAR(255) NOT NULL,
    balance DECIMAL(10,2) NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);

-- Verificar se as tabelas foram criadas
SHOW TABLES; 