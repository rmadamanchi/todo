package tasks

import "testing"

func TestMemoryMapRepository(t *testing.T) {
	repository := NewRepository(MemoryMap)

	repository.create(&Task{Title: "Get Milk"})
	repository.create(&Task{Title: "Get Bread"})
	repository.create(&Task{Title: "Fill Gas", Done: true})

	all := repository.all()
	assert(t, len(all), 3)

	task := repository.get(1)
	assert(t, task.Title, "Get Milk")
}

func assert(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
