## A rules engine that gives human readable explanations

#### Have you ever wanted to explain in plan english why some condition is true or false? 

This project takes an abstract syntax tree (AST) of conditions and data to run checks against and gives you its result along with a human readable explanation.

## Example:

**Condition:** Is this person eligible for the US Military draft? 

They need to be over 18 years old and a male.

```
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
	"data": {
		"age": "18",
		"gender": "M"
	}
}
```

**Result:** TRUE because "M IS EQUAL TO M (gender) AND 18 (age) IS GREATER THAN OR EQUAL TO 18"

## Run server
`go run main.go`

## Example curl command
**Request**
`curl -H "Content-Type: application/json" --request POST -d '{"condition":{"operator":"AND","left_operand":{"operator":"EQL","left_operand":{"value":"M"},"right_operand":{"operator":"GET","right_operand":{"value":"gender"}}},"right_operand":{"operator":">=","left_operand":{"operator":"GET","right_operand":{"value":"age"}},"right_operand":{"value":"18"}}},"subject":"Eddie","outcome":"eligible for the draft","data":{"age":"18","gender":"M"}}' http://localhost:8090/rules_engine`

**Response**
```
{
  "Outcome": "true",
  "Explanation": "M IS EQUAL TO M (gender) AND 18 (age) IS GREATER THAN OR EQUAL TO 18",
  "ConditionResults": {
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
        },
        "result": "M",
        "explanation": "M (gender)"
      },
      "result": "true",
      "explanation": "M IS EQUAL TO M (gender)"
    },
    "right_operand": {
      "operator": ">=",
      "left_operand": {
        "operator": "GET",
        "right_operand": {
          "value": "age"
        },
        "result": "18",
        "explanation": "18 (age)"
      },
      "right_operand": {
        "value": "18"
      },
      "result": "true",
      "explanation": "18 (age) IS GREATER THAN OR EQUAL TO 18"
    },
    "result": "true",
    "explanation": "M IS EQUAL TO M (gender) AND 18 (age) IS GREATER THAN OR EQUAL TO 18"
  }
}
```


## TODO
Expand operators. Test more complicated conditions.