# My Answers to the Official Go Tour

Found at: https://go.dev/tour/
My first time learning and using Go




# Problem Descriptions

## Exercise Loops and Functions
Implement a square root function using Newton's method. 
Function returns when the error is less than or equal to `epsilon` or `maxitr` number of max iterations are reached.


## Exercise Slices
Pic should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it output base 64 representing your picture, interpreting the integers as grayscale values. Otherwise if you're using the Go Tour page it will display your picture in bluescale. The choice of image generation fn is free.


## Exercise Maps
Implement WordCount using stirngs.Fields. It returns a map of the counts of each "word" (space separated) in the string `s`. Uses the golang tour `wc.Test` function to run a test suite against the provided function and prints success or failure.


## Exercise Fibonacci Closure
The fibonacci function returns a function that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...) using a closure. 


## Exercise Stringers
Implement `fmt.Stringer` on the `IPAddr` type to print the address as a dotted quad. `IPAddr{1, 2, 3, 4}` should print as "1.2.3.4".


## Exercise Errors
Copy the `Sqrt` function from the earlier "Loops and Functions" exercise. It is modified to return an error value when given a negative number. Implements the `Error()` function.


## Exercise Readers
Implemented a Reader type that emits an infinite stream of ASCII "A".


## Exercise rot13Reader
Implement the `io.Reader` interface on the `rot13Reader` type by implementing the `Read` method. The `rot13Reader` implements `io.Reader` and reads from an `io.Reader`, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.


## Exercise Images
Rewrite the earlier picture generator in the "Slices" exercise to implement image.Image instead of a slice. A custom Image type is defined and the `ColorModel`, `Bounds` and `At` functions are implemented to be able to call pic.ShowImage.


## Exercise Equivalent Binary Trees
Using the golang tour tree package, implement and test the `Walk` function. Implement the `Same` function and test it.


## Exercise Web Crawler
Used Go's concurrency features to parallelise a web crawler. Modified the `Crawl` function to fetch URLs in parallel without fetching the same URL twice. Crawl uses `fetcher` to recursively crawl pages starting with `url`, to a maximum of `depth`.