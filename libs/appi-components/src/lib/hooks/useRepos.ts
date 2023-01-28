import { gql, useQuery } from '@apollo/client';
import { useMemo } from 'react';

const GET_REPOS =  gql`
query FetchRepos($username: String!)  {
  user(login: $username) {
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

export const useRepos = (username: string) => {
  const { loading, error, data } = useQuery(GET_REPOS,{
    variables: {
      username: username
    }
  });

  const repos = useMemo(() => data?.user.repositories.nodes
    .map((n: { name: string, url: string, description: string }) => ({...n, id: n.name})
  ), [data]) || []

  return { loading, error, repos }
}
