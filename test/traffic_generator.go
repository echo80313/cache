package test

import (
	"math/rand"
)

type OpType int

const (
	GetOp OpType = iota
	PutOp
)

const (
	keySize = 5
	valSize = 1024
)

type Op struct {
	t   OpType
	key string
	val interface{} // If PUT
}

func GenRandomOps(n int) []*Op {
	ops := make([]*Op, n)
	for i := 0; i < n; i++ {
		t := OpType(rand.Intn(2))
		if t == GetOp {
			ops[i] = &Op{t, randString(keySize), nil}
		} else {
			ops[i] = &Op{t, randString(keySize), randString(valSize)}
		}
	}
	return ops
}
