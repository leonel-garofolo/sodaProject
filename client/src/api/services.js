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
        await super.get("delivery", null)            
            .then(data => this.reloadData(data)) 
    }

    async getDeliveryClient() {
        responseData = []
        await super.get("deliveriesToClient", null)
            .then(data => this.reloadData(data)) 
    }

    reloadData(data){
        (data != null? responseData = data:responseData= [])
    }
}


