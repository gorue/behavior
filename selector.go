package behavior

type selector struct {
	Choices []Behavior
}

type selectorState struct {
	choice int
}

func (s *selector) Init() BehaviorData {
	bd := &selectorState{}
	return bd
}

func (s *selector) Step(r *Runner, bd BehaviorData) (Result, BehaviorData) {
	d := bd.(*selectorState)
	result := r.Next(s.Choices[d.choice])
	if result == SUCCESS {
		return SUCCESS, d
	}
	if result == FAILURE {
		d.choice++
		if d.choice >= len(s.Choices) {
			return FAILURE, d
		}
	}
	return RUNNING, d
}

func Selector(Choices ...Behavior) Behavior {
	return &selector{
		Choices: Choices,
	}
}
