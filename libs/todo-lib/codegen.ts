
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:8080/query",
  // schema: [{ "http://localhost:8080/query": {
  //     customFetch: "graphql-request"
  //   }}], // there's an issue with that
  documents: ["libs/**/**/*.{ts,tsx}", "apps/**/**/*.{ts,tsx}"],
  generates: {
    "libs/todo-lib/src/lib/gql/": {
      preset: "client",
      plugins: []
      //plugins: ['typescript-react-query']
    },
  }
};

export default config;
