import { ApolloError } from '@apollo/client/errors';
import { useTheme } from '@mui/material';
import { tokens } from '../../theme';
import { GridEventListener } from '@mui/x-data-grid';
import Box from '@mui/material/Box';
import DataTable from '../common/DataTable';

const userTableStyles = {
  height: '650px',
};

interface ReposType {
  loading: boolean,
  error: ApolloError | undefined,
  repos: { id: string, name: string, url: string, description: string }[] | undefined
}

export function ReposTable({ repos }: ReposType) {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);

  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
    console.log(`-- "${params.row.name}" clicked`);
    // trigget drawer
  };

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
