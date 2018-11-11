package tasks

type RepositoryType int

const (
	MemoryMap RepositoryType = iota
)

type Repository interface {
	all() []Task
	get(id int16) Task
	update(t Task)
	delete(id int16)
	create(t Task)
}

var counter int16 = 1

type memoryMapRepository map[int16]Task

func (r *memoryMapRepository) all() []Task {
	var tasks = make([]Task, len(*r))
	for _, task := range *r {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r memoryMapRepository) get(id int16) Task {
	return r[id]
}

func (r memoryMapRepository) update(t Task) {
	r[t.Id] = t
}

func (r memoryMapRepository) delete(id int16) {
	delete(r, id)
}

func (r memoryMapRepository) create(t Task) {
	t.Id = counter
	r[t.Id] = t
	counter += 1
}

func NewRepository(repositoryType RepositoryType) Repository {
	switch repositoryType {
	case MemoryMap:
		return new(memoryMapRepository)
	default:
		return nil
	}
}
