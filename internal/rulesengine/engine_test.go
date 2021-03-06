package rulesengine

import (
	"testing"
)

func TestRule1(t *testing.T) {
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
			"age": "18",
			"gender": "M"
		}
	}
	`

	rulesEngine, err := FromJson(rulesJson)

	if err != nil {
		panic(err)
	}

	result, _ := rulesEngine.Run()

	if result != "true" {
		t.Errorf("incorrect result")
	}

	if rulesEngine.Explanation != "M IS EQUAL TO M (gender) AND 18 (age) IS GREATER THAN OR EQUAL TO 18" {
		t.Errorf("incorrect explanation")
	}

	if rulesEngine.JsonResponse().ConditionResults.Explanation != "M IS EQUAL TO M (gender) AND 18 (age) IS GREATER THAN OR EQUAL TO 18" {
		t.Errorf("incorrect explanation in JSON payload")
	}

}

func TestRule2(t *testing.T) {
	rulesJson := `
	{
		"condition": {
			"operator": "OR",
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
			"age": "10",
			"gender": "M"
		}
	}
	`

	rulesEngine, err := FromJson(rulesJson)

	if err != nil {
		panic(err)
	}

	result, _ := rulesEngine.Run()

	if result != "true" {
		t.Errorf("incorrect result")
	}

	if rulesEngine.Explanation != "M IS EQUAL TO M (gender) OR 10 (age) IS NOT GREATER THAN OR EQUAL TO 18" {
		t.Errorf("incorrect explanation")
	}

}

func TestRule3(t *testing.T) {
	rulesJson := `
	{
		"condition": {
			"operator": "GET",
			"right_operand": {
				"value": "gender"
			}
		},
		"subject": "Eddie",
		"data": {
			"age": "10",
			"gender": "M"
		}
	}
	`

	rulesEngine, err := FromJson(rulesJson)

	if err != nil {
		panic(err)
	}

	result, _ := rulesEngine.Run()

	if result != "M" {
		t.Errorf("incorrect result")
	}

	if rulesEngine.Explanation != "M (gender)" {
		t.Errorf("incorrect explanation")
	}

}
