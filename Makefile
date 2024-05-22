# Makefile
build:
    docker build -t myapp .

run:
    docker run --env DATABASE_URL=<your_database_url> -p 8000:8000 myapp
