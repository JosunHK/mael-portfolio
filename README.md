# Getting started 

## Installation

### Tailwind CSS
`https://tailwindcss.com/blog/standalone-cli`
use -> https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.19/tailwindcss-linux-x64

### Go dependencies
install air base on your environment
`https://github.com/air-verse/air`

install sqlc base on your environment
`https://docs.sqlc.dev/en/stable/overview/install.html`

install templ base on your environment
`go install github.com/a-h/templ/cmd/templ@latest`

install goose to manage database migration
`go install github.com/pressly/goose/v3/cmd/goose@latest`

## Setting Up
create a log file 
`./logs/app.log`

create a .env file and configure your port and db connection string
`./.env`

### MAKEFILE
The project uses a makefile to run the server and the client.
To run the server, use the following command:

`make dev`

`make tailwind-watch` to watch & update css changes

### Database Schema 
To update database schema, we do this through goose

use `goose create <migration_name> sql` to creat new migration

see [documentations](https://github.com/pressly/goose) for how to do this.

`goose up` to apply changes

`goose down` to revert changes

`sqlc generate` to auto generate schema as structs
