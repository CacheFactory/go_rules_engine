package rulesengine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type RulesConfig struct {
	Condition Rule              `json:condition`
	Subject   string            `json:subject`
	Outcome   string            `json:outcome`
	Data      map[string]string `json:data`
}

type RulesEngine struct {
	Config  RulesConfig
	stack   []RuleResult
	Outcome string
}

type Operator struct {
	Term    string
	Func    func(*Rule, *Rule) interface{}
	Explain func(*Rule, *Rule) string
}

type RulesResponse struct {
	Outcome     string `json:outcome`
	Explanation string `json:explanation`
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
		return toFloat64(v.Run(re), re)
	case Rule:
		return toFloat64(v.Run(re), re)

	case string:
		retVal, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("error parsing float")
			return 0
		}
		return retVal
	case int:
		return float64(v)
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
		return toBool(v.Run(re), re)
	case *Rule:
		return toBool(v.Run(re), re)
	case bool:
		return v
	case string:
		if v == "true" {
			return true
		} else {
			return false
		}
	case int:
		if v != 1 {
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
		fmt.Println(value)
		fmt.Println("unknown type for toBool")
	}
	return false
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
		return toString(v.Run(re), re)
	case *Rule:
		return toString(v.Run(re), re)
	case string:
		return v
	case int:
		return strconv.Itoa(v)
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

func (re *RulesEngine) Run() (string, []RuleResult) {
	re.stack = []RuleResult{}
	result := re.runRule(re.Config.Condition)
	re.Outcome = result
	return result, re.stack
}

func (re *RulesEngine) Operators() map[string]Operator {
	return map[string]Operator{
		">=": {
			Term: "greater than or equal to",
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) >= toFloat64(right, re)
			},
			// Explain: func(left *Rule, right *Rule) string {
			// 	return toFloat64(left, re) >= toFloat64(right, re)
			// },
		},
		">": {
			Term: "greater than",
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) > toFloat64(right, re)
			},
		},
		"<=": {
			Term: "less than or equal to",
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) <= toFloat64(right, re)
			},
		},
		"<": {
			Term: "less than",
			Func: func(left *Rule, right *Rule) interface{} {
				return toFloat64(left, re) < toFloat64(right, re)
			},
		},
		"EQL": {
			Term: "equal to",
			Func: func(left *Rule, right *Rule) interface{} {
				return toString(left, re) == toString(right, re)
			},
		},
		"AND": {
			Term: "and",
			Func: func(left *Rule, right *Rule) interface{} {
				return toBool(left, re) && toBool(right, re)
			},
		},
		"GET": {
			Term: "and",
			Func: func(left *Rule, right *Rule) interface{} {
				return re.Config.Data[toString(right, re)]
			},
			Explain: func(left *Rule, right *Rule) string {
				return toString(right, re)
			},
		},
	}
}

func (re *RulesEngine) Explain() string {
	var explanation []string
	for _, result := range re.stack {
		explanation = append(explanation, result.Explanation)
	}

	str := strings.Join(explanation, " ")

	return str
}

func (re *RulesEngine) JsonResponse() RulesResponse {

	return RulesResponse{
		Outcome:     re.Outcome,
		Explanation: re.Explain(),
	}

}

func (re *RulesEngine) runRule(rule Rule) string {

	// result := RuleResult{
	// 	Outcome: rule.Run(re),
	// }

	return toString(rule, re)

}
