package main


import "math/big"


func lcm(m, n *big.Int) *big.Int {
	result := big.NewInt(1)
	mult := big.NewInt(1)
	gcd := big.NewInt(1)
	
	mult.Mul(m, n)
	gcd.GCD(nil, nil, m, n)
	result.Div(mult, gcd)

	return result
}
