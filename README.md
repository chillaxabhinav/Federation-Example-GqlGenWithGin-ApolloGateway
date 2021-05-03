## How to run

- First Install Node application inside "gateway" folder using below command.
  > yarn install
- Then open 3 terminal windows and run all the go-gqlgen applications first inside "product", "reviews" and "user" folder in separate terminal windows using command below, these are acting as our graphql microservice application.
  > go run server.go
- Now open another terminal window and in that finally run the gateway using below command.
  > node server.js
- Now open http://localhost:4000/, This you gateway endpoint and in that you can write graphql query as per the docs in the playground made by merging all graphql schemas of running go microservices making into a federated schema, allowing us to write graphql queries across multiple graphql microservices.

## Libraries/Tech Used

- Go graphql microservices

  - gqlgen - https://gqlgen.com/
  - gin - https://gin-gonic.com/

  Gin is wrapping gqlgen graphql application allowing us to hook multiple middlewares to the incoming request like I have hooked a middleware for rate limiting.

- Gateway - https://www.apollographql.com/docs/federation/

  Gateway uses apollo federation which helps us to merge schemas and query across these multiple graphql microservices using its query planner.
