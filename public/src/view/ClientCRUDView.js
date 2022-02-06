import React from "react";
import Services from "../api/services";
import Form from 'react-bootstrap/Form'
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import Button from 'react-bootstrap/Button';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import DeliveryList from "./components/DeliveryList";

const s = new Services()

export default class ClientCRUDView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clients: [],
        }        
    }

    async componentWillMount() {
        await s.getDeliveryClient()        
        this.setState({                
            clients: s.getData()
        })
    }

    goBefore(){
        alert("goBefore")
    }

    save(){
        alert("Save")
    }

    cancel(){
        alert("Cancel")
    }

    goAfter(){
        alert("goAfter")
    }

    render() { 
        if(this.state.clients.length > 0) {
            return (
                <Container className="p-3">    
                    <Row>
                        <DeliveryList />
                    </Row>  
                    <Row>
                        <Col><label>Orden</label></Col>                    
                    </Row>         
                    <Row>
                        <Col><input type="text" id="orden" /></Col>
                    </Row>
                    <Row><Col><label>Precio por Sifon</label></Col></Row>         
                    <Row><Col><input type="text" id="pricePerSoda" /></Col></Row>
                    <Row><Col><label>Dirección</label></Col></Row>         
                    <Row><Col><input type="text" id="address" /></Col></Row>
                    <Row><Col><label>Numero</label></Col></Row>         
                    <Row><Col><input type="text" id="numberAddress" /></Col></Row>
                    <Row><Col><label>Deuda</label></Col></Row>         
                    <Row><Col><input type="text" id="debt" /></Col></Row>
                    <Row>
                        <ButtonGroup aria-label="Basic example">
                            <Button variant="outline-info" onClick={this.goBefore}>Anterior</Button>
                            <Button variant="outline-danger" onClick={this.cancel}>Cancelar</Button>
                            <Button variant="outline-success" onClick={this.save}>Guardar</Button>
                            <Button variant="outline-primary" onClick={this.goAfter}>Siguiente</Button>
                        </ButtonGroup>                
                    </Row>                
                </Container>
            );        
        } else {
            return (
                <div>
                    <h3>Cargando</h3>                        
                </div>
            );        
        }
    }
}