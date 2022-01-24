import React from 'react';
import './App.css';
import Main from './layout/Main';
import Menu from './layout/Menu';
import Services from './api/services';

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
        <header className="App-header">
          <Menu onClick={this.onMenuSelected} />
          <div className="App-content">
            <Main content={this.state.content}/>
          </div>
        </header>
      </div>
    );
  }
 
}

export default App;
