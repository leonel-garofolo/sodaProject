import React from 'react';
import './App.css';
import Main from './layout/Main';
import Services from './api/services';
import MenuDashboard from './layout/MenuDashboard';
import { Layout } from 'antd';

const { Content } = Layout;

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      content: "list",
    };
    this.onMenuSelected= this.onMenuSelected.bind(this) 
  }

  onMenuSelected(sActionClicked){        
    var s = new Services()
    s.getDeliveries()
    this.setState({content: sActionClicked})
  }

  render() {
    return (
      <div className="App">
        <Layout className="layout">
          <MenuDashboard onClick={this.onMenuSelected} />
          <Content>
              <Main content={this.state.content}/>
          </Content>
        </Layout>        
      </div>
    );
  }
 
}

export default App;
