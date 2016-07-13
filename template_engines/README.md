## Folder information

This folder shows you mini examples on how to use the built'n Iris' template engines.

Most examples are only written for the `template/html` engine but exists for other template engines also, except the layout functionality which is not implemented for django and amber, because they have their own way to implement this functionality.

> Note: markdown template engine doesn't permit any iris built'n functions, neither the layout, it is used only for special cases.


- **Q: Can I use more than one template engine for the same app?**
A: Yes you can, each template engine has its own file extension, so you can regiser html template engine with .html and pug/jade for .jade extension, the context.Render will find the correct template engine to render the template.

