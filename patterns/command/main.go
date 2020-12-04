package main

import "github.com/rlgino/hundred-days-of-code/patterns/command/task"

func main() {
	ag1 := task.New(task.Clear, nil)
	ag2 := task.New(task.Clear, ag1)
	ag3 := task.New(task.Wash, ag2)
	ag4 := task.New(task.Dry, ag3)
	ag4.Execute()
}
