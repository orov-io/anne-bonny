# Anne Bonny

A template to bootstraping microservices with go.

Also helps to [fight global warming](https://en.wikipedia.org/wiki/Flying_Spaghetti_Monster#Pirates_and_global_warming) acting as [Anne Bonny](https://en.wikipedia.org/wiki/Anne_Bonny).

## Testing

Please, provide env vars for testing in a file called testing.env at repo root folder level. See the [Workspace declaration](./AnneBonny.code-workspace) to infer the testing usage.

## TODOS

- [] Tutorial
- [x] Add Database to a single service
- [] Use squirrel to build SQL statements
- [] Use echo+Validator to validate request data
- [] Add hot reloading to docker-compose
- [] Healthcheck in dockerfile
- [] Add a front client (use flutter)
- [] Swagger auto generation
- [] Add a public golang sdk api client for the exposed API.
- [] Add a private golang sdk api client for both the exposed and the internal services API.
- [] Deployer
- [] IaC
- [] KrakenD : It will coexist with caddy? Caddy will do the krakenD tas for no local overcharging?
- [] Jaeaguer
- [] Kubernetes
- [] Local CI/CD tool with Jenkins or GitLabCI
- [] Configure caddy and localhost to generate local certificates for *.annebonny.dev and serve they as in [this tutorial](https://medium.com/@devahmedshendy/traditional-setup-run-local-development-over-https-using-caddy-964884e75232)
- [x] Add migrations with [goose]("https://github.com/pressly/goose")
- [] Add ghost as a blog for Anne Bonny
- [] Check golang generics in order to know if you can add the factory injection in maryread.
- [] Concurrency example
- [] Cronjobs. Create a new user service and allow to create temp users (6 hours, 3 days, 1 week).
- [] Use release-it to app version management.
- [] GRPC. Let maryread start in grcp mode?
- [] Websockets example.
- [] Integration tests with the client.
- [] Basic SQL test data seeder per service
- [] Use the client to seed the application with test data
- [] Add Benchmark tests.
- [] Add to the client mock responses. At now, to test the video handler, you must up and expose the storage service
- [] Find a way to load non tracked .env files to the makefile.

## Tutorial

You must have both a "testing.env" file and a "override.env" file in the root of the project in order to invoque the make commands.

- General Env Vars
- Repos Env Vars
- Running locally
- Running test
-- Running all test in vs
-- Debugging test in vs code
-- Generate a test report and review it in a browse.
- Deploying
- Tools
- direnv to load .envrc (simplify your life testing migrations, for example).
