package gdax

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
	"github.com/shopspring/decimal"
)

var sharedTestPublicClient *Client

func initTestClients() {
	sharedTestPublicClient = NewClient(&http.Client{Timeout: 15 * time.Second,}, nil)

	roAuth := Authentication{
		Secret:     os.Getenv("COINBASE_SECRET_RO"),
		Key:        os.Getenv("COINBASE_KEY_RO"),
		Passphrase: os.Getenv("COINBASE_PASSPHRASE_RO"),
	}
	sharedTestReadOnlyClient = NewClient(sharedTestPublicClient.HttpClient, &roAuth)
	sharedTestReadOnlyClient.publicLimiter = sharedTestPublicClient.publicLimiter
	sharedTestReadOnlyClient.privateLimiter = sharedTestPublicClient.privateLimiter

	rwAuth := Authentication{
		Secret:     os.Getenv("COINBASE_SECRET"),
		Key:        os.Getenv("COINBASE_KEY"),
		Passphrase: os.Getenv("COINBASE_PASSPHRASE"),
	}
	sharedTestReadWriteClient = NewClient(sharedTestPublicClient.HttpClient, &rwAuth)
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

func compareAllFields(expected, actual interface{}) error {
	v := reflect.ValueOf(expected)
	vType := v.Type()
	fields := make([]string, 0, 10)
	for i := 0; i < v.Type().NumField(); i++ {
		fields = append(fields, vType.Field(i).Name)
	}
	return compareFields(expected, actual, fields)
}

func compareFields(expected, actual interface{}, properties []string) error {
	expectedValueOf := reflect.ValueOf(expected)
	actualValueOf := reflect.ValueOf(actual)

	for _, field := range properties {
		expectedValue := reflect.Indirect(expectedValueOf).FieldByName(field).Interface()
		actualValue := reflect.Indirect(actualValueOf).FieldByName(field).Interface()

		expectedDecimalRef, expectedIsDecimalRef := expectedValue.(*decimal.Decimal)
		if expectedIsDecimalRef && expectedDecimalRef != nil {
			expectedValue = *expectedDecimalRef
		}

		actualDecimalRef, actualIsDecimalRef := actualValue.(*decimal.Decimal)
		if actualIsDecimalRef && actualDecimalRef != nil {
			actualValue = *actualDecimalRef
		}

		expectedDecimal, expectedIsDecimal := expectedValue.(decimal.Decimal)
		actualDecimal, actualIsDecimal := actualValue.(decimal.Decimal)
		if expectedIsDecimal && actualIsDecimal {
			expectedValue, actualValue = expectedDecimal.String(), actualDecimal.String()
		}

		if expectedValue != actualValue {
			return errors.New(fmt.Sprintf("for field %s: (%s) not equal to (%s)", field, expectedValue, actualValue))
		}
	}

	return nil
}

func timeMust(time time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return time
}
