package test

import (
	"github.com/xiuos/mozi/models"
	"testing"
)

func Benchmark_GetUserByID(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		models.GetUserByID(1)
	}
}
