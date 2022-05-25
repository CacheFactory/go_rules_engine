package rulesengine

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type RulesConfig struct {
	Condition Rule              `json:condition`
	Subject   string            `json:subject`
	Outcome   string            `json:outcome`
	Data      map[string]string `json:data`
}

type RulesEngine struct {
	Config      RulesConfig
	stack       []RuleResult
	Explanation string
	Outcome     string
}

type Operator struct {
	Term        string
	Func        func(*Rule, *Rule) interface{}
	Explanation func(*Rule, *Rule, bool) string
}

type RulesResponse struct {
	Outcome     string       `json:outcome`
	Explanation string       `json:explanation`
	Results     []RuleResult `json:results`
}

func isOrNot(x bool) string {
	if x == true {
		return "IS"
	} else {
		return "IS NOT"
	}
}

func toFloat64(value interface{}, re *RulesEngine) float64 {
	switch v := value.(type) {
	case *Rule:
		if v.Value != "" {
			return toFloat64(v.Value, re)
		}
	}

	switch v := value.(type) {
	case *Rule:
		result, _ := v.Run(re)
		return toFloat64(result, re)
	case Rule:
		result, _ := v.Run(re)
		return toFloat64(result, re)

	case string:
		retVal, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("error parsing float")
			return 0
		}
		return retVal
	case float64:
		return v
	default:

		fmt.Println("unknown type for toFloat64")
	}
	return 0
}

func toBool(value interface{}, re *RulesEngine) bool {
	switch v := value.(type) {
	case *Rule:
		if v.Value != "" {

			return toBool(v.Value, re)
		}
	}

	switch v := value.(type) {
	case Rule:
		result, _ := v.Run(re)
		return toBool(result, re)
	case *Rule:
		result, _ := v.Run(re)
		return toBool(result, re)
	case bool:
		return v
	case string:
		if v == "true" {
			return true
		} else {
			return false
		}
	case float64:
		if v != 1 {
			return true
		} else {
			return false
		}
	default:
		fmt.Println("unknown type for toBool")
	}
	return false
}

func toExplanation(value interface{}, re *RulesEngine) string {
	switch v := value.(type) {
	case *Rule:
		if v.Value != "" {

			return toString(v.Value, re)
		}
	}

	switch v := value.(type) {
	case Rule:
		result := re.Operators()[v.Operator].Func(v.LeftOperand, v.RightOperand)
		return re.Operators()[v.Operator].Explanation(v.LeftOperand, v.RightOperand, toBool(result, re))
	case *Rule:
		result := re.Operators()[v.Operator].Func(v.LeftOperand, v.RightOperand)
		return re.Operators()[v.Operator].Explanation(v.LeftOperand, v.RightOperand, toBool(result, re))
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 64)
	case bool:
		if v == true {
			return "true"
		} else {
			return "false"
		}

	default:
		fmt.Println("unknown type for toExplanation")
	}
	return ""
}

func toString(value interface{}, re *RulesEngine) string {
	switch v := value.(type) {
	case *Rule:
		if v.Value != "" {

			return toString(v.Value, re)
		}
	}

	switch v := value.(type) {
	case Rule:
		result, _ := v.Run(re)
		return toString(result, re)
	case *Rule:
		result, _ := v.Run(re)
		return toString(result, re)
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 64)
	case bool:
		if v == true {
			return "true"
		} else {
			return "false"
		}

	default:
		fmt.Println("unknown type for toString")
	}
	return ""
}

func FromJson(jsonRulesConfig string) (*RulesEngine, error) {
	rulesConfig := RulesConfig{}

	err := json.Unmarshal([]byte(jsonRulesConfig), &rulesConfig)

	if err != nil {
		return nil, err
	}

	return New(rulesConfig), nil
}

func New(config RulesConfig) *RulesEngine {

	return &RulesEngine{
		Config: config,
	}
}

func (re *RulesEngine) Run() (string, string) {
	result, explanation := re.runRule(re.Config.Condition)
	re.Outcome = result
	re.Explanation = explanation
	return result, explanation
}

func (re *RulesEngine) Operators() map[string]Operator {
	return map[string]Operator{
		">=": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) >= toFloat64(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return toExplanation(left, re) + " " + isOrNot(positiveResult) + " GREATER THAN OR EQUAL TO " + toExplanation(right, re)
			},
		},
		">": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) > toFloat64(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return toExplanation(left, re) + " " + isOrNot(positiveResult) + " GREATER THAN " + toExplanation(right, re)
			},
		},
		"<=": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) <= toFloat64(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return toExplanation(left, re) + " " + isOrNot(positiveResult) + " LESS THAN OR EQUAL TO " + toExplanation(right, re)
			},
		},
		"<": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) < toFloat64(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return toExplanation(left, re) + " " + isOrNot(positiveResult) + " LESS THAN " + toExplanation(right, re)
			},
		},
		"EQL": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toString(left, re) == toString(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return toExplanation(left, re) + " " + isOrNot(positiveResult) + " EQUAL TO " + toExplanation(right, re)
			},
		},
		"AND": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toBool(left, re) && toBool(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				// ??
				return toExplanation(left, re) + " AND " + toExplanation(right, re)
			},
		},
		"OR": {
			Func: func(left *Rule, right *Rule) interface{} {
				return toBool(left, re) || toBool(right, re)
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				// ??
				return toExplanation(left, re) + " OR " + toExplanation(right, re)
			},
		},
		"GET": {
			Func: func(left *Rule, right *Rule) interface{} {
				return re.Config.Data[toString(right, re)]
			},
			Explanation: func(left *Rule, right *Rule, positiveResult bool) string {
				return re.Config.Data[toString(right, re)] + " (" + toString(right, re) + ")"
			},
		},
	}
}

func (re *RulesEngine) JsonResponse() RulesResponse {

	return RulesResponse{
		Outcome:     re.Outcome,
		Results:     re.stack,
		Explanation: re.Explanation,
	}

}

func (re *RulesEngine) runRule(rule Rule) (string, string) {

	result, explanation := rule.Run(re)

	return toString(result, re), explanation

}
