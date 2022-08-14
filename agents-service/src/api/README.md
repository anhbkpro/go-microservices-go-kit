# go-kit Practice

- Practice go-kit from the [video](https://www.youtube.com/watch?v=a462f8NvUvI&ab_channel=PacktVideo)
  - `cd agents-service/src/api/`, `go build main.go`, copy `main` to `agents-service` folder
  - `docker buildx build --platform=linux/amd64 -t agents-service . -f Dockerfile`
- *Issue*: Running this code on Mac M1, I do not find a way to run `dockerx` on docker-compose, and connection issue to the database and run `CMD
to run Go service via Docker image. So I run service from code directly, create MariaDB using below steps:

- [Update 1]: Now I can create database via docker-compose:
<img width="1871" alt="image" src="https://user-images.githubusercontent.com/2937629/184545475-b9d6e340-5dfc-4705-8c31-53bee45b5dcf.png">


- Create [MariaDB](https://mariadb.com/resources/blog/get-started-with-mariadb-using-docker-in-3-steps/#:~:text=Execute%20the%20following%20to%20connect,mariadb%20%2D%2Duser%20root%20%2DpPassword123!&text=And%20that's%20it!,start%20using%20(querying)%20MariaDB.):
  - Create MariaDB Docker Container: `docker run -p 127.0.0.1:3306:3306  --name mdb -e MARIADB_ROOT_PASSWORD=Password123! -d mariadb:latest`
  - Connect to MariaDB: `docker exec -it mdb mariadb --user root -pPassword123!`
  - create database: `create database agentsdb;`
  - connect to the database: `use agentsdb;`
  - create `manager` table: 
    ```sql
    create table manager(
    id int auto_increment,
    manager varchar(255) not null,
    account int,
    created_at timestamp default current_timestamp,
    primary key(id)
    );
    ```
  - create `manager_player` table:
    ```sql
    create table manager_player(
    id int auto_increment,
    manager_id int,
    player_id int,
    created_at timestamp default current_timestamp,
    primary key(id)
    );
    ```
  - insert data to `manager` table:
    ```sql
    insert into manager(manager, account, created_at)
    values('John',10, current_date);
    ```
- Test Go service with Postman:
  - POST: http://127.0.0.1:8080/agent-player
     ```json lines
     {
       "agent_id": 10,
       "player_id": 2022
     }
     ```
  - GET: http://127.0.0.1:8080/agent/1
