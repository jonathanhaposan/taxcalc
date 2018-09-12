# Tax Calculator

### Installation

This app requires [Docker](https://www.docker.com/products/docker-engine) and [Docker Compose](https://docs.docker.com/compose/install/)

Clone the repo.
```sh
$ git clone https://github.com/jonathanhaposan/taxcalc
```

Go to project directory.
```sh
$ cd github.com/jonathanhaposan/taxcalc
```

Run the docker command inside this folder.
```sh
$ docker-compose up
```

The shell will give logs if its success building the container, make sure you see this 2 log.
```sh
postgrestaxcalc    | PostgreSQL init process complete; ready for start up.
postgrestaxcalc    | LOG:  database system was shut down at 2018-09-12 19:29:53 UTC
postgrestaxcalc    | LOG:  MultiXact member wraparound protections are now enabled
postgrestaxcalc    | LOG:  database system is ready to accept connections
postgrestaxcalc    | LOG:  autovacuum launcher started

```

```sh
taxcalc            | 2018/09/12 19:29:56 Succes connect to DB
taxcalc            | 2018/09/12 19:29:56 Server start on :9001
```

After that you can try the app on your web browser by visit this link
```
localhost:9001/bill/list
```

If you made change on code please do re-build the image
```sh
$ docker-compose up --build
```

### Documentation

##### [Link to documentation](https://github.com/jonathanhaposan/taxcalc/tree/master/docs)