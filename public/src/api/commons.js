const END_POINT ='http://localhost:3000/';

export default class Commons {
    constructor(){}

    async get(serviceName, request){
        var init = {
            method:"GET"            
        }
        if(request != null){
            init.body = request
        }      
        const response =await fetch(END_POINT + serviceName, init)  
        return response.json()
    }    
}

