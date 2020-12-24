package chunkify

import (
	"reflect"
	"testing"
)

func TestChunks(t *testing.T) {
	var cases = []struct {
		collectionSize int
		chunkSize      int
		expectedChunks []Chunk
	}{
		{6, 2, []Chunk{{0, 2}, {2, 4}, {4, 6}}},
		{7, 2, []Chunk{{0, 2}, {2, 4}, {4, 6}, {6, 7}}},
		{6, 10, []Chunk{{0, 6}}},
	}

	for _, testCase := range cases {
		chunker, err := NewChunker(testCase.collectionSize, testCase.chunkSize)
		if err != nil {
			t.Error("unexpected error on chunker creation", err)
		}

		var actualChunks []Chunk
		for chunk := range chunker.Chunks() {
			actualChunks = append(actualChunks, chunk)
		}

		if !reflect.DeepEqual(testCase.expectedChunks, actualChunks) {
			t.Errorf("expected %d, actual %d", testCase.expectedChunks, actualChunks)
		}
	}
}

func BenchmarkLargeCollection(b *testing.B) {
	chunker, err := NewChunker(1e8, 50)
	if err != nil {
		b.Errorf(err.Error())
	}
	benchmarkChunker(chunker, b)
}

func BenchmarkMediumCollection(b *testing.B) {
	chunker, err := NewChunker(1e4, 50)
	if err != nil {
		b.Errorf(err.Error())
	}
	benchmarkChunker(chunker, b)
}

func BenchmarkSmallCollection(b *testing.B) {
	chunker, err := NewChunker(1e2, 50)
	if err != nil {
		b.Errorf(err.Error())
	}
	benchmarkChunker(chunker, b)
}

func benchmarkChunker(chunker Chunker, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _ = range chunker.Chunks() {
		}
	}
}
