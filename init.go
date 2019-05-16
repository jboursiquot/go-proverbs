package proverbs

import (
	"encoding/csv"
	"strings"
)

const data = `"Don't communicate by sharing memory, share memory by communicating.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s"
"Concurrency is not parallelism.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=3m42s"
"Channels orchestrate; mutexes serialize.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=4m20s"
"The bigger the interface, the weaker the abstraction.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s"
"Make the zero value useful.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=6m25s"
"interface{} says nothing.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s"
"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=8m43s"
"A little copying is better than a little dependency.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s"
"Syscall must always be guarded with build tags.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m10s"
"Cgo must always be guarded with build tags.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=11m53s"
"Ggo is not Go.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=12m37s"
"With the unsafe package there are no guarantees.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=13m49s"
"Clear is better than clever.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=14m35s"
"Reflection is never clear.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=15m22s"
"Errors are values.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=16m13s"
"Don't just check errors, handle them gracefully.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=17m25s"
"Design the architecture, name the components, document the details.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=18m09s"
"Documentation is for users.","https://www.youtube.com/watch?v=PAAkCSZUG1c&t=19m07s"
"Don't panic.","https://github.com/golang/go/wiki/CodeReviewComments#dont-panic"
`

// init initializes the package-global dictionary of proverbs.
func init() {
	r := csv.NewReader(strings.NewReader(data))
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true
	records, _ := r.ReadAll() // safely ignoring error here

	dict = make(map[int]*Proverb, len(records))
	for i, r := range records {
		dict[i] = &Proverb{Saying: r[0], Link: r[1]}
	}
}
