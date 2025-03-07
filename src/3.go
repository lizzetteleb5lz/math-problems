package main

import "math/rand"

func randomProblem() int {
  // Generate a random number between 1 and 10
  num := rand.Intn(10) + 1

  // Generate a random operation (+, -, *, /)
  op := []string{"+", "-", "*", "/"}[rand.Intn(4)]

  // Generate a random number to solve the problem
  sol := rand.Intn(num) + 1

  return num, op, sol
}
