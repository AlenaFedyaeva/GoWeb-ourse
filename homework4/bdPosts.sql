-- CREATE database posts;

DROP TABLE IF EXISTS posts;
CREATE TABLE posts(
  id SERIAL PRIMARY KEY,
  author VARCHAR(255),
  postText TEXT,
  title VARCHAR(255),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
 );
 
INSERT INTO posts(author, postText,title) 
 VALUES ("Petrov I I","some text1", "some title1"),
        ("Ouen I V","some tex2", "some title2");
 
SELECT * FROM posts.posts;

-- запрос на обновление
update posts set postText = "empty"
where id = 1;