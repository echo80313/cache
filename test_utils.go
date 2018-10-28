package cache

import (
	"math/rand"
)

const charSet = "qwertyuiop[]asdfghjkl;'zxcvbnm,./`1234567890-="
const charSetSize = len(charSet)

func randChar() byte {
	return charSet[rand.Intn(charSetSize)]
}

func randString(length int) string {
	if length == 0 {
		return ""
	}
	str := ""
	for i := 0; i < length; i++ {
		str += string(randChar())
	}
	return str
}

func genKeyValueSet(sz, keySize, valueSize int) map[string]interface{} {
	kvSet := make(map[string]interface{})
	for i := 0; i < sz; i++ {
		kvSet[randString(keySize)] = randString(valueSize)
	}
	return kvSet
}
