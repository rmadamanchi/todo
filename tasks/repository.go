package tasks

import "fmt"

type RepositoryType int

const (
	MemoryMap RepositoryType = iota
)

type Repository interface {
	all() []Task
	get(id int16) Task
	update(t *Task)
	delete(id int16)
	create(t *Task)
}

var counter int16 = 1

type memoryMapRepository map[int16]Task

func (repository *memoryMapRepository) all() []Task {
	var tasks []Task
	for _, task := range *repository {
		tasks = append(tasks, task)
	}
	return tasks
}

func (repository memoryMapRepository) get(id int16) Task {
	return repository[id]
}

func (repository memoryMapRepository) update(t *Task) {
	repository[t.Id] = *t
}

func (repository memoryMapRepository) delete(id int16) {
	delete(repository, id)
}

func (repository memoryMapRepository) create(t *Task) {
	fmt.Print(repository)
	t.Id = counter
	repository[t.Id] = *t
	counter += 1
}

func NewRepository(repositoryType RepositoryType) Repository {
	switch repositoryType {
	case MemoryMap:
		return &memoryMapRepository{}
	default:
		return nil
	}
}
