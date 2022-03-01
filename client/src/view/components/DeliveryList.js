import React from "react";
import Form from 'react-bootstrap/Form'
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import Services from "../../api/services";

const s = new Services()
export default class DeliveryList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {            
            deliveries: [],
            deliveryNameSelected:"",
        }
    }

    async componentWillMount() {          
        await s.getDeliveriesCode()        
        var tempDeliveries= s.getData()
        this.setState({                
            deliveries: tempDeliveries,
            deliveryNameSelected: (tempDeliveries.length > 0?tempDeliveries.name:"")
        })  
        if(this.state.deliveries.length > 0){
            this.props.onDeliverySelected(this.state.deliveries[0].code)
        }        
    }

    onChange(event){                
        this.setState({
            deliveries: this.state.deliveries,
            deliveryNameSelected: this.state.deliveries[event.target.selectedIndex].name
        })
        this.props.onDeliverySelected(event.target.value)        
    }
    
    render(){
        return (
            <Container>
                <Row>
                    <Col sm={2} ><label>Seleccione el Reparto</label></Col>
                    <Col sm={2}>
                        <Form.Select
                            onChange={this.onChange.bind(this)}
                            >
                            {
                                this.state.deliveries.map((deliveries, i) => {     
                                    return (<option key={i} id={i}>{deliveries.code}</option>) 
                                })
                            }
                        </Form.Select>
                    </Col>
                    <Col>
                        <Form.Label>{this.state.deliveryNameSelected}</Form.Label>
                    </Col>
                </Row>                
            </Container>   
        )
    }
}
