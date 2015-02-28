package primes

import (
    "container/list"
    "math"
)

// Simple function for checking if a value coming in on a channel is divisible by a 
// divisor, if it is then it is passed to the output channel
func checkDivisor(divisor int64, in chan int64, out chan int64) {
    for {
        val := <-in
        if val % divisor != 0 {
            out <- val
        }
    }
}

// A simple sieve implementation
func PrimeSieveOfErat(end int64) []int64 {
    if end < int64(2) {
        return make([]int64, 0)
    }
    
    ar := make([]bool, end)
    ar[0] = true
    ar[1] = true
    
    limit := int64(math.Sqrt(float64(end))) + int64(1)
    
    for i := int64(2); i < limit; i++ {
        if ar[i] == false {
            for j := i * 2; j < end; j += i {
                ar[j] = true
            }
        }
    }
    
    count := 0
    for _, v := range ar {
        if v == false {
            count++
        }
    }
    
    result := make([]int64, count)
    index := int64(0)
    for i, v := range ar {
        if v == false {
            result[index] = int64(i)
            index++
        }
    }
    
    return result
}

// Determnes if a number is prime by check to see if it is evenly divisible
// by a supplied list of prime numbers
func isPrime(n int64, primes *list.List) bool {
    if primes == nil {
        return false
    }
    
    for e := primes.Front(); e != nil; e = e.Next() {
        if n % e.Value.(int64) == 0 {
            return false
        }
    }
    
    return true
}

// Gets the Nth prime number
func GetNthPrime(n int64) int64 {
    primes := list.New()
    primes.PushBack(int64(2))
    
    for i := int64(3); int64(primes.Len()) < n; i += int64(2) {
        if isPrime(i, primes) {
            primes.PushBack(i)
        }
    }
    
    return primes.Back().Value.(int64)
}

// Creates a prime sieve
func MakePrimeSieve() chan int64 {
    // Initial seed channel
    seed := make(chan int64)

    // Channel to return to the caller
    primeCh := make(chan int64, 10)

    // Add "2" to the return channel as we know this is a prime and the only even one
    primeCh <- 2

    // Create a goroutine to add values to the seed channel
    go func() {
        for i := int64(3); ; i += 2 {
            seed <- i
        }
    }()

    // Create the output channel and create the first divisor goroutine to check for division by 2
    output := make(chan int64)
    go checkDivisor(2, seed, output)

    // Create a new go routine for processing prime numbers.  When a new value is detected on the output
    // channel then it has passed divisor checks and must be a prime number, so a new goroutine is created 
    // whose input is the output of the last goroutine.  And so each new detected prime results in a new 
    // divisor check
    go func() {
        for {
            prime := <-output
            primeCh <- prime

            input := output
            output = make(chan int64)

            go checkDivisor(prime, input, output)
        }
    }()

    return primeCh
}

// Determines the prime factors for a given value, returning a map of prime values and the number of
// times each is used
func PrimeFactors(value int64) map[int64]int64 {
    ps := MakePrimeSieve()
    prime := <-ps

    result := make(map[int64]int64)

    for value != 1 {
        for value % prime == 0 {
            if _, ok := result[prime]; ok == false {
                result[prime] = 1
            } else {
                result[prime] += 1
            }
            value /= prime
        }
        prime = <-ps
    }

    return result
}