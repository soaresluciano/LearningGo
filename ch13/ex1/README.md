# Exercise 1

## Question
Create a middleware generating function that puts a timeout into the context. The function should have one parameter, which is the number of milliseconds that a request is allowed to run. It should return a `func(http.Handler) http.Handler`.