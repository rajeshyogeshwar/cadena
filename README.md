# cadena
Middleware is piece of code that will be executed on every incoming request. Upon the execution the handler for the requested url/endpoint/route is to be executed. If it is only one middleware that is to be implemented it is fairly simple. But usually, whenever we are developing a fairly sizeable web application we might end up with multiple middlewares to be implemented. This middlewares need to be chained, meaning the output of processing of one middleware becomes the input to next middleware that is wrapped around the current.

In go, [httprouter](https://github.com/julienschmidt/httprouter) is one of the very popular mux or request router that is being used. *cadena* is a simple utility module that will chain all the provided middlewares and apply it to the incoming request before serving up the defined handler. 

Let's say we have following middlewares defined somewhere

    var LoggingMiddleware =  func(fn httprouter.Handle) httprouter.Handle {
    	return httprouter.Handle(
    		func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    				// Your code to be implemented in this middleware
    				fn(writer, request, params)
    		})
    }

    var AuthMiddleware =  func(fn httprouter.Handle) httprouter.Handle {
        return httprouter.Handle(
	        func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		    // Your code to be implemented in this middleware
		    fn(writer, request, params)
	    })
	}

Now somewhere in your application you might have defined router and the routes for your application handlers. Let's say it looks something like this.

    router := httprouter.New()
    router.GET("/", Home)
    router.GET("/", About)

What we want to do is apply our middlewares to the defined routes. We can modify our route definitions as follows.

    middlewaresToApply := []Middleware {
	    AuthMiddleware,
	    LoggingMiddleware,
    }

    router.GET("/", cadena.ChainMiddleware(Home, middlewaresToApply...))

The important thing to notice is the order of middlewares defined, last one is applied first and it moves upwards.