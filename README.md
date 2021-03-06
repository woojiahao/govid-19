# govid-19

Go API for retrieving Covid-19 statistics

## Changelog

- Data structure (see below)
- Time series data is now stored globally at runtime and updated daily when fetching a new set of results
- Removed the `/all` endpoint and replaced it with separate endpoints to handle different levels of data
    analysis 

## Endpoints

- `GET /latest` - retrieve the latest statistics
- `GET /countries` - retrieve all countries
- `GET /ping` - ping to see if the server is running
- `GET /all` - retrieve statistics of all countries

### GET /ping

Test if the server is running.

### GET /stats?show=country,state,date

### GET /all

Returns all statistics. 

#### Query parameters

- `country` -  case sensitive search for the country to return
- `state` - case sensitive search for the state to return
- `first` - get the first *n* number of records
- `last` - get the last *n* number of records
- `sort-total` - sort the results by the total value of each category
- `sort-records` - sort the results by the total value of each category

#### Response structure

```json
{
  "<country>": {
    "<state>": {
      "long": 0.0,
      "lat": 0.0,
      "confirmed": {
        "total": 0,
        "data": {
          "<date - dd-MM-yyyy>": 0
        }   
      },
      "recovered": {
        "total": 0,
        "data": {
          "<date - dd-MM-yyyy>": 0
        }   
      },
      "deaths": {
        "total": 0,
        "data": {
          "<date - dd-MM-yyyy>": 0
        }   
      }
    } 
  }
}
```

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
- [X] Fully migrate to use database to query data
- [ ] Use middleware to handle errors thrown
- [ ] Use bubble chart for showing most prominent countries
- [ ] Endpoint for grouping countries that have states

## Lessons

- CORS must be configured before any API routes are created otherwise the client will encounter CORS issues
- GORM is limited in its functionality with tags; for foreign keys, the association must be declared during table 
    creation and not during struct declaration
