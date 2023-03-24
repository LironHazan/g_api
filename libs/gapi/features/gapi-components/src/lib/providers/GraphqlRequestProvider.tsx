import { createContext, ReactNode, useContext } from 'react';
import { GraphQLClient } from 'graphql-request';

const GraphqlRequestContext = createContext({} as GraphQLClient);

export const useGraphqlRequest = () => useContext<GraphQLClient>(GraphqlRequestContext);

interface GraphqlRequestProviderProps { client: GraphQLClient; children: ReactNode }
export function GraphqlRequestProvider({ client, children }: GraphqlRequestProviderProps) {
  return (
    <GraphqlRequestContext.Provider value={client}>
      {children}
    </GraphqlRequestContext.Provider>
  );
}
