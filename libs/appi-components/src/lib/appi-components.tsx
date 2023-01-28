import styled from '@emotion/styled';

import { useQuery, gql } from '@apollo/client';

const GET_LOCATIONS = gql`
query {
  repository(owner:"octocat", name:"Hello-World") {
    issues(last:20, states:CLOSED) {
      edges {
        node {
          title
          url
          labels(first:5) {
            edges {
              node {
                name
              }
            }
          }
        }
      }
    }
  }
}
`;

function DisplayLocations() {
  const { loading, error, data } = useQuery(GET_LOCATIONS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return data.locations.map(({ id, name, description, photo }: any) => (
    <div key={id}>
      <h3>{name}</h3>
      <img width="400" height="250" alt="location-reference" src={`${photo}`} />
      <br />
      <b>About this location:</b>
      <p>{description}</p>
      <br />
    </div>
  ));
}

/* eslint-disable-next-line */
export interface AppiComponentsProps {}

const StyledAppiComponents = styled.div`
  color: pink;
`;

export function AppiComponents(props: AppiComponentsProps) {
  return (
    <StyledAppiComponents>
      <h2>My first Apollo app 🚀</h2>
      <br/>
      <DisplayLocations />
    </StyledAppiComponents>
  );
}

export default AppiComponents;
