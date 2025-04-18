CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       title VARCHAR(100) NOT NULL,
                       description TEXT,
                       status VARCHAR(20) DEFAULT 'pending',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);
