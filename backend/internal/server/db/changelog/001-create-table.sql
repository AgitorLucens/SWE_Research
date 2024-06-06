CREATE TABLE IF NOT EXISTS record (
    id SERIAL PRIMARY KEY,
    record_name VARCHAR(100) NOT NULL,
    img TEXT NOT NULL, 
    price NUMERIC(6,2) NOT NULL,
    descr VARCHAR(200) NOT NULL, 
    quant int NOT NULL,
    topic VARCHAR(100) NOT NULL,
    created DATE NOT NULL,
    published timestamp DEFAULT NOW() NOT NULL)