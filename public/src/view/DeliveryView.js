import React from "react";
import Services from '../api/services';
import Table from 'react-bootstrap/Table'

const s = new Services()

export default class DeliveryView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            deliveries: [],
        }
    }
    async componentWillMount() {
        await s.getDeliveries()   
        this.setState({                
            deliveries: s.getData()
        })
    }

    render() {    
        if(this.state.deliveries.length > 0) {
            return (
                <Table striped bordered hover responsive="sm" size="sm">
                    <thead>
                        <tr>
                        <th>#</th>
                        <th>Repartidor</th>                        
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.state.deliveries.map((delivery, i) => {     
                                return (<tr><td>{delivery.id}</td><td>{delivery.name}</td></tr>) 
                            })
                        } 
                    </tbody>
                </Table>
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