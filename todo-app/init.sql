CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       title VARCHAR(100) NOT NULL,
                       des TEXT,
                       status VARCHAR(20) DEFAULT 'pending',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       finished_at TIMESTAMP DEFAULT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE

);

INSERT INTO users (name, email) VALUES ('FIRST_USER', 'firstemail@gmail.com');
INSERT INTO tasks (user_id, title, des) VALUES (1, 'FIRST_TASK', 'This is the first task description');
