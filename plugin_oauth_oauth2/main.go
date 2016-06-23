package main

import (
	"html/template"
	"sort"
	"strings"

	"github.com/iris-contrib/plugin/oauth"
	"github.com/kataras/iris"
)

// register your auth via configs, providers with non-empty values will be registered to goth automatically by Iris
var configs = oauth.Config{
	Path: "/auth", //defaults to /auth

	GithubKey:    "YOUR_GITHUB_KEY",
	GithubSecret: "YOUR_GITHUB_SECRET",
	GithubName:   "github", // defaults to github

	FacebookKey:    "YOUR_FACEBOOK_KEY",
	FacebookSecret: "YOUR_FACEBOOK_KEY",
	FacebookName:   "facebook", // defaults to facebook
}

func main() {
	// create the plugin with our configs
	authentication := oauth.New(configs)
	// register the plugin to iris
	iris.Plugins.Add(authentication)

	m := make(map[string]string)
	m[configs.GithubName] = "Github" // same as authentication.Config.GithubName
	m[configs.FacebookName] = "Facebook"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	// set a  login success handler( you can use more than one handler)
	// if user succeed to logged in
	// client comes here from: localhost:3000/auth/lowercase_provider_name/callback
	authentication.Success(func(ctx *iris.Context) {
		// if user couldn't validate then server sends StatusUnauthorized, which you can handle by:  authentication.Fail OR iris.OnError(iris.StatusUnauthorized, func(ctx *iris.Context){})
		user := authentication.User(ctx)

		// you can get the url by the predefined-named-route 'oauth' which you can change by Config's field: RouteName
		println("came from " + iris.URL("oauth", strings.ToLower(user.Provider)))

		t, _ := template.New("foo").Parse(userTemplate)
		ctx.ExecuteTemplate(t, user)
	})

	// customize the error page using: authentication.Fail(func(ctx *iris.Context){....})

	iris.Get("/", func(ctx *iris.Context) {
		t, _ := template.New("foo").Parse(indexTemplate)
		ctx.ExecuteTemplate(t, providerIndex)
	})

	iris.Listen(":3000")
}

// ProviderIndex ...
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="` + configs.Path + `/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
