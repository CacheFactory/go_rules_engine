## A rules engine and an HTTP server that serves it up.

#### Have you ever wanted to explain in plan english why some condition is true or false? 

Takes an abstract syntax tree (AST) of conditions and data to run checks against and gives you its boolean result along with a human readable explanation.

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

## TODO
Expand operators. Test more complicated conditions.