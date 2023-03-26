package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brianstrauch/spotify"
	"github.com/kyokomi/emoji/v2"
	"github.com/olekukonko/tablewriter"
)

// playback

type Line struct {
	Emoji  string
	Output string
}

func (e Line) String() string {
	return fmt.Sprintf("   %s\r%s\n", e.Emoji, e.Output)
}

func status(api *spotify.API) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New("No Device Found")
	}

	return Show(playback), nil
}

func Show(playback *spotify.Playback) string {
	var artistLine string
	switch playback.Item.Type {
	case "track":
		artistLine = JoinArtists(playback.Item.Artists)
	case "episode":
		artistLine = playback.Item.Show.Name
	}

	var isPlayingEmoji string
	if playback.IsPlaying {
		isPlayingEmoji = emoji.Sprint(":play-button: ")
	} else {
		isPlayingEmoji = emoji.Sprint(":pause-button: ")
	}

	progressBar := getProgressBar(playback.ProgressMs, playback.Item.Duration)

	output := Line{Emoji: "🎵 ", Output: playback.Item.Name}.String()
	output += Line{Emoji: "👤 ", Output: artistLine}.String()
	output += Line{Emoji: isPlayingEmoji, Output: progressBar}.String()
	fmt.Printf("▶️")
	return output
}

func JoinArtists(artists []spotify.Artist) string {
	list := artists[0].Name
	for i := 1; i < len(artists); i++ {
		list += ", " + artists[i].Name
	}
	return list
}

func getProgressBar(progress int, duration *spotify.Duration) string {
	const length = 16
	bars := length * progress / int(duration.Milliseconds())

	status := fmt.Sprintf("%s [", formatTime(progress))
	for i := 0; i < bars; i++ {
		status += "="
	}
	for i := bars; i < length; i++ {
		status += " "
	}
	status += fmt.Sprintf("] %s", formatTime(int(duration.Milliseconds())))

	return status
}

func formatTime(ms int) string {
	s := ms / 1000
	m := s / 60
	h := m / 60

	if h == 0 {
		return fmt.Sprintf("%d:%02d", m, s%60)
	} else {
		return fmt.Sprintf("%d:%02d:%02d", h, m%60, s%60)
	}
}

// playlist

func DisplayPlaylist(playlist *spotify.Playlist) (string, error) {
	output := new(strings.Builder)

	table := tablewriter.NewWriter(output)

	table.SetHeader([]string{"Track", "Artist(s)"})

	for _, playlistTrack := range playlist.Tracks.Items {
		track := playlistTrack.Track

		artists := make([]string, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j] = artist.Name
		}

		table.Append([]string{track.Name, strings.Join(artists, ", ")})
	}
	table.Render()

	return output.String(), nil
}

// album

func DisplayAlbum(album *spotify.Album) (string, error) {

	output := new(strings.Builder)

	table := tablewriter.NewWriter(output)

	table.SetHeader([]string{"#", "Track"})

	for _, albumTrack := range album.Tracks.Items {
		track := albumTrack

		artists := make([]string, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j] = artist.Name
		}

		table.Append([]string{track.Name, strings.Join(artists, ", ")})
	}
	table.Render()

	return output.String(), nil
}
