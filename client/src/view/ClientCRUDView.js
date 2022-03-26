import React, { useEffect, useState } from "react";
import Services from "../api/services";
import DeliveryList from "./components/DeliveryList";
import { Row, Col, Divider, Form, Input, Button, Space, InputNumber, message } from 'antd';

const validateMessages = {
    required: '${label}: es requerido!',
    string: {
        range: '${label}: valores entre ${min} y ${max}',
    },
    number: {
        range: '${label}: valores entre ${min} y ${max}',
    },
};

const s = new Services()
var newClientData = {
    id: -1,
    order: 0,
    address: "",
    num_address: "",
    price_per_soda: "",
    price_per_box: "",
    debt: 0,
    id_delivery: -1,
    id_root: -1
}


export default class ClientCRUDView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            clients: [],
            currentClient: newClientData,
            currentErrors: [],
            index: 0,
        }

        this.onDeliverySelected = this.onDeliverySelected.bind(this)
        this.goBefore = this.goBefore.bind(this)
        this.newClient = this.newClient.bind(this)
        this.save = this.save.bind(this)
        this.goAfter = this.goAfter.bind(this)
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
        console.log("client", this.state.currentClient)
        await s.saveClient(this.state.currentClient)
        if (s.getResponse().errors.length > 0) {                      
            message.error('Oops! No se guardo el cliente, revisa los errores!')
            this.setState({
                currentErrors: s.getResponse().errors
            })
        } else {            
            message.success('El cliente se guardo correctamente!')
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

    async onDeliverySelected(deliverySelected) {
        await s.getDeliveryClient(deliverySelected.code)
        var response = s.getResponse()
        if (response.errors.length === 0 && response.data.length > 0) {
            this.setState({
                clients: response.data,
                index: 0,
                currentClient: response.data[this.state.index],
            })            
        }
        newClientData.id_delivery = deliverySelected.id
        newClientData.id_root = deliverySelected.code
    }

    render() {
        return (
            <>
                <Row>
                    <DeliveryList onDeliverySelected={this.onDeliverySelected} />
                </Row>
                <Divider orientation="left">Detalle del Cliente</Divider>
                <ClientForm currentClient={this.state.currentClient} currentErrors={this.state.currentErrors} />
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
            </>
        )
    }
}

function ClientForm(props) {
    const [form] = Form.useForm()
    initFormValues(form, props.currentClient)
    form.setFields(errorHandling(props.currentErrors))
    return (
        <Form
            name="userForm"
            layout="horizontal"
            form={form}
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            autoComplete="off"
            validateMessages={validateMessages}
        >
            <Row>
                <Col span={6}>
                    <Form.Item label="Orden" name="order" rules={[{ type: 'number', min: 1, max: 9999 }]} >
                        <InputNumber id="order" onChange={(value) => props.currentClient.order = value} />
                    </Form.Item>
                </Col>
            </Row>
            <Row>
                <Col span={6}>
                    <Form.Item label="Dirección" name="address" rules={[{ type: 'string', min: 3, max: 200 }]} >
                        <Input id="address" onChange={(e) => props.currentClient.address = e.target.value} />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="Número" name="numberAddress" rules={[{ type: 'number', min: 1, max: 99999 }]}>
                        <InputNumber id="numberAddress" onChange={(value) => props.currentClient.num_address = value} />
                    </Form.Item>
                </Col>
            </Row>
            <Row>
                <Col span={6}>
                    <Form.Item label="Precio por Sifon" name="pricePerSoda" rules={[{ type: 'number', min: 0, max: 99999 }]}>
                        <InputNumber id="pricePerSoda" onChange={(value) => props.currentClient.price_per_soda = value} type={"number"} />
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="Precio por Cajon" name="pricePerBox" rules={[{ type: 'number', min: 0, max: 99999 }]}>
                        <InputNumber id="pricePerBox" onChange={(value) => props.currentClient.price_per_box = value} type={"number"} />
                    </Form.Item>
                </Col>
            </Row>
            <Row>
                <Col span={6}>
                    <Form.Item label="Deuda" name="debt" rules={[{ type: 'number', min: 0, max: 999999 }]}>
                        <InputNumber id="debt" onChange={(value) => props.currentClient.debt = value} type={"number"} />
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    )
}

function initFormValues(form, currentClient){
    if(currentClient != null){
        form.setFieldsValue({
            order: currentClient.order,
        });
        form.setFieldsValue({
            address: currentClient.address,
        });
        form.setFieldsValue({
            numberAddress: currentClient.num_address,
        });
        form.setFieldsValue({
            pricePerSoda: currentClient.price_per_soda,
        });
        form.setFieldsValue({
            pricePerBox: currentClient.price_per_box,
        });
        form.setFieldsValue({
            debt: currentClient.debt,
        });       
    }
}

//example of Error Object: {FailedField: 'Client.Debt', Tag: 'required', Value: ''}
function errorHandling(currentErrors){
    var formErrors = []
    var i = 0;
    currentErrors.forEach(error => {
        console.log(error)
        switch (error.FailedField) {
            case 'Client.Order':                
                formErrors[i] = {
                        name: 'order',
                        errors: errorTextHandling(error),
                    }                
                break;
            case 'Client.Address':                
                formErrors[i] = {
                        name: 'address',
                        errors: errorTextHandling(error),
                    }                
                break;    
            case 'Client.NumAddress':                
                formErrors[i] = {
                        name: 'numberAddress',
                        errors: errorTextHandling(error),
                    }                
                break;
            case 'Client.PricePerSoda':                
                formErrors[i] = {
                        name: 'pricePerSoda',
                        errors: errorTextHandling(error),
                    }                
                break;            
            case 'Client.PricePerBox':                
                formErrors[i] = {
                        name: 'pricePerBox',
                        errors: errorTextHandling(error),
                    }                
                break;    
            case 'Client.Debt':                
                formErrors[i] = {
                        name: 'debt',
                        errors: errorTextHandling(error),
                    }                
                break;                
            default:
                break;
        }
        i++;
    });
    return formErrors
}

function errorTextHandling(error){
    var i = 0;
    var errorsString= []
    var requiredString = 'Campo requerido'
    var minString = 'El valor debe ser mayor a $1'
    var maxString = 'El valor debe ser menor a $1'
    if(error.Tag === 'required'){
        errorsString[i++] = requiredString
    }
    if(error.Tag === 'min'){
        errorsString[i++] = minString.replace("$1", error.Value)
    }
    if(error.Tag === 'max'){
        errorsString[i++] = maxString.replace("$1", error.Value)
    }
    return errorsString
}