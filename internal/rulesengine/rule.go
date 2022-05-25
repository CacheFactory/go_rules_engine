package rulesengine

type Rule struct {
	Operator     string `json:"operator,omitempty"`
	LeftOperand  *Rule  `json:"left_operand,omitempty"`
	RightOperand *Rule  `json:"right_operand,omitempty"`
	Result       string `json:"result,omitempty"`
	Value        string `json:"value,omitempty"`
	Explanation  string `json:"explanation,omitempty"`
}

func (rule *Rule) Run(re *RulesEngine) (interface{}, string) {
	operators := re.Operators()

	result := operators[rule.Operator].Func(rule.LeftOperand, rule.RightOperand)
	rule.Result = toString(result, re)
	explanation := operators[rule.Operator].Explanation(rule.LeftOperand, rule.RightOperand, toBool(result, re))
	rule.Explanation = explanation
	return result, explanation
}

func (rule *Rule) Explain(re *RulesEngine) string {
	result := re.Operators()[rule.Operator].Func(rule.LeftOperand, rule.RightOperand)
	return re.Operators()[rule.Operator].Explanation(rule.LeftOperand, rule.RightOperand, toBool(result, re))
}
