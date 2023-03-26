package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/brianstrauch/spotify"
)

func Play(api *spotify.API, name string, mode int) (string, error) {

	playback, err := api.GetPlayback()
	if err != nil {
		fmt.Println("error->", playback)
		return "", err
	}

	if playback == nil {
		return "", errors.New("No Device Found")
	}

	isPlaying := playback.IsPlaying
	id := playback.Item.ID
	progressMs := playback.ProgressMs
	if mode == 1 {

		err := playTrack(name, api)
		if err != nil {
			return "", err
		}
	} else {
		err := playAlbum(name, api)
		if err != nil {
			return "", err
		}
	}
	playback, err = WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		hasChanged := playback.Item.ID != "" && (playback.Item.ID != id || playback.ProgressMs < progressMs)
		return !isPlaying && playback.IsPlaying || hasChanged
	})
	if err != nil {
		return "", err
	}

	if mode == 1 {

		return Show(playback), err

	}
	return "", nil
}

func playTrack(name string, api *spotify.API) error {
	track, err := SearchTrack(api, name, true)
	if err != nil {
		return err
	}

	if err := api.Play("", track.URI); err != nil {
		return err
	}
	return nil
}

func playAlbum(name string, api *spotify.API) error {
	album, err := SearchAlbum(api, name)
	if err != nil {
		return err
	}

	if err := api.Play(album.URI); err != nil {
		return err
	}
	res, err := DisplayAlbum(album)

	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

func WaitForUpdatedPlayback(api *spotify.API, isUpdated func(*spotify.Playback) bool) (*spotify.Playback, error) {
	timeout := time.After(time.Second)
	tick := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-timeout:
			return nil, errors.New("Request Timeout")
		case <-tick.C:
			playback, err := api.GetPlayback()
			if err != nil {
				return nil, err
			}

			if isUpdated(playback) {
				return playback, nil
			}
		}
	}
}
