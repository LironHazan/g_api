import { UserReposTable } from '../components/repos/ReposTable';
import { SetStateAction, useState } from 'react';
import { SearchField } from '../components/SearchField';

export function ReposPage() {
  const [uname, setUserName] = useState('');

  const handleChange = (uname: SetStateAction<string>) => {
    setUserName(uname);
  }
  return <>
    <SearchField liftUserName={handleChange}/>
    <UserReposTable username={uname}/>
  </>
}
