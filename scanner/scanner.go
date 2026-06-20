package scanner

type Scanner interface {
	Scan(target string) (interface{}, error)
}
