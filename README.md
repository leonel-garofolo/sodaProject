# Soda
Soda is a system to manager delivery of SODA to little company.

## Features
- Show delivery
- Show Clients and update data client

## Libreries used
- [GoReport](https://github.com/mikeshimura/goreport)


## Enviroment
Install docker compose from https://docs.docker.com/compose/install/

**Develop**
First, you have to up the docker container that have the mysql database.

```
 docker-compose up
 ```

After that, you need to run the Go API server:

```
 go run .\app.go
 ```

**Test**

Run de mock API service to consume from React client, you have to move to public folder and run the follow command:
```
json-server --watch test/mock.json  
```

After that, up the React Client with command:
```
npm start   
```

**Commandos**
For help the migration database from xBase to Mysql we have a command client to do this proccess more easy.

You have to execute:
```
soda migrate <name_file>.dbx
```


**Production**
Execute docker commands to up three containers.

```
docker compose --profile prd up
```

