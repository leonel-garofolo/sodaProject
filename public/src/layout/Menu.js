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
                    <li><a href="#list" onClick={this.onClicked}>Listado</a></li>
                    <li><a href="#delivery" onClick={this.onClicked}>Repartidores</a></li>
                    <li><a href="#print" onClick={this.onClicked}>Imprimir</a></li>                    
                    <li><a href="#exit" onClick={this.onClicked}>Salir</a></li>
                </ul>
            </div>
        );   
    }
}



