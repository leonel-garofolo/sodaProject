import Commons from './commons';

var responseData = []
export default class Services extends Commons{
    constructor() {  
        super()              
    }

    getData(){
        return responseData
    }

    async getDeliveries() {
        responseData = []
        await super.get("deliveries", null)            
            .then(data => responseData = data)        
    }

    async getDeliveryClient() {
        responseData = []
        await super.get("clients", null)
            .then(data => responseData = data) 
    }
}


