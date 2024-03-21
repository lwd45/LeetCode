package mid

//type StockPrice struct {
//	maxPrice, minPrice hp
//	priceMap           map[int]int
//	maxTime            int
//}
//
//func Constructor() StockPrice {
//	return StockPrice{priceMap: make(map[int]int)}
//}
//
//func (s *StockPrice) Update(timestamp int, price int) {
//	heap.Push(&s.maxPrice, pair{-price, timestamp})
//	heap.Push(&s.minPrice, pair{price, timestamp})
//	s.priceMap[timestamp] = price
//	if timestamp > s.maxTime {
//		s.maxTime = timestamp
//	}
//}
//
//func (s *StockPrice) Current() int {
//	return s.priceMap[s.maxTime]
//}
//
//func (s *StockPrice) Maximum() int {
//	for {
//		if p := s.maxPrice[0]; -p.price == s.priceMap[p.timestamp] {
//			return -p.price
//		}
//		heap.Pop(&s.maxPrice)
//	}
//}
//
//func (s *StockPrice) Minimum() int {
//	for {
//		if p := s.minPrice[0]; p.price == s.priceMap[p.timestamp] {
//			return p.price
//		}
//		heap.Pop(&s.minPrice)
//	}
//}
//
//type pair struct {
//	price     int
//	timestamp int
//}

//type hp []pair

//func (h *hp) Len() int           { return len(*h) }
//func (h *hp) Less(i, j int) bool { return (*h)[i].price < (*h)[j].price }
//func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
//func (h *hp) Push(x interface{}) { *h = append(*h, x.(pair)) }
//func (h *hp) Pop() interface{}   { v := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return v }

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
