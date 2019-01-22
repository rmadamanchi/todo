package tasks

var repository Repository

func init() {
	repository, _ = NewRepository(MemoryMap)
	repository.create(&Task{Title: "Get Milk"})
	repository.create(&Task{Title: "Get Bread"})
}
