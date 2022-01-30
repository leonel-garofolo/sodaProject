import React from "react";
import Services from "../api/services";

const s = new Services()

export default class ClientCRUDView extends React.Component {
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

    goBefore(){
        alert("goBefore")
    }

    save(){
        alert("Save")
    }

    cancel(){
        alert("Cancel")
    }

    goAfter(){
        alert("goAfter")
    }

    render() { 
        if(this.state.clients.length > 0) {
            return (
                // add label form to focus tab move into it
                <form>
                    <div>
                        <h3>Clientes</h3> 
                        <div>
                            <div><label>Repartos</label></div>
                            <div>
                                <select>
                                    <option id="1">Anselmo</option>
                                    <option id="2">Ariel</option>
                                </select>
                            </div>
                        </div>
                        <div>
                            <div><label>Orden</label></div>
                            <div><input type="text" id="orden" /></div>   
                        </div>
                        <div>
                            <div><label>Precio por Sifon</label></div>
                            <div><input type="text" id="pricePerSoda"/></div>   
                        </div>
                        <div>
                            <div><label>Dirección</label></div>
                            <div><input type="text" id="address"/></div>   
                        </div>
                        <div>
                            <div><label>Numero</label></div>
                            <div><input type="text" id="numberAddress"/></div>   
                        </div>
                        <div>
                            <div><label>Deuda</label></div>
                            <div><input type="text" id="debt"/></div>   
                        </div>

                        <div>
                            <input type="button" value="Anterior" onClick={this.goBefore}/>
                            <input type="button" value="Cancelar" onClick={this.cancel}/>
                            <input type="button" value="Guardar" onClick={this.save}/>                                                    
                            <input type="button" value="Siguiente" onClick={this.goAfter}/>
                        </div>                        
                    </div>
                </form>
            );        
        } else {
            return (
                <div>
                    <h1>Clientes</h1>                        
                </div>
            );        
        }
    }
}