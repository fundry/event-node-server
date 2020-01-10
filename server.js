require('dotenv').config();
import { ApolloServer } from 'apollo-server';
import { Prisma } from 'prisma-binding';
import { schema } from './src/schema';

// switching environments .... NODE_ENV is undefined
const env = process.argv[2];
let env_variable;

{
  env === 'dev'
    ? (env_variable = process.env.DEV_PRISMA)
    : (env_variable = process.env.PROD_PRISMA);
}
// ===================================================>

const server = new ApolloServer({
  introspection: true,
  schema,
  context: (req) => ({
    ...req,
    db: new Prisma({
      typeDefs: './prisma/prisma.graphql',
      endpoint: `${env_variable}` ,
      debug: true,
    }),
    prisma: new Prisma({
      typeDefs: './prisma/prisma.graphql',
      endpoint: `${env_variable}` , // process.env.env_variable ... using env_variable cause prisma to disconnect
      debug: true,
    }),
  }),
});

server
  .listen()
  .then(({ url }) => {
    console.log(`💩  is running at ${url}`);
    console.log(env_variable);
  })
  .catch((e) => console.log(e));
