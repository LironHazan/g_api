
Generated gql api in order to repro a bug..

Install deps:

```
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/graphql/handler/transport@v0.17.26
go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.26
`````

Generate api project

`nx g @nx-go/nx-go:app graph-api
`

Init a project

`go run github.com/99designs/gqlgen init
`

Start graph API 

`nx run graph-api:serve`

todo: containerize 
