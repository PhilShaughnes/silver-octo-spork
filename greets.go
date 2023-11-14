package main

func hi(s string) string {
  greet := "Hello"
  if s != "" {
    greet += " " + s
  }
  return greet + "!"
}
