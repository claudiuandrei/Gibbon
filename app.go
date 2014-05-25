package gibbon

import "net/http"

// Simple handler with middleware attached
type App struct {
	http.Handler
	middleware []http.Handler
}

// Entry point into the applicaton
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Load the response writer
	rw, ok := w.(*ResponseWriter)

	// Convert the type
	if !ok {
		rw = NewResponseWriter(w)
	}

	// Just go through the stack
	for _, h := range a.middleware {

		// Response has been
		if rw.Flushed {
			return
		}

		// Serve middleware
		h.ServeHTTP(rw, r)
	}
}

// Create a new app
func NewApp() *App {
	return &App{}
}

// Add handlers to the middleware stack
func (a *App) Use(h http.Handler) {
	a.middleware = append(a.middleware, h)
}

// Run the server
func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a)
}
