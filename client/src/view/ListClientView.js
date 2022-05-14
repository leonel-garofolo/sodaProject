import React from "react";
import Button from "react-bootstrap/esm/Button";
import Services from "../api/services";
import DeliveryList from "./components/DeliveryList";

import { Row, Col, Divider, Table } from 'antd';


const s = new Services()


export default class ListClientView extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      clients: [],
      filteredInfo: null,
      sortedInfo: null,
    }
    this.onDeliverySelected = this.onDeliverySelected.bind(this)
    this.handleChange = this.handleChange.bind(this)
  }

  async print() {
    await s.getReport("10")
    console.log(s.data)
  }

  async onDeliverySelected(deliverySelected) {
    console.log(deliverySelected)
    await s.getDeliveryClient(deliverySelected.code)
    this.setState({
      clients: s.getResponse().data
    })
  }

  handleChange = (pagination, filters, sorter) => {
    this.setState({
      filteredInfo: filters,
      sortedInfo: sorter,
    });
  };

  render() {
    let { sortedInfo, filteredInfo } = this.state;
    sortedInfo = sortedInfo || {};
    filteredInfo = filteredInfo || {};

    const columns = [
      {
        title: 'Orden',
        dataIndex: 'order',
        key: 'order',
        width: '5%',
      },
      {
        title: 'Dirección',
        dataIndex: 'address',
        key: 'address',
        width: '30%',
        filteredValue: filteredInfo.address || null,
        onFilter: (value, record) => record.address.includes(value),
        sorter: (a, b) => a.address.length - b.address.length,
        sortOrder: sortedInfo.columnKey === 'address' && sortedInfo.order,
        ellipsis: true,
      },
      {
        title: 'Numero',
        dataIndex: 'num_address',
        key: 'num_address',
        sorter: {
          compare: (a, b) => a.num_address - b.num_address,
        },
      },
      {
        title: 'recio x Soda',
        dataIndex: 'price_per_soda',
        key: 'price_per_soda',
      },
      {
        title: 'Precio x Cajon',
        dataIndex: 'price_per_box',
        key: 'price_per_box',
      },
      {
        title: 'Deuda',
        dataIndex: 'debt',
        key: 'debt',
      }
    ];

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
        <Table dataSource={this.state.clients} columns={columns} onChange={this.handleChange} />
      </>
    )
  }
}
