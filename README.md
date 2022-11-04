# OneLab-Lab5
Simple CRUD server that stores books

## Author
@ZakirAvrora

## Installation

```bash
git clone https://github.com/ZakirAvrora/OneLab-Lab5
```

## Usage

- Firstly to __build__ postgres docker container use command:
```bash
make postgres:
```
or directly docker commands:
```bash
docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
```
- Then create database docker_api using:
```bash
make createdb 
```
or directly docker command:
```bash
docker exec -it postgres15 createdb --username=root --owner=root books_api
```
- Run application on port 8000:
```bash
go run ./cmd/main.go
```
