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

func testPublicClient() *Client {
	if sharedTestPublicClient == nil {
		sharedTestPublicClient = NewPublicClient(
			&http.Client{
				Timeout: 15 * time.Second,
			},
		)
	}
	return sharedTestPublicClient
}

var sharedTestReadOnlyClient *Client

func testReadOnlyClient() *Client {
	if sharedTestReadOnlyClient == nil {
		sharedTestReadOnlyClient = NewClient(
			&http.Client{
				Timeout: 15 * time.Second,
			},
			os.Getenv("COINBASE_SECRET_RO"),
			os.Getenv("COINBASE_KEY_RO"),
			os.Getenv("COINBASE_PASSPHRASE_RO"),
		)
	}
	return sharedTestReadOnlyClient
}

var sharedTestReadWriteClient *Client

func testReadWriteClient() *Client {
	if sharedTestReadWriteClient == nil {
		sharedTestReadWriteClient = NewClient(
			&http.Client{
				Timeout: 15 * time.Second,
			},
			os.Getenv("COINBASE_SECRET"),
			os.Getenv("COINBASE_KEY"),
			os.Getenv("COINBASE_PASSPHRASE"),
		)
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

func compareProperties(a, b interface{}, properties []string) (bool, error) {
	aValueOf := reflect.ValueOf(a)
	bValueOf := reflect.ValueOf(b)

	for _, property := range properties {
		aValue := reflect.Indirect(aValueOf).FieldByName(property).Interface()
		bValue := reflect.Indirect(bValueOf).FieldByName(property).Interface()

		if aValue != bValue {
			return false, errors.New(fmt.Sprintf("%s not equal", property))
		}
	}

	return true, nil
}
