enum Env {
  PROD= 'prod',
  DEV = 'dev'
}

const ENV = Env.DEV

interface EnvVar {
  emoticGraphQLBaseURL: string
}

const envVars: {
  [env: string]:  EnvVar
} = {
  'prod': {
    emoticGraphQLBaseURL: 'https://api-emotic.short-d.com/graphql'
  },
  'dev': {
    emoticGraphQLBaseURL: 'http://localhost:8080/graphql'
  }
};

export class Environment {
  getEnv(): EnvVar {
    return envVars[ENV];
  }
}
