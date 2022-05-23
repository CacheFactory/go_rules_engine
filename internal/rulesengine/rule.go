package rulesengine

type Rule struct {
	Operator     string `json:"operator,omitempty"`
	LeftOperand  *Rule  `json:"left_operand,omitempty"`
	RightOperand *Rule  `json:"right_operand,omitempty"`
	Explanation  string `json:"explanation,omitempty"`
	Value        string `json:"value,omitempty"`
}

type RuleResult struct {
	Outcome  interface{}
	Left     string
	Right    string
	Operator string
}

func (rule *Rule) Run(re *RulesEngine) interface{} {
	operators := re.Operators()

	var left string
	var right string

	// if rule.LeftOperand != nil {
	// 	left = toString(rule.LeftOperand, re)
	// }

	// if rule.RightOperand != nil {
	// 	right = toString(rule.RightOperand, re)
	// }

	re.stack = append(re.stack, RuleResult{
		Operator: rule.Operator,
		Left:     left,
		Right:    right,
	})

	return operators[rule.Operator](rule.LeftOperand, rule.RightOperand)
}
