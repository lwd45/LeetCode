package mid

import "math"

type Pair struct {
	Day   int
	Price int
}

type StockSpanner struct {
	CurrentDay int
	Stack      []Pair
}

func Constructor() StockSpanner {
	return StockSpanner{
		CurrentDay: -1,
		Stack:      []Pair{{Day: -1, Price: math.MaxInt64}},
	}
}

func (s *StockSpanner) Next(price int) int {
	for price >= s.Stack[len(s.Stack)-1].Price {
		s.Stack = s.Stack[:len(s.Stack)-1]
	}
	s.CurrentDay++
	s.Stack = append(s.Stack, Pair{
		Day:   s.CurrentDay,
		Price: price,
	})
	return s.CurrentDay - s.Stack[len(s.Stack)-2].Day
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
