package mid

type ThroneInheritance struct {
	childs   map[string][]string
	deaths   map[string]bool
	kingName string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		childs:   make(map[string][]string),
		deaths:   make(map[string]bool),
		kingName: kingName,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.childs[parentName] = append(this.childs[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.deaths[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string

	var trace func(king string)
	trace = func(king string) {
		if _, ok := this.deaths[king]; !ok {
			ans = append(ans, king)
		}
		for _, child := range this.childs[king] {
			trace(child)
		}
	}

	trace(this.kingName)
	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
