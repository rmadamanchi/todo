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

func (r memoryMapRepository) all() []Task {
	var tasks []Task
	for _, task := range r.db {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r memoryMapRepository) get(id int16) Task {
	return r.db[id]
}

func (r memoryMapRepository) update(t *Task) {
	r.db[t.Id] = *t
}

func (r memoryMapRepository) delete(id int16) {
	delete(r.db, id)
}

func (r *memoryMapRepository) create(t *Task) {
	t.Id = r.idCounter
	r.db[t.Id] = *t
	r.idCounter += 1
}

func NewRepository(repositoryType RepositoryType) Repository {
	switch repositoryType {
	case MemoryMap:
		return &memoryMapRepository{db: make(map[int16]Task), idCounter: 1}
	default:
		return nil
	}
}
