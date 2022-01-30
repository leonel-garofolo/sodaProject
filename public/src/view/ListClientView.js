import React from "react";
import Services from "../api/services";
import DataTable from "./table/DataTable";

const s = new Services()

export default class ListClientView extends React.Component {
    constructor(props) {
        super(props);   
        this.state = {
            clients: [],
        }        
    }

    async componentWillMount() {
        await s.getDeliveryClient()        
        this.setState({                
            clients: s.getData()
        })
    }

    save(){

    }

    cancel(){

    }

    print(){

    }

    render() { 
        if(this.state.clients.length > 0) {
            return (
                <div>
                    <h3>Listado de Clientes</h3>
                    <div>
                        <div><label>Repartos</label></div>
                        <div>
                            <select>
                                <option id="1">Anselmo</option>
                                <option id="2">Ariel</option>
                            </select>
                        </div>
                    </div>    
                    <DataTable data={this.state.clients}/>
                    <input type="button" value="Cancelar" onClick={this.cancel}/>
                    <input type="button" value="Guardar" onClick={this.save}/>                                            
                    <input type="button" value="Imprimir" onClick={this.print}/>
                </div>
            );        
        } else {
            return (
                <div>
                    <h3>Listado de Clientes</h3>                        
                </div>
            );        
        }
    }
}