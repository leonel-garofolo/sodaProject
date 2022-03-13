import React from "react";
import Services from '../api/services';
import { Layout, Table } from 'antd';
const { Content } = Layout;

const s = new Services()
const columns = [
    {
      title: 'Id',
      dataIndex: 'id',
      key: 'id',      
    },
    {
      title: 'Nombre',
      dataIndex: 'name',
      key: 'name',      
    }
  ];
export default class DeliveryView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            deliveries: [],
        }
    }
    async componentWillMount() {
        await s.getDeliveries()   
        this.setState({                
            deliveries: s.getResponse().data
        })
    }

    render() {    
        if(this.state.deliveries.length > 0) {
            return (
                <Layout>
                    <Content>
                        <Table pagination={{position:['none']}} dataSource={this.state.deliveries} columns={columns} />
                    </Content>                    
                </Layout>                
            );        
        } else {
            return (
                <div>
                    <h3>Cargando...</h3>
                </div>
            );        
        }       
    }
}