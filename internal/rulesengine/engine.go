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
	Config  RulesConfig
	stack   []RuleResult
	Outcome string
}

type RulesResponse struct {
	Outcome string `json:outcome`
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
		fmt.Println("A")
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

func toString(value interface{}, re *RulesEngine) string {
	switch v := value.(type) {
	case *Rule:
		if v.Value != "" {

			return toString(v.Value, re)
		}
	}

	switch v := value.(type) {
	case Rule:
		fmt.Println("A")
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

func (re *RulesEngine) Operators() map[string]func(*Rule, *Rule) interface{} {
	return map[string]func(*Rule, *Rule) interface{}{
		">=": func(left *Rule, right *Rule) interface{} {
			return toFloat64(left, re) >= toFloat64(right, re)
		},
		">": func(left *Rule, right *Rule) interface{} {
			return toFloat64(left, re) > toFloat64(right, re)
		},
		"<=": func(left *Rule, right *Rule) interface{} {

			return toFloat64(left, re) <= toFloat64(right, re)
		},
		"<": func(left *Rule, right *Rule) interface{} {
			return toFloat64(left, re) < toFloat64(right, re)
		},
		"==": func(left *Rule, right *Rule) interface{} {
			return toString(left, re) == toString(right, re)
		},
		"GET": func(left *Rule, right *Rule) interface{} {
			return re.Config.Data[toString(right, re)]
		},
	}
}

func (re *RulesEngine) JsonResponse() RulesResponse {

	return RulesResponse{
		Outcome: re.Outcome,
	}

}

func (re *RulesEngine) runRule(rule Rule) string {

	// result := RuleResult{
	// 	Outcome: rule.Run(re),
	// }

	return toString(rule, re)

}
