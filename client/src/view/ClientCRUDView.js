import React from "react";
import Services from "../api/services";
import DeliveryList from "./components/DeliveryList";
import { Row, Col, Divider, Form, Input, Button, Space, InputNumber } from 'antd';

const validateMessages = {
    required: '${label}: es requerido!',
    string:{
        range: '${label}: valores entre ${min} y ${max}',
    },
    number: {
        range: '${label}: valores entre ${min} y ${max}',
    },
};

const s = new Services()
const newClientData = {
    id: -1,
    order: 0,
    address: "",
    num_address: "",
    price_per_soda: "",
    price_per_box: "",
    debt: "",
    id_delivery: "",
    id_root: ""
}

export default class ClientCRUDView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clients: [],
            currentClient: newClientData,
            index: 0,
        }

        this.onDeliverySelected = this.onDeliverySelected.bind(this)
        this.goBefore = this.goBefore.bind(this)
        this.newClient = this.newClient.bind(this)
        this.save = this.save.bind(this)
        this.goAfter = this.goAfter.bind(this)
        this.onChangeData = this.onChangeData.bind(this)
    }

    goBefore() {
        if (this.state.index > 0) {
            const newIndex = this.state.index - 1
            this.setState({
                index: newIndex,
                currentClient: this.state.clients[newIndex],
            })
        }
    }

    async save() {
        await s.saveClient(this.state.currentClient)
        if (s.getResponse().errors.length > 0) {
            //TODO show error
        } else {
            alert("save is ok")
        }
    }

    newClient() {
        this.setState({
            clients: this.state.clients,
            index: -1,
            currentClient: newClientData,
        })
    }

    goAfter() {
        if (this.state.index < this.state.clients.length) {
            const newIndex = this.state.index + 1
            this.setState({
                clients: this.state.clients,
                index: newIndex,
                currentClient: this.state.clients[newIndex],
            })
        }
    }

    onChangeData(e) {
        var currentClientUpdate = this.state.currentClient
        switch (e.target.attributes[1].value) {
            case "orden":
                currentClientUpdate.order = e.target.value
                break;
            case "address":
                currentClientUpdate.address = e.target.value
                break;
            case "numberAddress":
                currentClientUpdate.num_address = e.target.value
                break;
            case "pricePerSoda":
                currentClientUpdate.price_per_soda = e.target.value
                break;
            case "pricePerBox":
                currentClientUpdate.price_per_box = e.target.value
                break;
            case "debt":
                currentClientUpdate.debt = e.target.value
                break;

            default:
                break;
        }

        this.setState({
            currentClient: currentClientUpdate,
        })
    }

    async onDeliverySelected(code) {
        await s.getDeliveryClient(code)
        var response = s.getResponse()
        console.log(response)
        if (response.errors.length == 0 && response.data.length > 0) {
            console.log("updateState")
            this.setState({
                clients: response.data,
                index: 0,
                currentClient: response.data[this.state.index],
            })
        }
    }

    render() {
        return (
            <>
                <Row>
                    <DeliveryList onDeliverySelected={this.onDeliverySelected} />
                </Row>
                <Divider orientation="left">Detalle del Cliente</Divider>
                <Form
                    name="basic"
                    layout="horizontal"
                    labelCol={{ span: 8 }}
                    wrapperCol={{ span: 16 }}
                    initialValues={{ remember: true }}
                    //   onFinish={onFinish}
                    //   onFinishFailed={onFinishFailed}
                    autoComplete="off"
                    validateMessages={validateMessages}
                >
                    <Row>
                        <Col span={6}>
                            <Form.Item label="Orden" name="Orden" rules={[{ type: 'number', min: 1, max: 9999 }]}>
                                <InputNumber id="orden" />
                            </Form.Item>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={6}>
                            <Form.Item label="Dirección" name="Dirección" rules={[{ type: 'string', min: 3, max: 200 }]}>
                                <Input id="address" type={"text"} />
                            </Form.Item>
                        </Col>
                        <Col span={6}>
                            <Form.Item label="Número" name="Número" rules={[{ type: 'number', min: 1, max: 99999 }]}>
                                <InputNumber id="numberAddress" />
                            </Form.Item>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={6}>
                            <Form.Item label="Precio por Sifon" name="Precio por Sifon" rules={[{ type: 'number', min: 0, max: 99999 }]}>
                                <InputNumber id="pricePerSoda" type={"number"} />
                            </Form.Item>
                        </Col>
                        <Col span={6}>
                            <Form.Item label="Precio por Cajon" name="Precio por Cajon" rules={[{ type: 'number', min: 0, max: 99999 }]}>
                                <InputNumber id="pricePerBox" type={"number"} />
                            </Form.Item>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={6}>
                            <Form.Item label="Deuda" name="Deuda" rules={[{ type: 'number', min: 0, max: 999999 }]}>
                                <InputNumber id="debt" type={"number"} />
                            </Form.Item>
                        </Col>
                    </Row>
                    <Row>
                        <Col span={12} offset={2}>
                            <Space>
                                <Button type="secondary" htmlType="button" onClick={this.goBefore}>Anterior</Button>
                                <Button type="primary" danger ghost htmlType="button" onClick={this.newClient}>Nuevo Cliente</Button>
                                <Button type="primary" htmlType="button" onClick={this.save}>Guardar</Button>
                                <Button type="secondary" htmlType="button" onClick={this.goAfter}>Siguiente</Button>
                            </Space>
                        </Col>
                    </Row>
                </Form>
            </>
        );
    }
}