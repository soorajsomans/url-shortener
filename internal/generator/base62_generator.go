package generator

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type Base62Generator struct{}

func NewBase62Generator() *Base62Generator {
	return &Base62Generator{}
}

func (g *Base62Generator) Generate(id int64) string {
	if id == 0 {
		return "0"
	}

	var encoded []byte

	for id > 0 {
		reminder := id % 62

		encoded = append(encoded, charset[reminder])

		id /= 62
	}

	reverse(encoded)
	return string(encoded)
}

func reverse(data []byte) {
	for left, right := 0, len(data)-1; left < right; {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}

// compile time check

var _ CodeGenerator = (*Base62Generator)(nil)
