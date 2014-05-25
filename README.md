# Gibbon

Gibbon is a very simple way to setup middleware, that uses `net/http` Handler interface.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

~~~ go
package main

import (
  "fmt"
  "github.com/claudiuandrei/gibbon"
  "net/http"
)

func main() {

  // Setup the app
  g := gibbon.NewApp()

  // Setup a middleware
  g.Use(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("X-Powered-By", "Gibbon")
  }))

  // Setup the router
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to Gibbon!")
  })
  g.Use(mux)

  // Run the server
  g.Run(":3000")
}
~~~

Then install the Gibbon package (**go 1.1** and greater is required):
~~~
go get github.com/claudiuandrei/gibbon
~~~

Then run your server:
~~~
go run server.go
~~~

You will now have a Go net/http webserver running on `localhost:3000`.

## `Use()` - Middleware & Routing

Gibbon `Use` function adds middleware to the request execution chain. Everything that acts like a `net/http` Handler can be a middleware and will run through each of them in the order added until the ResponseWriter is flushed.

Gibbon is BYOR (Bring your own Router). Gibbon is fully supporting `net/http`.

~~~ go
// New Application
g := gibbon.NewApp()

// Middleware
g.Use(FirstMiddleware)

// Setup the router
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Welcome to Gibbon!")

  // Flush the response
  // This will break the execution chain and the remaining middleware is skipped
  w.(http.Flusher).Flush()
})
g.Use(mux)

// More middleware
g.Use(LastMiddleware)

// Run the server
g.Run(":3000")
~~~

## `Run()` - Server

Gibbon has a convenience function called `Run`. `Run` takes an addr string identical to [http.ListenAndServe](http://golang.org/pkg/net/http#ListenAndServe).

## Authors

[Claudiu Andrei](http://claudiuandrei.com/)