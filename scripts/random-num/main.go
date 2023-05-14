package main

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func Handle_error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(rand.Int())
	rand_num := rand.Int()
	fmt.Printf("%T %d\n", rand_num, rand_num)
	// set the seed as the current tim in nano seconds
	// this will set the initial value of the RNG so that
	// each time a new sequence is generated based on the seed value

	// deprecated
	//rand.Seed(time.Now().UnixNano())

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Int())
	rand_num = r.Int()
	fmt.Printf("%T %d\n", rand_num, rand_num)

	num := r.Int()
	fmt.Println(num)

	rand_float32 := r.Float32()
	fmt.Println(rand_float32)

	rand_float64 := r.Float64()
	fmt.Println(rand_float64)

	rand_exp_float := r.ExpFloat64()
	fmt.Println(rand_exp_float)

	rand_norm_float := r.NormFloat64()
	fmt.Println(rand_norm_float)

	for i := 0; i < 5; i++ {
		rand_float := r.Float32()
		fmt.Println(rand_float)
	}

	for i := 0; i < 10; i++ {
		// generate a integer between 0 and 5
		dice_throw := r.Intn(6)
		// Move the Offset of 0
		fmt.Println(dice_throw + 1)

	}

	// Cryptographic random numbers
	var max *big.Int = big.NewInt(6)
	// big is a package for high precision arithmetic
	for i := 0; i < 10; i++ {
		crypt_rand_num, err := crypto_rand.Int(crypto_rand.Reader, max)
		Handle_error(err)
		// Move the Offset of 0 by adding 1
		crypt_num := crypt_rand_num.Add(crypt_rand_num, big.NewInt(1))
		fmt.Println(crypt_num)
	}
	prime, err := crypto_rand.Prime(crypto_rand.Reader, 8)
	Handle_error(err)
	fmt.Println(prime)

	rand_byte := make([]byte, 5)
	fmt.Println(rand_byte)
	_, err = r.Read(rand_byte)
	Handle_error(err)
	fmt.Println(rand_byte)
	fmt.Printf("%c \n", rand_byte)

	crypt_rand_byte := make([]byte, 5)
	fmt.Println(crypt_rand_byte)
	_, err = crypto_rand.Read(crypt_rand_byte)
	Handle_error(err)
	fmt.Println(crypt_rand_byte)
	fmt.Printf("%c \n", crypt_rand_byte)

}
