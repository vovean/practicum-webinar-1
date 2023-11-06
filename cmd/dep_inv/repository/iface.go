package repository

type IRepo interface {
	Get(key int) (int, bool)
	Set(key, value int)
}
