/*
    Copyright (C) 2016 cacahuatl < cacahuatl at autistici dot org >
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package gen

import (
	"crypto/rsa"
	"math/big"
)

var bigOne = big.NewInt(1)

func GenerateKey(p *big.Int, q *big.Int, exp int) (*rsa.PrivateKey,error) {
	priv := new(rsa.PrivateKey)
	priv.E = exp
	n := new(big.Int).Set(bigOne)
	totient := new(big.Int).Set(bigOne)
	pminus1 := new(big.Int)
	primes := []*big.Int{p, q}
	for _, prime := range primes {
		n.Mul(n, prime)
		pminus1.Sub(prime, bigOne)
		totient.Mul(totient, pminus1)
	}
	g := new(big.Int)
	priv.D = new(big.Int)
	e := big.NewInt(int64(priv.E))
	g.GCD(priv.D, nil, e, totient)
	if g.Cmp(bigOne) == 0 {
		if priv.D.Sign() < 0 {
			priv.D.Add(priv.D, totient)
		}
	}
	priv.N = n
	priv.Primes = primes
	priv.Precompute()
	return priv,priv.Validate()
}
