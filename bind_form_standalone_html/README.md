## Folder information


This folder contains an example on how to use the ReadForm to bind a form  to object.
You may see other frameworks uses binders and all that, but you know better than me that Iris uses simplicity for productivity.
In my opinion this method is easier than others.


> You can use ctx.ReadJSON & ctx.ReadXML also, instead of reading the form data these functions are reading the whole page Body and convert it to an object.


## How to use


```go

//...

type Visitor struct {
	Username string
	VisitDate string `formam: "visit_date"`
}

func(ctx *iris.Context){
	visitor := &Visitor{}
	ctx.ReadForm(visitor)
	//....
}

//...

```
