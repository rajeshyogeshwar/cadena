package cadena

import (
	"github.com/julienschmidt/httprouter"
)

// Middleware is a struct for middlewares to be implemented
type Middleware func(httprouter.Handle) httprouter.Handle

// ChainMiddlewares is a simple method that accepts the httprouter.Handle and list of middlewares to apply to the handle and return httprouter.Handle
func ChainMiddlewares(fn httprouter.Handle, middlewares ...Middleware) httprouter.Handle {

	// If the middleware to be applied is the last one to be applied, simply call the handler
	if len(middlewares) == 0 {
		return fn
	}

	wrappedHandler := fn

	// Reverse looping the provided middlewares to preserve the order. Last one is applied first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrappedHandler = middlewares[i](wrappedHandler)
	}

	// After all the middlewares are applied, returns the original handler
	return wrappedHandler

}
