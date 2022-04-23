package utils

type bigger interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

func Min[T bigger](elems ...T) T {
	least := elems[0]
	for _, elem := range elems {
		if elem < least {
			least = elem
		}
	}

	return least
}
