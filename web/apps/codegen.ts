import type { CodegenConfig } from "@graphql-codegen/cli";

import pkgjson from "./package.json";

const schema = pkgjson.graphql.schema;
const rootGQLDirectory = "src/gql/__gen__/";
const pluginsDirectory = `${rootGQLDirectory}/plugins`;

const config: CodegenConfig = {
  overwrite: true,
  schema,
  documents: "src/**/*.graphql",
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    [rootGQLDirectory]: {
      preset: "client",
    },
    [`${pluginsDirectory}/graphql-request.ts`]: {
      plugins: ["typescript", "typescript-operations", "typescript-graphql-request"],
    },
  },
};

export default config;
