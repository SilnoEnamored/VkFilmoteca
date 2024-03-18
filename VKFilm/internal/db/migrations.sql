CREATE DATABASE filmoteca;

CREATE TABLE actors (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        gender VARCHAR(50),
                        birth_date DATE NOT NULL
);

CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(150) NOT NULL,
                        description TEXT,
                        release_date DATE NOT NULL,
                        rating DECIMAL(2, 1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE movie_actors (
                              movie_id INT NOT NULL,
                              actor_id INT NOT NULL,
                              PRIMARY KEY (movie_id, actor_id),
                              FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE,
                              FOREIGN KEY (actor_id) REFERENCES actors (id) ON DELETE CASCADE
);
