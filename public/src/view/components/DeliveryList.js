import React from "react";
import Form from 'react-bootstrap/Form'
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";


export default class DeliveryList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            deliveries: [
                {
                    id:"1",
                    name:"Anselmo"
                },
                {
                    id:"2",
                    name:"Ariel"
                }
            ],
        }        
    }

    render(){
        return (
            <Container>
                <Row>
                    <Col sm={2} ><label>Seleccione el Reparto</label></Col>
                    <Col sm={2}>
                        <Form.Select>
                            {
                                this.state.deliveries.map((deliveries, i) => {     
                                    return (<option id={deliveries.id}>{deliveries.name}</option>) 
                                })
                            }
                        </Form.Select>
                    </Col>
                </Row>                
            </Container>   
        )
    }
}
