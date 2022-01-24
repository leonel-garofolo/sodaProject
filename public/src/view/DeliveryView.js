import React from "react";
import Services from '../api/services';
const s = new Services()

export default class DeliveryView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            deliveries: [],
        }
        this.callService()        
    }

    async callService(){
       await s.getDeliveries()        
       this.setState({                
        deliveries: s.getData()
       })
    }
    

    render() {    
        if(this.state.deliveries.length > 0) {
            return (
                <div>
                    <h1>Repartidores</h1>
                    {
                        this.state.deliveries.map((delivery, i) => {     
                            return (<p id={delivery.id}>{delivery.name}</p>) 
                        })
                    }                    
                </div>
            );        
        } else {
            return (
                <div>
                    <h3>Cargando...</h3>
                </div>
            );        
        }       
    }
}