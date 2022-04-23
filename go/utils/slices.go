package utils

func Chunk[T any](source []T, size int) [][]T {
	out := [][]T{}

	for i := 0; i < len(source); i += size {
		sizeLeft := len(source) - i
		chunkSize := Min(sizeLeft, size)

		out = append(out, source[i:i+chunkSize])
	}

	return out
}
