import { Routes, Route } from 'react-router-dom';
import { AppiComponents, Tasks, AppMainLayout } from '@g_api/appi-components';
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';

const GITHUB_BASE_URL = 'https://api.github.com/graphql';

const client = new ApolloClient({
  uri: GITHUB_BASE_URL,
  cache: new InMemoryCache(),
});

export function App() {
  return (
    <ApolloProvider client={client}>
    <Routes>
        <Route path="/" element={<AppMainLayout />}>
          <Route path="tasks" element={<Tasks />} />
          <Route path="appi" element={<AppiComponents />} />
        </Route>
    </Routes>
    </ApolloProvider>
  );
}

export default App;
