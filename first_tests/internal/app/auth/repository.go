package authserver

// TODO: работа с базой данных
type Setter interface {
	Save(msg string)
}

type Getter interface {
	Get(msg string)
}

type UserRepository interface {
	Save(key string, msg string)
	Get(key string) string
}

func NewRepoMock() UserRepository {
	return &repo{Stor: map[string]string{}}
}

type repo struct {
	Stor map[string]string
}

func (r *repo) Save(key string, msg string) {
	r.Stor[key] = msg
}
func (r *repo) Get(key string) string {
	return r.Stor[key]
}
