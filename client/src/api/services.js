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

    async getDeliveriesCode() {
        responseData = []
        await super.get("deliveryCode", null)            
            .then(data => this.reloadData(data)) 
    }

    async getDeliveryClient(code) {
        responseData = []
        await super.get("deliveriesToClient?cod="+code, null)            
            .then(data => this.reloadData(data)) 
    }

    reloadData(data){
        (data != null? responseData = data:responseData= [])
    }
}


