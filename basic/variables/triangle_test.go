package main

import "testing"

func TestTriangle(t *testing.T) {
	testCases := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, testCase := range testCases {
		if actual := calcTriangle(testCase.a, testCase.b); actual != testCase.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d",
				testCase.a, testCase.b, actual, testCase.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	c1, c2, h := 12000, 35000, 37000
	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(c1, c2); actual != h {
			b.Errorf("calcTriangle(%d, %d); got %d; expected %d",
				c1, c2, actual, h)
		}
	}
}
