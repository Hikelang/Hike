# Package for generating documentation, via docstring format in Hike:

For example, this is a simple documentation for a function:

```hike
// @brief sum function
// @description function that sums two numbers
// @param a first number
// @param b second number
// @return sum of two numbers
sum = <T>(a T, b T) int {
  return int(a) + int(b);
}
```