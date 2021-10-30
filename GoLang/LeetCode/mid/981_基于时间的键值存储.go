package mid

//type pair struct {
//	timestamp int
//	value     string
//}
//
//type TimeMap struct {
//	m map[string][]pair
//}
//
//func Constructor() TimeMap {
//	return TimeMap{m: map[string][]pair{}}
//}
//
//func (this *TimeMap) Set(key string, value string, timestamp int) {
//	this.m[key] = append(this.m[key], pair{
//		timestamp: timestamp,
//		value:     value,
//	})
//}
//
//func (this *TimeMap) Get(key string, timestamp int) string {
//	pairs := this.m[key]
//	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
//	if i > 0 {
//		return pairs[i].value
//	}
//	return ""
//}
