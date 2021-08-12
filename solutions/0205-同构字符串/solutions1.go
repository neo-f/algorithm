package _205_同构字符串

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		si, ti := s[i], t[i]
		// 如果存在s->t映射 且值不一致 则return
		if t, ok := s2t[si]; ok && t != ti {
			return false
		}
		// 如果存在t->s映射 且值不一致 则return
		if s, ok := t2s[ti]; ok && s != si {
			return false
		}
		s2t[si] = ti
		t2s[ti] = si
	}
	return true
}
