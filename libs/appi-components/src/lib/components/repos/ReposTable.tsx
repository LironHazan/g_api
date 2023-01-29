import Box from '@mui/material/Box';
import { useTheme } from '@mui/material';
import { tokens } from '../../theme';
import DataTable from '../common/DataTable';
import { GridEventListener } from '@mui/x-data-grid';
import { ApolloError } from '@apollo/client/errors';
import { useRepos } from '../../hooks/useRepos';

const userTableStyles = {
  height: '650px',
};

interface ReposType {
  loading: boolean,
  error: ApolloError | undefined,
  repos: { id: string, name: string, url: string, description: string }[] | undefined
}

function ReposTable({ repos }: ReposType) {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);

  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
    console.log(`-- "${params.row.name}" clicked`);
    // trigget drawer
  };

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
    <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
      { !repos ? <div> Search for a github user </div> :
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
              // loading={!repos?.length}
              sx={userTableStyles}
              onRowClick={handleRowClick}
            />
          </Box>
        </Box>
      }
    </Box>
  );
}

export function UserReposTable({ username }: { username: string }) {
  const { loading, error, repos } = useRepos(username)
  return (
    <ReposTable error={error} loading={loading} repos={repos}/>
  );
}
