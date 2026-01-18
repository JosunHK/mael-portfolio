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

## Setting Up
create a log file 
`./logs/app.log`

create a .env file and configure your port and db connection string
`./.env`

### MAKEFILE
The project uses a makefile to run the server and the client.
To run the server, use the following command:

`make dev`
