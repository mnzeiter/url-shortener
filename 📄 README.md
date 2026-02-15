# Mini URL Shortener – Go + Gin + MySQL

A simple URL shortening REST API built with Go (1.25), Gin, and MySQL.  
Users can shorten long URLs, get redirect links, and view basic statistics.

---

## Features

- Shorten long URLs into short codes
- Redirect from short URL to original URL
- Track number of clicks per short URL
- MySQL for persistent storage
- Dockerized for easy setup

---

## Tech Stack

- Go 1.25
- Gin Web Framework
- MySQL
- GORM
- Docker & Docker Compose

---

## Run with Docker

```bash
docker-compose up --build
