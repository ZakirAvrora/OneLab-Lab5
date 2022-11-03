# OneLab-Lab5
Simple CRUD server that stores books

## Author
@ZakirAvrora

## Installation

```bash
git clone https://github.com/ZakirAvrora/OneLab-Lab4
```

## Usage

- To __build__ docker container use command:
```bash
make build
```
or directly docker commands:
```bash
docker build --rm -t crud-web .
docker image prune --filter label=stage=builder -f
```
- To __run__ docker container use:
```bash
make run 
```
or directly docker command:
```bash
docker run --rm --name crud-web -p 8080:8080 crud-web
```