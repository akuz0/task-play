package mp

type Playlister interface {
	Status() (bool, error)
	Clear() error
	PlaylistLoad(string, int, int) error
	Add(string) error
	Play(int) error
}
