package slg

type Player struct {
	ID int64
}

type PlayerMgr struct {
	players map[int64]*Player
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players: make(map[int64]*Player),
	}
}

func (p *PlayerMgr) AddPlayer(id int64) {
	p.players[id] = &Player{ID: id}
}
