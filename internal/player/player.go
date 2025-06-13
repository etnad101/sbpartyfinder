package player

type CataClass int

const (
	Healer CataClass = iota
	Mage
	Berserk
	Archer
	Tank
)

type Levels struct {
	SbLvl        int
	CataLvl      int
	CataClassLvl int
}

type Player struct {
	Name          string
	Level         Levels
	SelectedClass CataClass
}

func NewPlayer(name string, sbLvl, cataLvl, classLvl int, cataClass CataClass) *Player {
	levels := Levels{SbLvl: sbLvl, CataLvl: cataLvl, CataClassLvl: classLvl}
	return &Player{
		Name:          name,
		Level:         levels,
		SelectedClass: cataClass,
	}
}
