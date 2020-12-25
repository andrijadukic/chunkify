package chunkify

import (
	"errors"
)

// Represents a type which returns a receive-only channel of Chunk instances.
type Chunker interface {
	Chunks() <-chan Chunk
}

// Specifies the range of indices the Chunk consists of.
// Start holds the starting index of this Chunk in the collection.
// End holds the ending index of this Chunk in the collection.
// Start is inclusive, End exclusive.
type Chunk struct {
	Start int
	End   int
}

// Concrete implementation of the Chunker interface.
// Holds the size of the collection and the desired chunk size.
type chunker struct {
	collectionSize int
	chunkSize      int
}

// Constructor of a Chunker instance.
// Returns an error if given chunkSize or collectionSIze argument is not a positive integer.
func NewChunker(collectionSize, chunkSize int) (Chunker, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunk size must be a positive integer")
	}
	if collectionSize <= 0 {
		return nil, errors.New("collection size must be a positive integer")
	}
	return &chunker{collectionSize: collectionSize, chunkSize: chunkSize}, nil
}

// Returns a buffered channel of chunks of an indexable collection.
// Consecutive chunks are sent to the channel, each of which is the same size.
// The final chunk may be smaller than the others which also means that if the given chunk size is greater than
// the size of the collection, the chunk returned includes the entire collection.
// The returned channel is a buffered receive-only channel. The buffer size is set to the total number of chunks that
// will be created so as to prevent unnecessary blocking.
func (c *chunker) Chunks() <-chan Chunk {
	fullChunks := c.collectionSize / c.chunkSize
	totalChunks := fullChunks
	if c.collectionSize%c.chunkSize != 0 {
		totalChunks++
	}

	chunks := make(chan Chunk, totalChunks)
	go func() {
		defer close(chunks)

		var i int
		for ; i < fullChunks; i++ {
			chunks <- Chunk{Start: i * c.chunkSize, End: (i + 1) * c.chunkSize}
		}
		if totalChunks != fullChunks {
			chunks <- Chunk{Start: i * c.chunkSize, End: c.collectionSize}
		}
	}()

	return chunks
}
