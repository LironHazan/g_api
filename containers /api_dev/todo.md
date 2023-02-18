Instructions:
```
export MY_PASSWORD=12345
docker build -t my-postgres-image .
docker run --name my-postgres-container -p 5432:5432 -e POSTGRES_PASSWORD=$MY_PASSWORD -d my-postgres-image
```
