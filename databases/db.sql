DROP TABLE IF EXISTS Books CASCADE;
DROP TABLE IF EXISTS Authors CASCADE;
DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS Borrows CASCADE;

CREATE TABLE Books(
  id BIGSERIAL PRIMARY KEY,
  author_id BIGINT NOT NULL,
  title VARCHAR UNIQUE NOT NULL,
  description VARCHAR NOT NULL,
  quantity BIGINT NOT NULL,
  cover VARCHAR,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

CREATE TABLE Authors(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

CREATE TABLE Users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,  
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

CREATE TABLE Borrows (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL, 
    borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP,  
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

ALTER TABLE Books 
ADD FOREIGN KEY (author_id) REFERENCES Authors(id);

ALTER TABLE Borrows 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE Borrows 
ADD FOREIGN KEY (book_id) REFERENCES Books(id);

INSERT INTO Authors (name, created_at, updated_at)
VALUES
('Messi', '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Ronaldo',  '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Kaka', '2022-07-01 15:00:00', '2022-07-01 15:00:00');

INSERT INTO users (name, email, password, created_at, updated_at)
VALUES
('Ariana Grande', 'ariana@gmail.com', '$2a$10$4Pfk3cP5S.1iNZLGq6odL.wFz20HqQqm4e.b76ueGq8YPBJbiwLBO' , '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Margot Robbie', 'margot@gmail.com', '$2a$10$4Pfk3cP5S.1iNZLGq6odL.wFz20HqQqm4e.b76ueGq8YPBJbiwLBO' , '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Emma Watson', 'emma@gmail.com', '$2a$10$4Pfk3cP5S.1iNZLGq6odL.wFz20HqQqm4e.b76ueGq8YPBJbiwLBO' ,'2022-07-01 15:00:00', '2022-07-01 15:00:00');

INSERT INTO Books (title, author_id, description, quantity, cover, created_at, updated_at)
VALUES
('Laskar Pelangi', 1, 'Buku ini bagus bre, beli aja lahhh', 2, 'Hardcover', '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Frieren', 2, 'My lord mbak frierennnnn', 1, 'Softcover', '2022-07-01 15:00:00', '2022-07-01 15:00:00'),
('Buku Kunci Sukses', 3, 'Ayo ambil kunci untuk menuju jalan kesuksesan', 3, 'Softcover', '2022-07-01 15:00:00', '2022-07-01 15:00:00');
