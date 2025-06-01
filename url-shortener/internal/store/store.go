package store

type URLStore interface {
	Save(url string) (string, error)
	Get(code string) (string, error)
}
