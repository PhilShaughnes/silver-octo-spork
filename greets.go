package main

func hi(s string) string {
  greet := "Hello"
  if s != "" {
    greet += " " + s
  }
  return greet + "!"
}

func bye(s string) string {
  farewell := "Goodbye"
  return farewell + " " + s
}
