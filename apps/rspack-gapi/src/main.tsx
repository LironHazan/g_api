import { StrictMode } from 'react';
import * as ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
//import { ApolloClient, ApolloProvider, InMemoryCache, HttpLink } from '@apollo/client';
import App from './app/App';

//const GITHUB_BASE_URL = 'https://api.github.com/graphql';
//const GITHUB_TOKEN = process.env['NX_GITHUB_TOKEN'];

// https://www.apollographql.com/docs/resources/graphql-glossary/
// const client = new ApolloClient({
//   cache: new InMemoryCache(), // cache query results after fetching them
//   link: new HttpLink({
//     uri: GITHUB_BASE_URL,
//     headers: {
//       authorization: `Bearer ${GITHUB_TOKEN}` //todo: make an env variable
//     },
//   }),
// });

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>,
  </StrictMode>
);
