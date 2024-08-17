-- Check if the database exists
SELECT datname FROM pg_database WHERE datname = 'hepha_db';

-- If the database doesn't exist, create it
DROP DATABASE IF EXISTS hepha_db;
CREATE DATABASE hepha_db;

-- Connect to the newly created database
\c hepha_db;

-- Create the tables
CREATE TABLE System (
    id SERIAL PRIMARY KEY,
    param VARCHAR(255),
    val INT
);

CREATE TABLE "User" (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255),
    name VARCHAR(255),
    secret_2fa VARCHAR(255),
    session_id VARCHAR(255)
);

CREATE TABLE Sheet (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    owner VARCHAR(255),
    protection INT CHECK(protection IN (0, 1, 2)),
    searchable BOOLEAN,
    exe_times INT,
    comments INT,
    evaluations INT,
    star FLOAT,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Box (
    id SERIAL PRIMARY KEY,
    sheet_id INT,
    position INT,
    type VARCHAR(15) CHECK(type IN ('text', 'code', 'input', 'output', 'equation', 'image', 'graph')),
    box_id INT,
    caption VARCHAR(255),
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Code (
    id SERIAL PRIMARY KEY,
    parent_id INT,
    language VARCHAR(63),
    language_version VARCHAR(63),
    executions INT,
    errors INT,
    last_error_date TIMESTAMP,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Input (
    id SERIAL PRIMARY KEY,
    file_id INT,
    code_file_id INT,
    error BOOLEAN
);

CREATE TABLE Output (
    id SERIAL PRIMARY KEY,
    file_id INT,
    code_file_id INT,
    error BOOLEAN
);

CREATE TABLE Text (
    id SERIAL PRIMARY KEY,
    file_id INT,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Equation (
    id SERIAL PRIMARY KEY,
    file_id INT,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Image (
    id SERIAL PRIMARY KEY,
    file_id INT,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Graph (
    id SERIAL PRIMARY KEY,
    file_id INT,
    created_date TIMESTAMP,
    mod_date TIMESTAMP
);

CREATE TABLE Share (
    id SERIAL PRIMARY KEY,
    sheet_id INT,
    shared_user VARCHAR(255),
    privilege VARCHAR(63)
);
