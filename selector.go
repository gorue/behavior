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
	if result.Complete {
		return Result{Complete: true}, d
	}
	if result.Failed {
		d.choice++
		if d.choice >= len(s.Choices) {
			return Result{Failed: true}, d
		}
	}
	return Result{}, d
}

func Selector(Choices ...Behavior) Behavior {
	return &selector{
		Choices: Choices,
	}
}
