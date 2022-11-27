import React, { useState } from 'react'
import { DataGrid } from '@mui/x-data-grid';

const DataTable = ({
                     rows,
                     columns,
                     loading,
                     sx
                   }: any) => {
  const [pageSize, setPageSize] = useState(10);

  // const handleRowClick: GridEventListener<'rowClick'> = (params) => {
  //   setMessage(`Movie "${params.row.title}" clicked`);
  // };

  return (
    <DataGrid
      rows={rows}
      columns={columns}
      loading={loading}
      sx={sx}
      checkboxSelection
      pagination
      pageSize={pageSize}
      onPageSizeChange={(newPageSize) => setPageSize(newPageSize)}
      rowsPerPageOptions={[10, 15, 25]}
      // onRowClick={handleRowClick} {...data}
    />
  );
};

export default DataTable
