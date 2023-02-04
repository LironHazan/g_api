import TextField from '@mui/material/TextField';
import { SetStateAction } from 'react';

type LiftUserNameFn = (uname: SetStateAction<string>) => void;
export const SearchField = (props: { liftUserName: LiftUserNameFn}) => {
  const { liftUserName } = props;
  const handleChange = (event: { target: { value: string; }; }) => {
    liftUserName(event.target.value);
  }

  return (
    <TextField placeholder="Search..." size="small" onChange={handleChange} defaultValue="" />
);
}

