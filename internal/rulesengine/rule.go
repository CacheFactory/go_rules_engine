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

func (rule *Rule) Run(re *RulesEngine) (interface{}, string) {
	operators := re.Operators()

	result := operators[rule.Operator].Func(rule.LeftOperand, rule.RightOperand)

	explanation := operators[rule.Operator].Explanation(rule.LeftOperand, rule.RightOperand, toBool(result, re))

	return result, explanation
}
