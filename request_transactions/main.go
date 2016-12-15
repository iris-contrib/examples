package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.BeginTransaction(func(scope *iris.RequestTransactionScope) {
			// OPTIONAL STEP:
			// create a new custom type of error here to keep track of the status code and reason message
			err := iris.NewErrWithStatus()

			// we should use scope.Context if we want to rollback on any errors lives inside this function clojure.
			// if you want persistence then use the 'ctx'.
			scope.Context.Text(iris.StatusOK, "Blablabla this should not be sent to the client because we will fill the err with a message and status")

			//	var firstErr error  = do this()   // your code here
			//	var secondErr error = try_do_this() // your code here
			//	var thirdErr error  = try_do_this() // your code here
			//	var fail bool = false

			//	if firstErr != nil || secondErr != nil || thirdErr != nil {
			//			fail = true
			//	}
			// or err.AppendReason(firstErr.Error()) // ... err.Reason(dbErr.Error()).Status(500)

			// virtualize a fake error here, for the shake of the example
			fail := true
			if fail {
				err.Status(iris.StatusInternalServerError).
					// if status given but no reason then the default or the custom http error will be fired (like ctx.EmitError)
					Reason("Error: Virtual failure!!")
			}

			// OPTIONAl STEP:
			// but useful if we want to post back an error message to the client if the transaction failed.
			// if the reason is empty then the transaction completed succesfuly,
			// otherwise we rollback the whole response body and cookies and everything lives inside the scope.Request.
			scope.Complete(err)
		})

		ctx.BeginTransaction(func(scope *iris.RequestTransactionScope) {
			scope.Context.HTML(iris.StatusOK,
				"<h1>This will sent at all cases because it lives on different transaction and it doesn't fails</h1>")
			// * if we don't have any 'throw error' logic then no need of scope.Complete()
		})

		// OPTIONAL, depends on the usage:
		// at any case, what ever happens inside the context's transactions send this to the client
		ctx.HTML(iris.StatusOK, "<h1>I persist show this message to the client whatever happens!</h1>")
	})

	iris.Listen(":8080")
}
