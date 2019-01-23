# Full-stack Coding

In this exercise, you will be working on a very small project which
shows a list of companies and their mission statements. The goal will be
to alter the code so that it implements a `/companies` endpoint to
return that data from a companies table in MySQL.

## Some resources you may need

- [A Tour of Go](https://tour.golang.org/welcome/1) and
  [Effective Go](https://golang.org/doc/effective_go.html)
- [The Go standard library](https://golang.org/pkg/) packages
- [Pop](https://github.com/gobuffalo/pop), the ORM we use for working
  with the database.
- [A guide to Vue.js](https://vuejs.org/v2/guide/), the front-end
  framework we use.

## The first steps

Look around the repository. There's a Go program, which is our API
server, in `/cmd/api/main.go`. You can find the SQL that describes our
company data in `/data.sql` (and you can also find the program which
built that set in `/cmd/gendata/main.go`!). You can also find packages
which support the API server in `/pkg`.

Now look in the `/frontend` directory; that is the home of our Vue.js
framework code. The HTML we serve is defined initially by
`/frontend/src/index.html`, and we include the root component which is
defined in `/frontend/src/App.vue`. The JavaScript which we use to
create and mount the root component is in `/frontend/src/main.js`.

## Setting up the environment

You must have Docker installed to work with our full stack environment,
which is defined in `docker-compose.yml`; it's composed of a MySQL
service, an API service built from the `/cmd/api`, and a frontend
service built from `/frontend`.

You can use the `make` command to run the environment; you can type
`make up`, or just `make`, to bring the services up, and `make down` to
bring them back down again. When you make changes to the front or
backend files, run `make rebuild` to make them take effect in the docker
services.

You can connect to MySQL from the terminal application; type `mysql -h
127.0.0.1 -u root` at the prompt. You can import the data set in
`/data.sql` by running the following from the root of the
second-interview directory:

```
mysql -h 127.0.0.1 -u root < data.sql
```

## Expectations

In the ideal state, you will have the `/companies` endpoint working with
the API service, and be able to render that information from the `/`
route in Vue.js. But, most importantly, our expectation is that you
learn from the experience, with Go, Vue, and with the structure we use
to work on our projects.
