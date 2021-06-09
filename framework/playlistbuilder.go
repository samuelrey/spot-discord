package framework

type PlaylistBuilder struct {
	MusicServiceBuilder
}

type Playlist struct {
	ID  string
	URL string
}

type MusicServiceBuilder interface {
	AuthorizeUser()
	CreatePlaylist(playlistName string) (*Playlist, error)
}
