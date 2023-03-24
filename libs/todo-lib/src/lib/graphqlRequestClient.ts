import { GraphQLClient } from 'graphql-request';

const requestHeaders = {
  authorization: 'Bearer MY_TOKEN', // in case I'll add auth to API will pass an env var
};

const graphqlRequestClient = new GraphQLClient('http://localhost:8080/query', {
  headers: requestHeaders,
});

export default graphqlRequestClient;
