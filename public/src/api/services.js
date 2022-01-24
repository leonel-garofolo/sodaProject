import Commons from './commons';

var responseData
export default class Services extends Commons{
    constructor() {  
        super()              
    }

    getData(){
        return responseData
    }

    async getDeliveries() {
        await super.get("deliveries", null)            
            .then(data => responseData = data)        
    }

    getDeliveryClient() {
        super.get("delivery", null).then(response =>{
            console.log(response)
        })
    }
}


