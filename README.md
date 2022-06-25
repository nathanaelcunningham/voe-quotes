## Setup

install golang-migrate https://github.com/golang-migrate/migrate

## Run Migrations

```run from root
migrate -path migrations/ -database mysql://quotes:quotes@/quotes up
```
