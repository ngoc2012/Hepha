-- Check if the database exists
USE INFORMATION_SCHEMA;
SELECT SCHEMA_NAME
FROM SCHEMATA
WHERE SCHEMA_NAME = 'hepha_db';

-- If the database doesn't exist, create it
DROP DATABASE IF EXISTS hepha_db;
CREATE DATABASE IF NOT EXISTS hepha_db;

-- Use the newly created database
USE hepha_db;

-- Create the tables
-- CREATE TABLE IF NOT EXISTS system (
--   text_id INT AUTO_INCREMENT PRIMARY KEY,
--   code_id INT AUTO_INCREMENT PRIMARY KEY
-- );

CREATE TABLE System (
    id INT AUTO_INCREMENT PRIMARY KEY,
    param VARCHAR(255),
    val INT
);

CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    login VARCHAR(255),
    name VARCHAR(255),
    secret_2fa VARCHAR(255),
    session_id VARCHAR(255)
);

-- protection INT CHECK(protection IN (0, 1, 2)),
CREATE TABLE Sheet (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    owner VARCHAR(255),
    protection INT,
    searchable BOOLEAN,
    exe_times INT,
    comments INT,
    evaluations INT,
    star FLOAT,
    created_date DATETIME,
    mod_date DATETIME
);

-- type VARCHAR(15) CHECK(type IN ('text', 'code', 'input', 'output', 'equation', 'image', 'graph')),
CREATE TABLE Box (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sheet_id INT,
    position INT,
    type VARCHAR(15),
    box_id INT,
    caption VARCHAR(255),
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Code (
    id INT AUTO_INCREMENT PRIMARY KEY,
    parent_id INT,
    language VARCHAR(63),
    language_version VARCHAR(63),
    executions INT,
    errors INT,
    last_error_date DATETIME,
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Input (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    code_file_id INT,
    error BOOLEAN
);

CREATE TABLE Output (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    code_file_id INT,
    error BOOLEAN
);

CREATE TABLE Text (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Equation (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Image (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Graph (
    id INT AUTO_INCREMENT PRIMARY KEY,
    file_id INT,
    created_date DATETIME,
    mod_date DATETIME
);

CREATE TABLE Share (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sheet_id INT,
    shared_user VARCHAR(255),
    privilege VARCHAR(63)
);