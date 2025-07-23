package main

func intadd10(i *int) int {
	*i += 10
	return *i
}

func intmul2(i []int) []int {
	for j := range i {
		i[j] *= 2
	}
	return i
}

func main() {
	i := 10
	intadd10(&i)
	println(i)

	slice := []int{1, 2, 3, 4, 5}
	intmul2(slice)
	for _, v := range slice {
		println(v)
	}
}
