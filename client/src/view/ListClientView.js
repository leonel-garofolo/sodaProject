import React from "react";
import Button from "react-bootstrap/esm/Button";
import Services from "../api/services";
import DataTable from "./table/DataTable";
import DeliveryList from "./components/DeliveryList";

import { Row, Col, Divider } from 'antd';


const s = new Services()

export default class ListClientView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clients: [],
        }
        this.onDeliverySelected = this.onDeliverySelected.bind(this)
    }

    async print() {
        await s.getReport("10")
        console.log(s.data)
        /*
        const file = new Blob([s.data], { type: "application/pdf" });
        const fileURL = URL.createObjectURL(file);
        const pdfWindow = window.open();
        pdfWindow.location.href = fileURL;          
        */
    }

    async onDeliverySelected(code) {
        console.log("onDeliverySelected", code)
        await s.getDeliveryClient(code)
        this.setState({
            clients: s.getResponse().data
        })
    }

    render() {
        return (
            <>
                <Row>
                    <Col span={18}>
                        <DeliveryList onDeliverySelected={this.onDeliverySelected} />
                    </Col>
                    <Col span={6}>
                        <div style={{ float: "right", padding: '0 20px' }}>
                            <Button onClick={this.print}>Imprimir</Button>
                        </div>
                    </Col>
                </Row>
                <Divider orientation="left">Clientes</Divider>
                <DataTable data={this.state.clients} />
            </>
        )
    }
}