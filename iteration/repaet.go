package iteration

const repeatCount = 5

func Repeat(word string) (ret string) {
	for i := 0; i < repeatCount; i++ {
		ret += word
	}
	return ret
}