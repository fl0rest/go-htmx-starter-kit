# Go-HTMX Starter Kit

Designed to be a starting point for making apps with Go, HTMX and Bootstrap


## Installation

You don't need to install anything, just download the zip and start working, you can run the below to grab the latest release.

```bash
wget https://github.com/fl0rest/go-htmx-starter-kit/releases/latest/download/starter.zip
```

After downloading, just unzip it and start working.
## Running the app

To run this project run

```bash
  go run cmd/main.go
```


## Documentation

Use `internal/` for internal handlers, validations, models and such

Use `static/` for static files - HTML, CSS, JS

Use `assets/` for images and icons

You can define environment variables in .env or by adding them into your environment.

A current list of accepted environment variables:


| Key | Default Value | Description |
| --- | ------------- | ----------- |
| PORT | `8000` | Port where the application will be served |
| ENV | `DEVELOPMENT` | The environment in which the application is running |
| LOG_TO_SCREEN | `true` | Output log statements to screen |
| LOG_FILE | `./log/app.log` | Path for the Info, Sys and Warn logfile |
| ERROR_LOG_FILE | `./log/error.log` | Path for the Error logfile |
| DEBUG | `false` |  |


## Roadmap

- Databases

- Better starting pages

- Utilizing Go templating

