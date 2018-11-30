package tasks

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
	var tasks []Task
	for _, task := range repository.db {
		tasks = append(tasks, task)
	}
	return tasks
}

func (repository memoryMapRepository) get(id int16) Task {
	return repository.db[id]
}

func (repository memoryMapRepository) update(ask *Task) {
	repository.db[ask.Id] = *ask
}

func (repository memoryMapRepository) delete(id int16) {
	delete(repository.db, id)
}

func (repository *memoryMapRepository) create(task *Task) {
	task.Id = repository.idCounter
	repository.db[task.Id] = *task
	repository.idCounter += 1
}

func NewRepository(repositoryType RepositoryType) Repository {
	switch repositoryType {
	case MemoryMap:
		return &memoryMapRepository{db: make(map[int16]Task), idCounter: 1}
	default:
		return nil
	}
}
