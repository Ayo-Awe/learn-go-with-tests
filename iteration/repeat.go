package iteration

const DefaultRepeatCount = 5

func Repeat(char string, count int) string {

	if count < 1 {
		count = DefaultRepeatCount
	}

	var repeated string
	for i := 0; i < count; i++ {
		repeated += char
	}

	return repeated
}
