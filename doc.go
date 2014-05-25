// Gibbon is a very simple way to setup middleware, that uses net/http Handler interface
//
// http://github.com/claudiuandrei/gibbon
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/claudiuandrei/gibbon"
//    "net/http"
//  )
//
//  func main() {
//
//    g := gibbon.NewApp()
//
//    g.Use(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//      w.Header().Add("X-Powered-By", "Gibbon")
//    }))
//
//    mux := http.NewServeMux()
//    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//      fmt.Fprintf(w, "Welcome to Gibbon!")
//    })
//    g.Use(mux)
//
//    g.Run(":3000")
//  }
//
package gibbon
