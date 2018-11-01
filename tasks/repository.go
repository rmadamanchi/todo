package tasks

type Repository interface {
	all() []Task
	get(id int16) Task
	update(t Task)
	delete(id int16)
	create(t Task)
}

type memoryRepository []Task

func (r *memoryRepository) all() []Task {
	return *r
}

func (r memoryRepository) get(id int16) Task {
	return Task{}
}

func (r memoryRepository) update(t Task) {
}

func (r memoryRepository) delete(id int16) {
}

func (r memoryRepository) create(t Task) {
}

func NewRepository() Repository {
	return new(memoryRepository)
}
