# Project Euler in Go

Golang solutions for [Project Euler](https://projecteuler.net/) problems that
(importantly) don't give away the answers for free.

## Build

```
go fmt ./... && go test ./... && go install
```

## Examples of running the program

```
% euler solve -args '[[3, 5], 10]'

If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the
multiples of 3 or 5 below 1000.

"[[3, 5], 10]" -> 23
```

```
% euler solve -args '[[2489, 444, 34], 34834834]'

If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the
multiples of 3 or 5 below 1000.

"[[2489, 444, 34], 34834834]" -> 19367245639428
```

```
% euler solve -args '[[3, 5], 1000]'
panic
```
