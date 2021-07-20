package main

func main() {
	println(isPalindrome(123321))
}

func isPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	tmp := 0
	for x > tmp {
		tmp = tmp*10 + x%10
		x /= 10
	}

	return tmp/10 == x || tmp == x

}
