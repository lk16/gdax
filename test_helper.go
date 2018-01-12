package gdax

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

var sharedTestPublicClient *Client

func initTestClients() {
	sharedTestPublicClient = NewPublicClient(
		&http.Client{
			Timeout: 15 * time.Second,
		},
	)
	sharedTestReadOnlyClient = NewClient(
		sharedTestPublicClient.httpClient,
		os.Getenv("COINBASE_SECRET_RO"),
		os.Getenv("COINBASE_KEY_RO"),
		os.Getenv("COINBASE_PASSPHRASE_RO"),
	)
	sharedTestReadOnlyClient.publicLimiter = sharedTestPublicClient.publicLimiter
	sharedTestReadOnlyClient.privateLimiter = sharedTestPublicClient.privateLimiter

	sharedTestReadWriteClient = NewClient(
		sharedTestPublicClient.httpClient,
		os.Getenv("COINBASE_SECRET"),
		os.Getenv("COINBASE_KEY"),
		os.Getenv("COINBASE_PASSPHRASE"),
	)
	sharedTestReadWriteClient.publicLimiter = sharedTestPublicClient.publicLimiter
	sharedTestReadWriteClient.privateLimiter = sharedTestPublicClient.privateLimiter
}

func testPublicClient() *Client {
	if sharedTestPublicClient == nil {
		initTestClients()
	}
	return sharedTestPublicClient
}

var sharedTestReadOnlyClient *Client

func testReadOnlyClient() *Client {
	if sharedTestReadOnlyClient == nil {
		initTestClients()
	}
	return sharedTestReadOnlyClient
}

var sharedTestReadWriteClient *Client

func testReadWriteClient() *Client {
	if sharedTestReadWriteClient == nil {
		initTestClients()
	}
	return sharedTestReadWriteClient
}

func structHasZeroValues(i interface{}) bool {
	iv := reflect.ValueOf(i)

	//values := make([]interface{}, v.NumField())

	for i := 0; i < iv.NumField(); i++ {
		field := iv.Field(i)
		if reflect.Zero(field.Type()) == field {
			return true
		}
	}

	return false
}

func compareFields(a, b interface{}, properties []string) error {
	aValueOf := reflect.ValueOf(a)
	bValueOf := reflect.ValueOf(b)

	for _, field := range properties {
		aValue := reflect.Indirect(aValueOf).FieldByName(field).Interface()
		bValue := reflect.Indirect(bValueOf).FieldByName(field).Interface()

		if aValue != bValue {
			return errors.New(fmt.Sprintf("for field %s: (%s) not equal to (%s)", field, aValue, bValue))
		}
	}

	return nil
}
