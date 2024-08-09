package main

import "fmt"

/* 加法哈希 */
// 对输入的每个字符的ASCII码进行相加，将得到的总和作为哈希值
func addHash(key string) int {
	var hash int64
	var modulus int64

	// 大质数 1000000007
	// 质数是为了保证；哈希值均匀分布（避免哈希冲突）
	// 原因是质数不与其他数字存在公约数，减少取模操作而产生的周期性模式
	modulus = 1_000_0000_7
	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modulus
	}

	return int(hash)
}

/* 乘法哈希 */
// 利用乘法的不相关性，妹轮乘以一个常数，将各个字符的ASCII码累积到哈希值中
func mulHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1_0000_0000_7

	for _, b := range []byte(key) {
		hash = (31*hash + int64(b)) % modulus
	}

	return int(hash)

}

/* 异或哈希 */
// 将输入数据的每个元素通过异或操作累积到一个哈希值中
func xorHash(key string) int {
	hash := 0
	modulus := 1_0000_0000_7

	for _, b := range []byte(key) {
		fmt.Println(int(b))
		hash ^= int(b)
		hash = (31*hash + int(b)) % modulus
	}
	return hash & modulus
}

/* 旋转哈希 */
// 将每个字符的ASCII码累积到一个哈希值中，每次累积之前都会对哈希值进行旋转操作
func rotHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1_0000_0000_7
	for _, b := range []byte(key) {
		hash = ((hash << 4) ^ (hash >> 28) ^ int64(b)) % modulus
	}
	return int(hash)
}

func main() {
	add := addHash("hong")
	fmt.Println("add -> ", add)

	mul := mulHash("hong")
	fmt.Println("mul -> ", mul)

	xor := xorHash("hong")
	fmt.Println("xor -> ", xor)

	rot := rotHash("hong")
	fmt.Println("rot -> ", rot)
}
