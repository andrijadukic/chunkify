# chunkify

A simple package for collection partitioning in go. Since go does not have generics, this package does not produce
chunks of the given collection directly, rather it assumes the collection supports random access, so it returns a
channel of type Chunk, which contains only the starting and the ending index of a chunk.

## Usage

```  go

...
chunker, err := chunkify.NewChunker(len(collection), chunkSize)
if err != nil {
    fmt.Println(err)
}
for chunk := range chunker.Chunks() {
    process(chunk)
}
...

```
