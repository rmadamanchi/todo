package tasks

import "errors"

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

type memoryMapRepository struct {
	db        map[int16]Task
	idCounter int16
}

func (repository memoryMapRepository) all() []Task {
	// zero length slice
	var tasks = make([]Task, 0)
	for _, task := range repository.db {
		tasks = append(tasks, task)
	}
	return tasks
}

func (repository memoryMapRepository) get(id int16) Task {
	return repository.db[id]
}

func (repository memoryMapRepository) update(task *Task) {
	repository.db[task.Id] = *task
}

func (repository memoryMapRepository) delete(id int16) {
	delete(repository.db, id)
}

// must use pointer receiver to not get a copy of idCounter
func (repository *memoryMapRepository) create(task *Task) {
	task.Id = repository.idCounter
	repository.db[task.Id] = *task
	repository.idCounter += 1
}

func NewRepository(repositoryType RepositoryType) (Repository, error) {
	switch repositoryType {
	case MemoryMap:
		// https://stackoverflow.com/a/40824044
		return &memoryMapRepository{db: make(map[int16]Task), idCounter: 1}, nil
	default:
		return nil, errors.New("Unknown Repository Type")
	}
}
