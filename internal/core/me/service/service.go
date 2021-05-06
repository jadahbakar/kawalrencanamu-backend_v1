package service

// type service struct {
// 	gamesRepository ports.GamesRepository
// 	uidGen          uidgen.UIDGen
// }

func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidGen:          uidGen,
	}
}
