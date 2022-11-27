import Box from '@mui/material/Box';
import { useTheme } from '@mui/material';
import { useEffect, useState } from 'react';
import { tokens } from '../theme';
import DataTable from '../components/common/DataTable';

const columns = [
  { field: 'id', headerName: 'User ID', width: 150 },
  { field: 'name', headerName: 'Name', width: 150 },
  { field: 'username', headerName: 'Username', width: 150 },
  { field: 'email', headerName: 'E-mail', width: 200 },
];

const userTableStyles = {
  height: '650px',
};

export default function Tasks() {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);

  const [users, setUsers] = useState([]);

  async function repoDataURL() {
    await fetch("https://api.github.com/users/LiveDetermined/repos")
      .then((res) => res.json())
      .then(
        (result) => {
          console.log(36, result);
          const list = result.map((item: any) => (
            <div className="text-center">
              <a target="_blank" href={item.svn_url}>
                {item.name}
              </a>
            </div>
          ));
        },
        (error) => {
          console.log(error);
        }
      );
  }

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
      flex: 1
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
          />
        </Box>
      </Box>
    </Box>
  );
}
