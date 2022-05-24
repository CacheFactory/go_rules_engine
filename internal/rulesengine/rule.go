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

func (rule *Rule) Run(re *RulesEngine, doNotPushToStack bool) interface{} {
	operators := re.Operators()

	result := operators[rule.Operator].Func(rule.LeftOperand, rule.RightOperand)
	if doNotPushToStack != true {
		var explanation string

		if rule.LeftOperand != nil && rule.RightOperand != nil {
			term := " is "
			if toBool(result, re) == false {
				term = " is not "
			}

			explanation = toString(rule.LeftOperand, re, true) + term + operators[rule.Operator].Term + " " + toString(rule.RightOperand, re, true)
			//explanation = toExplanation(rule.LeftOperand, re, true) + " is " + operators[rule.Operator].Term + " " + toExplanation(rule.RightOperand, re, true)

			re.stack = append(re.stack, RuleResult{
				Operator:    rule.Operator,
				Outcome:     result,
				Explanation: explanation,
			})
		}

	}

	return result
}
