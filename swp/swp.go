package swp

func EncryptA(ka []byte, kb []byte, s string) ([][]byte, error) {
	var w []string = cut(s)
	// Plain words. W1, W2, ..., Wn
	
	var c [][]byte = make([][]byte, len(w))
	for i, v := range w {
		ciferBytes, err := encryptAESECB([]byte(v), ka)
		if err != nil {
			return nil, err
		}
		cifer := string(ciferBytes)
		// Block encryption. E(ka', Wi)
		
		var l, r string = cifer[:len(cifer)/2], cifer[len(cifer)/2:]
		// Split equal to Li, Ri
		
		k := random1(string(kb), l)
		// Ki = f(K'', Li)
		
		a := random3(v, k)
		// F(Ki,Li)
		
		vBytes := []byte(r)
		for ii, vv := range vBytes {
			vBytes[ii] = vv ^ byte(a)
		}
		// XOR(F(Ki, Li), Ri)
		
		c[i] = vBytes
	}
	return c, nil
}

func Search(k string, w string) bool {

}
