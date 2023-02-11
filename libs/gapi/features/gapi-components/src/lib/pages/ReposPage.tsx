import { UserReposTableContainer } from '../components/repos/ReposTableContainer';
import { SetStateAction, useState } from 'react';
import { SearchField } from '../components/SearchField';

export function ReposPage() {
  const [uname, setUserName] = useState('');

  const handleChange = (uname: SetStateAction<string>) => {
    setUserName(uname);
  }
  return <>
    <SearchField liftUserName={handleChange}/>
    <UserReposTableContainer username={uname}/>
  </>
}
