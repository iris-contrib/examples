package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"

	"github.com/kataras/iris/v12/_examples/sessions/overview/example"
)

// 1. Install Redis:
// 1.1 Windows: https://github.com/ServiceStack/redis-windows
// 1.2 Other: https://redis.io/download
// 2. Run command: go run -mod=mod main.go
//
// Tested with redis version 3.0.503.
func main() {
	// These are the default values,
	// you can replace them based on your running redis' server settings:
	db := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      getenv("REDIS_ADDR", "127.0.0.1:6379"),
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Username:  "",
		Password:  "",
		Database:  "",
		Prefix:    "myapp-",
		Driver:    redis.GoRedis(), // defaults to this driver.
		// To set a custom, existing go-redis client, use the "SetClient" method:
		// Driver: redis.GoRedis().SetClient(customGoRedisClient)
	})

	// Optionally configure the underline driver:
	// driver := redis.GoRedis()
	// // To set a custom client:
	// driver.SetClient(customGoRedisClient)
	// OR:
	// driver.ClientOptions = redis.Options{...}
	// driver.ClusterOptions = redis.ClusterOptions{...}
	// redis.New(redis.Config{Driver: driver, ...})

	defer db.Close() // close the database connection if application errored.

	sess := sessions.New(sessions.Config{
		Cookie:          "_session_id",
		Expires:         0, // defaults to 0: unlimited life. Another good value is: 45 * time.Minute,
		AllowReclaim:    true,
		CookieSecureTLS: true,
	})

	//
	// IMPORTANT:
	//
	sess.UseDatabase(db)

	app := example.NewApp(sess)

	// TIP scaling-out Iris sessions using redis:
	// $ docker-compose up
	// http://localhost:8080/set/$key/$value
	// The value will be available on all Iris servers as well.
	// E.g. http://localhost:9090/get/$key and vice versa.
	addr := fmt.Sprintf(":%s", getenv("PORT", "8080"))
	app.Listen(addr)
}

func getenv(key string, def string) string {
	if v := os.Getenv(strings.ToUpper(key)); v != "" {
		return v
	}

	return def
}
