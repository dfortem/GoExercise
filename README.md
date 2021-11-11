# Description
This is an exercise to get familiar with Golang, elastic search clusters and gRPC.

Here is the list of some of the documentation I've used to help me out. 

https://golang.org/doc/tutorial/create-
https://logz.io/blog/brew-install-elasticsearch-mac/
https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html
https://blog.logrocket.com/making-http-requests-in-go/
https://grpc.io/docs/languages/go/quickstart/
https://grpc.io/docs/languages/go/basics/

# Prerequisites
In order to run the code you'll need an elasticsearch cluster to monitor. If you don't have one, install elasticsearch on your machine and run it locally.

# Run the sample code
To compile and run the server, assuming you are in the root of the repository
folder, simply:

```sh
$ go run server/server.go
```

On an other terminal, run the client:

```sh
$ go run client/client.go <elasticsearchUrl>
```

# Alternative use
You can also just run the monitoring module by running the main app directly with:

```sh
$ go run .
```
