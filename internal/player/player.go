package player

type Player struct {
	Name     string
	SbLvl    int
	CataLvl  int
	ClassLvl int
}

func NewPlayer(name string, sbLvl, cataLvl, classLvl int) *Player {
	player := Player{
		Name:     name,
		SbLvl:    sbLvl,
		CataLvl:  cataLvl,
		ClassLvl: classLvl,
	}

	return &player
}
