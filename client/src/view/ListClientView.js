import React from "react";
import Button from "react-bootstrap/esm/Button";
import Services from "../api/services";
import DataTable from "./table/DataTable";
import Form from 'react-bootstrap/Form'
import Container from "react-bootstrap/esm/Container";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import DeliveryList from "./components/DeliveryList";


const s = new Services()

export default class ListClientView extends React.Component {
    constructor(props) {
        super(props);   
        this.state = {
            clients: [],
        }
        this.onDeliverySelected = this.onDeliverySelected.bind(this)
    }

    save(){

    }

    cancel(){

    }

    print(){

    }

    async onDeliverySelected(code){
        console.log("onDeliverySelected", code)        
        await s.getDeliveryClient(code)
        this.setState({
            clients: s.getData()
        })
    }

    render() {                 
        return (
            <Container>
                    <Row>
                        <Col >
                            <DeliveryList onDeliverySelected={this.onDeliverySelected}/>
                        </Col>                            
                        <Col xs lg="2">
                            <Button onClick={this.goBefore}>Imprimir</Button>
                        </Col>
                    </Row>
                    <Row>
                        <Col><DataTable data={this.state.clients}/></Col>                            
                    </Row>                        
            </Container> 
        )
    }
}