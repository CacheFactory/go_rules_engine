package rulesengine

import (
	"fmt"
	"testing"
)

func TestRule(t *testing.T) {
	rulesJson := `
	{
		"condition": {
			"operator": "AND",
			"left_operand": {
				"operator": "EQL",
				"left_operand": {
					"value": "M"
				},
				"right_operand": {
					"operator": "GET",
					"right_operand": {
						"value": "gender"
					}
				}
			},
			"right_operand": {
				"operator": ">=",
				"left_operand": {
					"operator": "GET",
					"right_operand": {
						"value": "age"
					}
				},
				"right_operand": {
					"value": "18"
				}
			}
		},
		"subject": "Eddie",
		"outcome": "eligible for the draft",
		"data": {
			"age": "35",
			"gender": "M"
		}
	}
	`

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
