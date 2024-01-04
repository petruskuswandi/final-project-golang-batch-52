-- migrate create -ext sql -dir database/migrations create_users_table
CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    born_date TIMESTAMP
);

-- RUN MIGRATE
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/final_project_golang_batch_52" -path database/migrations up  

-- DROP TABLE
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/final_project_golang_batch_52" -path database/migrations down