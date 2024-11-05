# Go Journal

Basic Journaling application based on Go, Postgres, Templ, HTMX and Tailwind.

## Run

Start Postgres via `docker compose`:

`docker compose -f docker-compose.yml up -d `

Then, run the app by executing the start script `./start.sh`.

No proper env config has been set up yet and this has only been tested locally.

To regenerate template, execute `templ generate`