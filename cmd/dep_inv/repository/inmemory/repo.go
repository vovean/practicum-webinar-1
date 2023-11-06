package inmemory

type Repo map[int]int

func (r *Repo) Get(key int) (int, bool) {
	v, ok := (*r)[key]
	return v, ok
}

func (r *Repo) Set(key, value int) {
	(*r)[key] = value
}
