import React from "react";
import ReactDOM from "react-dom";
import ClientCRUDView from "../view/ClientCRUDView";
import DeliveryView from "../view/DeliveryView";
import ListClientView from "../view/ListClientView";

export default class Main extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {    
        switch (this.props.content) {
            case 'clients':
                return (
                    <ClientCRUDView />
                );
            case 'list':
                return (
                    <ListClientView />
                );
            case 'delivery':
                return (
                    <DeliveryView />
                );
        
            default:
                return (
                    <div>
                        <h1>En construcción</h1>                    
                    </div>
                );
        }    
      }
}