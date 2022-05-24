package rulesengine

type Rule struct {
	Operator     string `json:"operator,omitempty"`
	LeftOperand  *Rule  `json:"left_operand,omitempty"`
	RightOperand *Rule  `json:"right_operand,omitempty"`
	Explanation  string `json:"explanation,omitempty"`
	Value        string `json:"value,omitempty"`
}

type RuleResult struct {
	Outcome     interface{}
	Left        string
	Right       string
	Operator    string
	Explanation string
}

func (rule *Rule) Run(re *RulesEngine) interface{} {
	operators := re.Operators()

	result := operators[rule.Operator].Func(rule.LeftOperand, rule.RightOperand)

	var explanation string

	if rule.LeftOperand != nil && rule.RightOperand != nil {
		explanation = toString(rule.LeftOperand, re) + " is " + operators[rule.Operator].Term + " " + toString(rule.RightOperand, re)
	}

	re.stack = append(re.stack, RuleResult{
		Operator:    rule.Operator,
		Outcome:     result,
		Explanation: explanation,
	})

	return result
}

// func (rule *Rule) Explain(re *RulesEngine) string {
// 	operators := re.Operators()
// 	operator := operators[rule.Operator]

// }
