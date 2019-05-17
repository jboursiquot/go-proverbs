# Go Proverbs

In 2015, Rob Pike (one of Go's creators) gave a talk at Gopherfest SV 2015 where he outlined what have become known as "Go Proverbs", a set of principles that every Go developer should keep in mind when working with the programming language.

This package simply exposes those proverbs and their relevant locations in the talk [video](https://www.youtube.com/watch?v=PAAkCSZUG1c).

[![Go Report Card](https://goreportcard.com/badge/github.com/jboursiquot/proverbs)](https://goreportcard.com/report/github.com/jboursiquot/proverbs)

## Usage Example

```go
package main

import "github.com/jboursiquot/proverbs"

func main() {
  fmt.Printf("%+v\n", proverbs.Random())
}
```

```
&{Saying:A little copying is better than a little dependency. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s}
```
