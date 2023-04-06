Cookbook - 
Experimenting a "real world app" development 
1. generate lib (nx-go)
2. root level installation of ent cli: `go get -d entgo.io/ent/cmd/ent`
3. cd to lib && generate schema: `go run -mod=mod entgo.io/ent/cmd/ent new Weather` (todo: wrap with nx generator) 
4. Edit the schema field 
5. Generate the code: `go generate ./ent`
6. run tests cmd: `nx run weather-lib:test` 
7. schema changes watch mode: `ent generate --target ./ent/schema --watch`

Weather pipeline flow:

	// query weather api search forecast
	// push results to topics according to country/region (collect)
	// consume & transform per country/region
	// store new struct in pg

	// API:
	// model gql schema + generate API
	// UI consumer (like the iphone app)

Kafka - followed confluent tutorial

- https://developer.confluent.io/get-started/go/#create-topic

`docker compose exec broker \
kafka-topics --create \
--topic ta-forecast \
--bootstrap-server localhost:9092 \
--replication-factor 1 \
--partitions 1`

Development:

Todo - move following to containers/weather
1. The service Dockerfile
2. The docker-compose.yaml

