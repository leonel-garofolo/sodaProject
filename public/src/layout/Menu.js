import React from "react";

export default class Menu extends React.Component {
    constructor(props) {
        super(props);
        this.onClicked = this.onClicked.bind(this);
      }

    onClicked(event){     
        this.props.onClick(event.target.getAttribute('href').replace('#', ''))
    }
    render() {
        return (
            <div>            
                <ul>                
                    <li><a href="#clients" onClick={this.onClicked}>Clientes</a></li>
                    <li><a href="#list" onClick={this.onClicked}>Listados de Repartos</a></li>
                    <li><a href="#delivery" onClick={this.onClicked}>Repartidores</a></li>                                        
                </ul>
            </div>
        );   
    }
}



