import { Layout, Menu } from 'antd';
import React from "react";
const { Header} = Layout;

export default class MenuDashboard extends React.Component {
    constructor(props) {
        super(props);
        this.onClicked = this.onClicked.bind(this);
    }

    onClicked(event) {
        console.log(event)
        this.props.onClick(event.target.getAttribute('href').replace('#', ''))
    }
    render() {
        return (
            <Header>
                <div className="logo" />
                <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['0']}
                    onSelect={(item) => this.props.onClick(item.key)}>
                    <Menu.Item key="clients" >Clientes</Menu.Item>
                    <Menu.Item key="list">Listados de Repartos</Menu.Item>
                    <Menu.Item key="delivery">Repartidores</Menu.Item>
                </Menu>
            </Header>
        );
    }
}



