# Soda 
Soda is a system to manager delivery of SODA to little company.

## Features
- Show delivery
- Show Clients and update data client

## Enviroment
**Develop**

Run local API server:
```
 go run .\main.go 
 ```

Or, Run de mock API service to consume from React client, you have to move to public folder and run the follow command: 
``` 
json-server --watch test/mock.json  
```

After that, up the React Client with command:
``` 
npm start   
```

**Production**

(Building)




