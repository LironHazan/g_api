import Box from '@mui/material/Box';
import { useTheme } from '@mui/material';
import { useMemo } from 'react';
import { tokens } from '../theme';
import DataTable from '../components/common/DataTable';
import { GridEventListener } from '@mui/x-data-grid';
import { useQuery, gql } from '@apollo/client';

const userTableStyles = {
  height: '650px',
};

const GET_REPOS = gql`
query {
  user(login: "lironhazan") {
    repositories(first: 100) {
      pageInfo {
        endCursor
        hasNextPage
      }
      nodes {
        name
        url
        description
      }
    }
  }
}
`;

export function Repos(){
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const { loading, error, data } = useQuery(GET_REPOS);

  const repos = useMemo(() => data?.user.repositories.nodes.map((n: any) => {
    return {...n, id: n.name};
  } ), [data]) || []

  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
    console.log(`-- "${params.row.name}" clicked`);
    // trigget drawer
  };

  const columns = [
    // { field: 'id', headerName: 'ID' },
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
            rows={repos}
            columns={columns}
            loading={(!repos as any).length}
            sx={userTableStyles}
            onRowClick={handleRowClick}
          />
        </Box>
      </Box>
    </Box>
  );
}
