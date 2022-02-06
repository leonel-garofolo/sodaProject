import React from 'react'
import {AgGridColumn, AgGridReact} from 'ag-grid-react';

import 'ag-grid-community/dist/styles/ag-grid.css';
import 'ag-grid-community/dist/styles/ag-theme-alpine.css';

function DataTable(jsonData) {
  const data = React.useMemo(() => jsonData.data) 
  return (
    <div className="ag-theme-alpine" style={{height: 400}}>
           <AgGridReact               
               rowData={data}>
               <AgGridColumn field="order" headerName="Orden" sortable={ true } filter={true} floatingFilter={true} maxWidth="120" />
               <AgGridColumn field="address" headerName="Dirección" sortable={ true } filter={true} floatingFilter={true} maxWidth="400"/>
               <AgGridColumn field="numAddress"  headerName="Numero" sortable={ true } filter={true} floatingFilter={true} maxWidth="130" />
               <AgGridColumn field="pricePerSoda" headerName="Precio x Soda" />
               <AgGridColumn field="pricePerBox" headerName="Precio x Cajon" />
               <AgGridColumn field="debt" headerName="Deuda" maxWidth="120" />
           </AgGridReact>
       </div>
  )
}

export default DataTable
