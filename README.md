# Go JSON CRUD API

## Overview

JSON-CRUD-API-GO is a simple JSON-based CRUD API built in Go using the Gin web framework and GORM for database interactions. This project provides basic functionality for creating, reading, updating, and deleting posts, where each post consists of a title and body.

![Project Architecture](/go_json_crud_api.png)

## Project Architecture

The project follows a straightforward architecture with the following components:

- **API Server (Gin/GORM):** The main component handling HTTP requests and database interactions.
- **Routes (HTTP API):** Defines HTTP endpoints for CRUD operations.
- **CRUD Handlers:** Request handlers for creating, reading, updating, and deleting posts.
- **Database (GORM):** Manages data storage and database operations.
- **Data Storage:** Stores post data, including title and body.

## API Description

### Create a Post

- **Endpoint:** `/posts`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "Title": "Post Title",
    "Body": "Post Body"
  }
**Response:**
* Status Code: `200 (Success)`
* Data:
    ```json
    {
        "post": {
        "ID": 1,
        "Title": "Post Title",
        "Body": "Post Body"
        }
    }

### Fetch All Posts
- Endpoint: `/posts`
- Method: `GET`
  
**Response:**
* Status Code: `200 (Success)`
* Data:
    ```json
    {
        "posts": [
                    {
                        "ID": 1,
                        "Title": "Post Title 1",
                        "Body": "Post Body 1"
                    },
                    {
                        "ID": 2,
                        "Title": "Post Title 2",
                        "Body": "Post Body 2"
                    }
                ]
    }
### Fetch a Single Post
- Endpoint: `/posts/:id`
- Method: `GET`
  
**Response:**
* Status Code: `200 (Success)`
* Data:
    ```json
    {
        "post": {
                "ID": 1,
                "Title": "Post Title",
                "Body": "Post Body"
            }
    }
### Update a Post
- Endpoint: `/posts/:id`
- Method: `PUT`
**Request Body:**
    ```json
    {
        "Title": "Updated Title",
        "Body": "Updated Body"
    }
**Response:**
* Status Code: `200 (Success)`
* Data:
    ```json
    {
        "post": {
            "ID": 1,
            "Title": "Updated Title",
            "Body": "Updated Body"
        }
    }
### Delete a Post
- Endpoint: `/posts/:id`
- Method: `DELETE`
  
**Response:**
* Status Code: `200 (Success)`

Feel free to use this API for creating, reading, updating, and deleting posts. Make sure to customize it as needed for your own projects.
