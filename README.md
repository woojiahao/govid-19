# govid-19

Go API for retrieving Covid-19 statistics

## Changelog

- Data structure (see below)
- Time series data is now stored globally at runtime and updated daily when fetching a new set of results

## Endpoints

- `GET /latest` - retrieve the latest statistics
- `GET /countries` - retrieve all countries
- `GET /ping` - ping to see if the server is running
- `GET /all` - retrieve statistics of all countries

### GET /ping

Test if the server is running.

### GET /all

Returns all statistics. 

#### Query parameters

- `country` -  case sensitive search for the country to return
- `state` - case sensitive search for the state to return
- `first` - get the first *n* number of records
- `last` - get the last *n* number of records
- `sort-data` - sort the results by the total value of each category

### GET /countries

Returns all available countries and states.

## Deployment

### Local

Ensure that you have Docker installed

Run the Docker compose file to load all necessary information.

```bash
$ docker-compose up
```

## TODO

- [X] Auto-pull the repository every day
- [X] Compute the overall changes of the data since the data is accumulative now
- [ ] Deployment guide
- [X] Endpoint for country information

## Lessons

- CORS must be configured before any API routes are created otherwise the client will encounter CORS issues
