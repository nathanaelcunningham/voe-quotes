## Setup

install golang-migrate https://github.com/golang-migrate/migrate

In the frontend folder `npm install`

### Run Migrations

```run from root
migrate -path migrations/ -database mysql://quotes:quotes@/quotes up
```
