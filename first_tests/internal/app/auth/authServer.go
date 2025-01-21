package authserver

type Authserver struct {
	log SuperLog
}

func (s *Authserver) Start() {

}

type SuperLog interface {
	Info(msg string, args ...any)
}

type Saver interface {
	Save(msg string)
}

type Getter interface {
	Get(msg string)
}

type SeverStruct struct {
}

func (s *SeverStruct) Save(msg string) {

}

func NewServer(log SuperLog, saver Saver) *Authserver {

	return &Authserver{}
}
