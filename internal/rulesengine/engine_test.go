package rulesengine

import (
	"fmt"
	"testing"
)

func TestRule(t *testing.T) {
	rulesJson := `{
		"condition": {
			"operator": "<=",
			"left_operand": {
				"operator": "GET",
				"right_operand": {"value": "age"}
			}, 
			"right_operand": {"value": "40"}
		},

		"subject": "Eddie",
		"outcome": "eligible for the draft",

		"data": {
			"age": "35"
		}
	}`

	rulesEngine, err := FromJson(rulesJson)

	if err != nil {
		panic(err)
	}

	result, stack := rulesEngine.Run()

	if result != "true" {
		t.Errorf("incorrect result")
	}

	fmt.Println(stack)

}
