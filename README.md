# Go Proverbs

In 2015, Rob Pike (one of Go's creators) gave a talk at Gopherfest SV 2015 where he outlined what have become known as "Go Proverbs", a set of principles that every Go developer should keep in mind when working with the programming language.

This package simply exposes those proverbs and their relevant locations in the talk [video](https://www.youtube.com/watch?v=PAAkCSZUG1c).

[![Go Report Card](https://goreportcard.com/badge/github.com/jboursiquot/proverbs)](https://goreportcard.com/report/github.com/jboursiquot/proverbs) [![GoDoc](https://godoc.org/github.com/jboursiquot/proverbs?status.svg)](https://godoc.org/github.com/jboursiquot/proverbs) [![CircleCI](https://circleci.com/gh/jboursiquot/proverbs.svg?style=svg)](https://circleci.com/gh/jboursiquot/proverbs)

## Usage

### Random Proverb

```go
package main

import "github.com/jboursiquot/proverbs"

func main() {
  fmt.Println(proverbs.Random())
  fmt.Println(proverbs.Random())
  fmt.Println(proverbs.Random())
}
```

#### Result
```plain
&{With the unsafe package there are no guarantees. https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s}
&{Reflection is never clear. https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s}
&{A little copying is better than a little dependency. https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s}
```

### Nth Proverb

```go
package main

import "github.com/jboursiquot/proverbs"

func main() {
  p, err := proverbs.Nth(4)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%+v\n", p)
}
```
#### Result
```plain
&{Saying:Make the zero value useful. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s}
```

### All Proverbs

```go
package main

import "github.com/jboursiquot/proverbs"

func main() {
  list, err := proverbs.All()
  if err != nil {
    fmt.Println(err)
  }
  for _, p := range list {
    fmt.Printf("%+v\n", p)
  }
}

```
#### Result
```plain
&{Saying:The bigger the interface, the weaker the abstraction. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s}
&{Saying:Syscall must always be guarded with build tags. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m10s}
&{Saying:Cgo must always be guarded with build tags. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m53s}
&{Saying:Errors are values. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=16m13s}
&{Saying:Concurrency is not parallelism. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s}
&{Saying:Channels orchestrate; mutexes serialize. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=4m20s}
&{Saying:Ggo is not Go. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=12m37s}
&{Saying:A little copying is better than a little dependency. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s}
&{Saying:Documentation is for users. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=19m07s}
&{Saying:Don't panic. Link:https://github.com/golang/go/wiki/CodeReviewComments#dont-panic}
&{Saying:Make the zero value useful. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s}
&{Saying:Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s}
&{Saying:With the unsafe package there are no guarantees. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s}
&{Saying:Clear is better than clever. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s}
&{Saying:Reflection is never clear. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s}
&{Saying:Don't just check errors, handle them gracefully. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s}
&{Saying:Design the architecture, name the components, document the details. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=18m09s}
&{Saying:Don't communicate by sharing memory, share memory by communicating. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s}
&{Saying:interface{} says nothing. Link:https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s}
```

## License
MIT