# goGCD
 Trying Out Go Generics & Fuzzing By Implementing Euclid's Algorithm 
 
Eculid's Algorithm is used to compute the greatest common divisor of two integers. This repo is an attempt to learn how to implement Generics and Fuzzing in Go by creating a `gcd` function applicable for all Integer types using Generics and fuzz testing the correctness of the implementation using the Go fuzzing feature.

When fuzzing the properties to check through were refrenced from [here](https://en.wikipedia.org/wiki/Greatest_common_divisor#Properties).

The fuzzing command is as follows

```
go test -v -fuzz=Fuzz -fuzztime 30s
```
 
