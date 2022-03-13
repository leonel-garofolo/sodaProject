import React from "react";
import Services from "../../api/services";

import { Row, Col, Select, Space } from 'antd';

const { Option } = Select;

const s = new Services()
export default class DeliveryList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            deliveries: [],
            deliveryNameSelected: "",
        }
        this.onChange = this.onChange.bind(this)
    }

    async componentWillMount() {
        await s.getDeliveriesCode()
        var response = s.getResponse()
        console.log(response)
        this.setState({
            deliveries: response.data,
            deliveryNameSelected: (response.data.length > 0 ? response.data[0].name : "")
        })
        if (this.state.deliveries.length > 0) {
            this.props.onDeliverySelected(this.state.deliveries[0].code)
        }
    }

    onChange(value) {
        this.setState({
            deliveries: this.state.deliveries,
            deliveryNameSelected: this.state.deliveries[value].name
        })
        this.props.onDeliverySelected(this.state.deliveries[value].code)        
    }

    render() {
        return (
            <>

                <Row gutter={50}>
                    <Col>
                        <Space align="center">
                            <label>Seleccione el Reparto</label>
                        </Space>
                    </Col>

                    <Col>
                        <Space align="center">
                            <Select defaultValue="1" style={{ width: 120 }} onChange={this.onChange}>
                                {
                                    this.state.deliveries.map((deliveries, i) => {
                                        return (<Option key={i} value={i}>{deliveries.code}</Option>)
                                    })
                                }
                            </Select>
                        </Space>
                    </Col>
                    <Col>
                        <a>{this.state.deliveryNameSelected}</a>
                    </Col>
                </Row>
            </>
        )
    }
}
