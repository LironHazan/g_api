import Box from '@mui/material/Box';
import { useTheme } from '@mui/material';
import { useEffect, useState } from 'react';
import { tokens } from '../theme';
import DataTable from '../components/common/DataTable';
import { GridEventListener } from '@mui/x-data-grid';

const userTableStyles = {
  height: '650px',
};

export function Tasks(){
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);

  const [users, setUsers] = useState([]);

  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
    console.log(`-- "${params.row.name}" clicked`);
    // trigget drawer
  };


    useEffect(() => {
    fetch('https://api.github.com/users/lironhazan/repos')
      .then((response) => response.json())
      .then((json) => {
        json = json.map((item: any) =>  {
          item.owner = item.owner.login;
          return item;
        } )
        return setUsers(json)
      })
      .catch(() => void 0)
  }, []);


  const columns = [
    { field: 'id', headerName: 'ID' },
    {
      field: 'name',
      headerName: 'Name',
      flex: 1,
      cellClassName: 'name-column--cell'
    },    {
      field: 'html_url',
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
    <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
      <Box m="20px">
        <Box
          m="40px 0 0 0"
          height="75vh"
          sx={{
            "& .MuiDataGrid-columnHeaders": {
              color: colors.blueAccent[300],
            },
            "& .MuiCheckbox-root": {
              color: `${colors.blueAccent[200]} !important`,
            },
          }}
        >
          <DataTable
            rows={users}
            columns={columns}
            loading={!users.length}
            sx={userTableStyles}
            onRowClick={handleRowClick}
          />
        </Box>
      </Box>
    </Box>
  );
}
