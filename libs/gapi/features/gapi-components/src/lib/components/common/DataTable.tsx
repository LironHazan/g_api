import React, { useState } from 'react'
import { DataGrid } from '@mui/x-data-grid';


const DataTable = ({
                     rows,
                     loading,
                     sx,
                     onRowClick
                   }: any) => {
  const [pageSize, setPageSize] = useState(10);

  const columns = [
    {
      field: 'name',
      headerName: 'Name',
      flex: 1,
      cellClassName: 'name-column--cell'
    },    {
      field: 'url',
      headerName: 'Link',
      flex: 1,
      cellClassName: 'link-column--cell'
    },
    {
      field: 'description',
      headerName: 'Description',
      flex: 1
    },
    {
      field: 'open_issues',
      headerName: 'Issues',
      flex: 1,
      width: 150
    },
    {
      field: "visibility",
      headerName: "Visibility",
      flex: 1,
    },
  ];

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
      onRowClick={onRowClick}
    />
  );
};

export default DataTable
