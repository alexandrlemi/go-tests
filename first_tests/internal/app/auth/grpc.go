package authserver

// TODO: транспортый слой
type GRPCService interface {
	Start(address string, port string)
}
