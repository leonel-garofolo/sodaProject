import React from "react";
import Container from "react-bootstrap/esm/Container";
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar'

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
            <Navbar bg="light"
                onSelect={(selectedKey) => this.props.onClick(selectedKey)}
                >
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav.Link eventKey="clients">Clientes</Nav.Link>
                    <Nav.Link eventKey="list">Listados de Repartos</Nav.Link>
                    <Nav.Link eventKey="delivery">Repartidores</Nav.Link>
                </Navbar.Collapse>
          </Navbar>
        );   
    }
}



