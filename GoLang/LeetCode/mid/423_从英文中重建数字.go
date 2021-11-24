package mid

import "bytes"

//zero              z
//two               w
//four				u
//three             r
//eight				t
//five              f
//one               o
//six               x
//seven             v
//nine              i
func originalDigits(s string) string {
	cnt := make(map[int32]int)
	for _, c := range s {
		cnt[c] += 1
	}

	var ans [10]int
	ans[0] = cnt['z']
	ans[2] = cnt['w']
	ans[4] = cnt['u']
	ans[3] = cnt['r'] - ans[0] - ans[4]
	ans[8] = cnt['t'] - ans[2] - ans[3]
	ans[5] = cnt['f'] - ans[4]
	ans[1] = cnt['o'] - ans[0] - ans[2] - ans[4]
	ans[6] = cnt['x']
	ans[7] = cnt['v'] - ans[5]
	ans[9] = cnt['i'] - ans[5] - ans[6] - ans[8]

	str := []byte{}
	for i, c := range ans {
		if c > 0 {
			str = append(str, bytes.Repeat([]byte{byte('0' + i)}, c)...)
		}
	}
	return string(str)
}
