<div align="center">
  <h3 align="center">Blog web scraper API with GO and PosgreSQL</h3>
  <p align="center">
    RSS aggregator backend API that allows clients to post feeds, that would be automatically and concurrently scraped.
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## Quick about

API endpoints are set up that allows the client to create a user and perform authenticated (by API key) requests with GO auth middleware.
Client can post feeds, follow them and get the most recently scraped information that was saved in the postgreSQL database.
Web scraper uses the base http library and go concurrency to perform multiple fetches at the same time and proccess each.

### Built With

 ![GO]
 ![PostgreSQL]

Side:
* Docker (to set up database container)
* SQLC for model and SQL query generation
* Chi router
* Goose for database migrations

## Set up

Clone the repository and install SQLC cmd
```go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest```

Then The Goose cmd
```go install github.com/pressly/goose/v3/cmd/goose@latest```

Add your PORT and database URL to .env e.g:
```
PORT=8000
DB_URL=postgres://user:password@localhost:5432/mydb?sslmode=disable
```
(The basic docker-compose.yaml file just has basic credentials hardcoded in for local use, change it with env files if need)

Lastly run
```source .env && goose -dir sql/schema postgres "$DB_URL" up``` and ```sqlc generate```
For model generation and schema migrations

To start the server:
```go build && ./rss-aggregator```

## Endpoints
* ```http://localhost:8000/v1/healthz```
* ```http://localhost:8000/v1/users``` With Header ```Authorization ApiKey {insertApiKey}```
* ```http://localhost:8000/v1/feeds```
* ```http://localhost:8000/v1/feed_follows```
* ```http://localhost:8000/v1/posts/my```
* ```http://localhost:8000/v1/feed_follows/{feedFollowId}```


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[GO]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge
[PostgreSQL]: https://img.shields.io/badge/postgresql-4169e1?style=for-the-badge&logo=postgresql&logoColor=white
