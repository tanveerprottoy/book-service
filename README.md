# Running the app

postgresql and book_service_db in postgres must be present for the service to run

# run in container with everything included
docker-compose up

# run service locally postgres in docker
``` start postgres if not available in the system through docker
docker compose -f postgres.yml up

```cli
go run cmd/bookservice/main.go
```
```makefile
make run
```