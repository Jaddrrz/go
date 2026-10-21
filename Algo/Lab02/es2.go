package main

import "fmt"
import "os"

/*
f(A, 1)
  if i == n
    return A[n]
  else 
    maxSub = f(A, i + 1)
    return max{A[i], maxSub]
*/ 
