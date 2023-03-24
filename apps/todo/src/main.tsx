import { StrictMode } from 'react';
import * as ReactDOM from 'react-dom/client';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import App from './app/app';

// import { ApolloClient, ApolloProvider, HttpLink, InMemoryCache } from '@apollo/client';

// const client = new ApolloClient({
//   cache: new InMemoryCache(), // cache query results after fetching them
//   link: new HttpLink({
//     uri: 'http://localhost:8080/query',
//   }),
// });


const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 1000,
    },
  },
});

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
    </QueryClientProvider>
    {/*<ApolloProvider client={client}>*/}
    {/*</ApolloProvider>*/}
  </StrictMode>
);
