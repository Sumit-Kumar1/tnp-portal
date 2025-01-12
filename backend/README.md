# TNP-portal backend

- Backend based on gofr framework
- run the following command to create a postgres db:
  - `docker run --name tnp-backend -e POSTGRES_DB=tnp_backend -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:latest`