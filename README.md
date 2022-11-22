# go-svelte-example
I have implemented a simple go backend api and svlete frontend for books CRUD. I have used the postres database for storing the data. The **init.sql** file contain the books and users tables create query. Running the docker compose up command will create the table from the init.sql file.
### Setting up DB
To create the database and the books table, navigate to go-api directory and run the following docker-compose command.
```sh
docker-compose up -d
```
Navigate to go-api folder and copy .env.example file content to .env file and update the connection string according to your variables.


