package tasks

type Repository interface {
	all() []Task
	get(id int16) Task
	update(t Task)
	delete(id int16)
	create(t Task)
}

type memoryMapRepository map[int16]Task

func (r *memoryMapRepository) all() []Task {
	var tasks []Task
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
	r[t.Id] = t
}

func NewRepository() Repository {
	return new(memoryMapRepository)
}
