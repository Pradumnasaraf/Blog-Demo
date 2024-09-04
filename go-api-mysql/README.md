## Golang API with MySQL 

A Restful API in Golang using Gin-Gonic framework and MySQL. This API allows you to create, read, update and delete schedules.

## Setting up the API

Once you clone the repository, copy the `.env.example` file to `.env` and update the values as per your MySQL configuration. Make sure your MySQL server is running and have the necessary permissions and the database created with the name you have provided in the `.env` file.

Once you are done with the configuration, you can run the following command to start the API:

```sh
go run main.go
```

This will start the API on `http://localhost:8080`.

Now you can use the API to interact with the MySQL database. See [Using the API](#using-the-api) to see the available endpoints.

## Using the API

Once you have the API running, you can use the following endpoints to interact with the API.
We have Get, Post, Patch and Delete methods to interact with the API.

| Method | Endpoint | Description | Body |
| --- | --- | --- | --- |
| GET | /health | Check the health of the API | None |
| GET | /api/schedules | Get all the schedules | None |
| POST | /api/schedules | Create a new schedule | `{"content": "Schedule content"}` |
| PATCH | /api/schedules/:id | Update a schedule | `{"content": "Updated schedule content"}` |
| DELETE | /api/schedules/:id | Delete a schedule | None |
