import { Repos } from '../components/repos/Repos';
import { useRepos } from '../hooks/useRepos';

export function ReposPage() {
  const { loading, error, repos } = useRepos("lironhazan")

  return (
    <Repos error={error} loading={loading} repos={repos}/>
  );
}
