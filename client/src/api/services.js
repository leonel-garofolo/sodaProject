import Commons from './commons';

var responseData= []
export default class Services extends Commons{    
    constructor() {  
        super()              
    }

    resetResponse(){
        responseData= {
            errors: [],
            data: []
        }
    }

    getResponse(){
        return responseData
    }

    async getDeliveries() {
        this.resetResponse()
        await super.get("delivery", null)  
            .then(function(response){
                return response.json()
            })          
            .then(data => this.reloadResponse(data, null)) 
            
    }

    async getDeliveriesCode() {
        this.resetResponse()
        await super.get("deliveryCode", null)            
            .then((response)=>{
                if(response.ok){
                    return response.json()    
                }
                return response.text().then(text => { throw new Error(text) })
            })          
            .then(data => this.reloadResponse(data, null)) 
            .catch(error => {
                console.log(error)
            })
    }

    async getDeliveryClient(code) {
        this.resetResponse()
        await super.get("deliveriesToClient?cod="+code, null)  
            .then((response)=>{
                return response.json()
            })                    
            .then(data => this.reloadResponse(data, null)) 
    }

    async getReport(code) {
        this.resetResponse()
        window.open().location.href = this.getENDPOINT() + "report?id="+code;
    }

    async saveClient(client){
        this.resetResponse()
        await super.post("client", client)
            .then(function(response) {                
                if(response.ok){
                    return response.json()
                }                
                return response.text().then(text => { throw new Error(text) })
            })
            .then(data => this.reloadResponse(data, null))
            .catch(error => {                
                this.reloadResponse(null, JSON.parse(error.message))
            });                    
    }

    reloadResponse(data, errors){
        responseData = {
            data: data != null? data: [],
            errors: errors != null? errors: []
        }        
    }    
}


