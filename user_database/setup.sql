CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    username VARCHAR (50) UNIQUE NOT NULL, 
    password VARCHAR (100) NOT NULL, 
    first_name VARCHAR (50),
    last_name VARCHAR (50),
    date_of_birth DATE,
    email VARCHAR (320),
    phone_number VARCHAR (50)
);
