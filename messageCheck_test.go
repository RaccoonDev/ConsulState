package main

import (
	"reflect"
	"testing"
)

// func TestMessasgeChecker(t *testing.T) {
// 	registrations := []MessageRegistration{
// 		{
// 			RegistrationID: "test-message-registration",
// 			ServiceName:    "testService",
// 			Pattern:        `{ "name": "test-message" }`,
// 		},
// 		{
// 			RegistrationID: "test-messageA-registration",
// 			ServiceName:    "testServiceA",
// 			Pattern:        `{ "name": "test-messageA" }`,
// 		},
// 	}

// 	msgChkr := NewMessageChecker(registrations)

// 	r, err := msgChkr.GetMessageRegistration(`{"id":"1234", "name": "test-message", "payload": { "expertId": 123} }`)
// 	if err != nil {
// 		t.Fatal("Message checker must except simple message")
// 	}

// 	if r.ServiceName != "testService" {
// 		t.Fatal("Test service must be detected by from the registration")
// 	}
// }

func TestCheckPattern(t *testing.T) {
	patternMap := make(map[string]interface{})
	messageMap := make(map[string]interface{})
	isMatch := checkPattern(patternMap, messageMap)

	if !isMatch {
		t.Error("Two empty maps must match")
	}

	patternMap["key1"] = 123
	isMatch = checkPattern(patternMap, messageMap)
	if isMatch {
		t.Error("Empty message must not match with pattern with one key")
	}

	messageMap["key1"] = 123
	isMatch = checkPattern(patternMap, messageMap)
	if !isMatch {
		t.Error("map with same keys and value should match")
	}

	messageMap["key2"] = 123
	isMatch = checkPattern(patternMap, messageMap)
	if !isMatch {
		t.Error("if message has more fields then pattern it still should match ")
	}

	patternMap["key2"] = 124
	isMatch = checkPattern(patternMap, messageMap)
	if isMatch {
		t.Error("message should match by all values")
	}

	messageMap["key2"] = 124
	isMatch = checkPattern(patternMap, messageMap)
	if !isMatch {
		t.Error("identical message and pattern should match")
	}

	patternMap["key3"] = 123
	isMatch = checkPattern(patternMap, messageMap)
	if isMatch {
		t.Error("should not match if pattern has more keys then messsage")
	}

	messageMap["key3"] = "someValue"
	isMatch = checkPattern(patternMap, messageMap)
	if isMatch {
		t.Error("should not match if message has another type")
	}
}

func TestParseToMap(t *testing.T) {
	testJSON := `{ "name": "Bob", "age": 33 }`
	parsedMap, err := parseToMap(testJSON)
	if err != nil {
		t.Error("this test json must be parsed")
	}

	if parsedMap["name"] != "Bob" {
		t.Error("Name property must be parsed out")
	}

	if parsedMap["age"] != float64(33) {
		t.Errorf("Age property must be parsed out, expected 33 got %+v of type %+v", parsedMap["age"], reflect.TypeOf(parsedMap["age"]))
	}
}
