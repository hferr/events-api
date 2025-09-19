## Events API

Simple events API

## Dependencies

- postgres as the database
- [migrate CLI](https://github.com/golang-migrate/migrate) to run migrations

## Running the application on Docker

The app can be started by running the following commands, which will pull the database, build and start the app:

```
$ docker-compose up --build -d
```

After the application and database starts, make sure to run the migration command to create the database tables. This can be done by using the migrate CLI command below, assuming that the default db values are being used:

```
$ migrate -database "postgres://postgres:topsecret@localhost:5432/events_db?sslmode=disable" -path migrations up
```

## Endpoints:

| Method | Endpoint     | Description        |
| ------ | ------------ | ------------------ |
| POST   | /events      | Create a new event |
| GET    | /events      | List events        |
| GET    | /events/{id} | Fetch event by id  |

## Example CURL requests that can be used for testing

### POST /events

```
curl --request POST \
  --url http://localhost:8080/events \
  --header 'Content-Type: application/json' \
  --data '{
  "title": "Title 1",
  "description": "Test description",
  "start_time": "2023-10-27T10:30:00+01:00",
  "end_time": "2023-11-27T10:30:00+01:00"
}
'
```

### GET /events

```
curl --request GET \
  --url http://localhost:8080/events \
  --header 'Content-Type: application/json'
```

### GET /events/{id}

```
curl --request GET \
  --url http://localhost:8080/events/bbb8b0ae-875a-4d59-bd13-017211350be4 \
  --header 'Content-Type: application/json'
```
