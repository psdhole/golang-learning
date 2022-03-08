package retrybackoff

import (
	"errors"
	"log"
	"time"

	"github.com/cenkalti/backoff"
)

func TestNewExponentialBackOff() {

	operation := func() error {
		log.Println("trying to get the result!")
		// some API call
		return errors.New("test error")
	}

	log.Println("Start...")

	backOff := backoff.NewExponentialBackOff()
	backOff.InitialInterval = 1 * time.Second
	backOff.MaxElapsedTime = 60 * time.Second
	backOff.Multiplier = 2

	backOffMaxRetry := backoff.WithMaxRetries(backOff, 5)

	var err = backoff.Retry(operation, backOffMaxRetry)
	if err != nil {
		log.Fatalf("Error while getting result : %v", err)
	}
	log.Println("End...")
}

func TestNewConstantBackOff() {

	operation := func() error {
		log.Println("trying to get the result!")

		return errors.New("test error")

	}

	log.Println("Start...")

	b := backoff.NewConstantBackOff(1 * time.Second)

	var err = backoff.Retry(operation, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
	log.Println("End...")
}
