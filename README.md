# billsplit

A simple bill splitting app with a Go backend and a Vuejs frontend communicating via a REST api.
Share expenses among friends easily by adding expenses to a bill split and automatically get the balance of each participant.

**Features:**
* View and add new bill splits

![HTTP Log-monitor](images/billsplitlist.png)

* View expenses associated to a bill split

![HTTP Log-monitor](images/billsplit.png)

* Add new expenses

![HTTP Log-monitor](images/newexpense.png)

* Add new participants to a bill split

![HTTP Log-monitor](images/manageparticipants.png)


* View balance of participants

![HTTP Log-monitor](images/balance.png)



## Requirements

The backend is written in Go and uses go modules (> go1.13).
The frontend uses Vuejs (Vuetify, axios, router)

## Backend build instructions

Backend serves at http://localhost:8010/

It uses PostgreSQL as database and you'll need to create a new database:

Database name, user and password can be set as env vars:

`DB_USER` (default `postgres`)

 `DB_PASSWORD` (default `password`)
 
 `DB_NAME` (default `test_bill`)
 
  `DB_HOST` (default `localhost`)
  
   `DB_PORT` (default `5432`)
 
The server address can also be set by an env var:
 `BACK_ADDR ` (default `:8010`)


Launch Postgres: `sudo -u postgres psql`, and then run `creadb DB_NAME`
Exit and then run: `psql -f database/init/setup.sql -d DB_NAME`


to build the backend run: `go build -o bill-split`

to run it:  `./bill-split`

If you want a demo mode with some fake data,  run `./bill-split -demo` instead

## Frontend build instructions

Frontend serves at http://localhost:8080/


The backend server address can be set by an env var in the `frontend/.env` file:

 `VUE_APP_BACK_ADDR ` (default `http://localhost:8010`)

##### Project setup
```
npm install
```

##### Compiles and hot-reloads for development
```
npm run serve
```




