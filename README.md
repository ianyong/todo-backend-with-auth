# To-do Backend

Authentication & authorisation using JSON Web Tokens.

Created for CS3219 Software Engineering Principles and Patterns Own Time Own Target (OTOT) Task C.
Built upon OTOT Task B.

Frontend: https://github.com/ianyong/todo-frontend

## Getting Started

1. Install Go `>= 1.17` by following the instructions [here](https://golang.org/doc/install).
1. Install PostgreSQL `>= 12` by following the instructions [here](https://www.postgresql.org/download/).
1. Make a copy of `.env.development` as `.env.development.local`.
   For development on your local system, it is recommended that you connect to PostgreSQL via Unix-domain sockets so that there is no need for database server credentials, allowing you to leave `DB_USER` and `DB_PASSWORD` empty.
   This can be done by setting the `DB_HOST` to the following locations:
   * Linux: `/var/run/postgresql`
   * macOS: `/tmp`

   Note that if you are connecting via `localhost`, you might need to set `DB_SSLMODE=disable`.
1. Create the database.
   ```sh
   $ make createdb
   ```
1. Migrate the database.
   ```sh
   $ make migratedb
   ```
1. Start the server.
   By default, the backend is accessible at http://localhost:8000/.
   ```sh
   $ make run
   ```

## Linting

1. Install `golangci-lint` by following the instructions [here](https://golangci-lint.run/usage/install/#local-installation).
1. Optionally, you can integrate the linter with your IDE if it is supported by following the instructions [here](https://golangci-lint.run/usage/integrations/).
   Otherwise, you will need to run `make fmt` and `make lintfix` to automatically format and fix any lint violations before you commit any changes, or add a pre-commit Git hook that does it for you automatically.

Note that `gosec` is a supported linter in `golangci-lint`.
As such, there is no need to separately install `gosec` for local development.

## Running Tests

1. Make a copy of `.env.test` as `.env.test.local`.
   The configuration should be similar to `.env.development.local` as described above, but with a different database name.
1. Create the test database.
   ```sh
   $ make createtestdb
   ```
1. Migrate the test database.
   ```sh
   $ make migratetestdb
   ```
1. Run the tests.
   ```sh
   $ make test
   ```

## Deployment

Deployment makes use of the [Serverless Framework](https://www.serverless.com/).
To deploy to AWS Lambda manually:

1. Install Node `>= 14` by following the instructions [here](https://nodejs.org/en/download/).
1. Install the Serverless Framework:
   ```sh
   $ npm install -g serverless
   ```
1. Build the server for use with AWS Lambda.
   Note that the normal server has to be wrapped with a translation layer so that the server is able to understand API requests and responses that go through AWS Lambda.
   ``` sh
   $ make buildlambda
   ```
1. Export the environment variables (see any of the `.env.*` files) for the database connection that will be used by the AWS Lambda function.
1. Migrate the database that will be used.
   Make sure that the correct environment variables are loaded from the step before.
   ```sh
   $ GO_ENV=production make migratedb
   ```
1. Deploy the server.
   ```sh
   $ serverless deploy
   ```

## Testing Authentication & Authorisation

1. Insert two users into the database, one with role set to `user`, and the other with role set to `admin`.
1. Two endpoints have been created for testing purposes:
   * `/api/v1/helloworld/user`: requires the user to be logged in
   * `/api/v1/helloworld/admin`: requires the user to be logged in and to have the `admin` role

   Using the access token generated when logging in, make API calls to the two endpoints with the following states:
   * Not logged in (i.e., no access token)
   * Logged in as a user with role `user`
   * Logged in as a user with role `admin`
