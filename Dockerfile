FROM golang:1.18.3 as development

# set working directory
WORKDIR /app

# Cache and install deps
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# install reload tool
RUN  go install github.com/cosmtrek/air@latest

#install migrate tool
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 3000

CMD air
