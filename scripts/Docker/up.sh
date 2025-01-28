docker run -d \
  --name gopher_social_db \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=social \
  -p 5432:5432 \
  -v postgres_data:/var/lib/postgresql/data \
  postgres:16.3