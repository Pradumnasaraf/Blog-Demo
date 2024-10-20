package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	flagsmith "github.com/Flagsmith/flagsmith-go-client/v3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Loading environment variable from the host system")
	} else {
		log.Printf("Loading environment from .env file")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		err, remainingLimit := rateLimitCall(c.ClientIP())
		if err != nil {
			c.JSON(
				http.StatusTooManyRequests,
				gin.H{"error": "Rate Limit Hit"})
		} else {
			c.JSON(
				http.StatusOK,
				gin.H{"Your left over API request is": remainingLimit})
		}
	})
	r.GET("/beta", func(c *gin.Context) {
		flags := getFeatureFlags()
		isEnabled, _ := flags.IsFeatureEnabled("beta")
		if isEnabled {
			c.JSON(
				http.StatusOK,
				gin.H{"message": "This is beta endpoint"})
		} else {
			c.String(http.StatusNotFound, "404 page not found")
		}
	})

	r.Run(":" + os.Getenv("PORT"))
}

func rateLimitCall(ClientIP string) (error, int) {

	ctx := context.Background()

	flags := getFeatureFlags()
	rateLimitInterface, _ := flags.GetFeatureValue("rate_limit")
	RATE_LIMIT := int(rateLimitInterface.(float64))
	fmt.Println("Current Rate Limit is", RATE_LIMIT)

	// Creating a new redis client to store the rate limit
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_URL")})
	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(ctx, ClientIP, redis_rate.PerHour(RATE_LIMIT))
	if err != nil {
		panic(err)
	}

	if res.Remaining == 0 {
		return errors.New("You have hit the Rate Limit for the API. Try again later"), 0
	}

	fmt.Println("remaining request for", ClientIP, "is", res.Remaining)
	return nil, res.Remaining
}

func getFeatureFlags() flagsmith.Flags {
	ctx := context.Background()
	client := flagsmith.NewClient(os.Getenv("FLAGSMITH_ENVIRONMENT_KEY"))
	flags, _ := client.GetEnvironmentFlags(ctx)
	return flags
}
