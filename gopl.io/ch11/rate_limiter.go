package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
	"gopl.io/ch11/limiter"
)

func Open() *ApiConnection {

	rateLimiter1 := rate.NewLimiter(rate.Limit(1), 1)
	rateLimiter2 := rate.NewLimiter(rate.Limit(5), 4)
	return &ApiConnection{
		// limit - number of events per second, burst - bucket size
		rateLimiter: limiter.MultiLimiter(rateLimiter1, rateLimiter2),
	}
}

type ApiConnection struct {
	rateLimiter limiter.RateLimiter
}

func (a *ApiConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *ApiConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func main() {
	defer log.Printf("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: &v", err)
			}
			log.Printf("ReadFile")
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: &v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}
