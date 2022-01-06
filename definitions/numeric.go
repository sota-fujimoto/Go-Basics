package main

import "fmt"

/* A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are:

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

*/

func main() {
	var (
		u8 uint8 = 255
		i8 int8  = 127
	)

	fmt.Printf("type = %T value = %v\n", u8, u8)
	fmt.Printf("type = %T value = %v", i8, i8)

	// shift from Decimal to Binary

	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
}