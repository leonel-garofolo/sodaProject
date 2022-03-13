const END_POINT ='http://localhost:3000/app/';

export default class Commons {
    constructor(){}

    getENDPOINT(){
        return END_POINT
    }

    async get(serviceName, data){
        var init = {
            method:"GET"            
        }
        if(data != null){
            init.body = JSON.stringify(data)
        }          
        return await fetch(END_POINT + serviceName, init)  
    }
    
    async post(serviceName, data){
        var init = {
            method:"POST", 
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'                
            },         

        }
        if(data != null){
            init.body = JSON.stringify(data)
        }                
        return await fetch(END_POINT + serviceName, init)  
    } 
}

