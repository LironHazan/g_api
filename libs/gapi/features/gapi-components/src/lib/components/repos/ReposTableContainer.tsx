import { useRepos } from '../../hooks/useRepos';
import { ReposTable } from './ReposTable';


export function UserReposTableContainer({ username }: { username: string }) {
  const { loading, error, repos } = useRepos(username)
  return (
    <ReposTable error={error} loading={loading} repos={repos}/>
  );
}
