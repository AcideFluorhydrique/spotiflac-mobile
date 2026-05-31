package gobackend

import "testing"

func TestCrossExtensionShareUsesAlbumCollectionItems(t *testing.T) {
	ext := &loadedExtension{
		Manifest: &ExtensionManifest{
			Capabilities: map[string]interface{}{
				"shareUrlTemplates": map[string]interface{}{
					"album": "https://music.apple.com/us/album/{id}",
				},
			},
		},
	}
	tracks := []ExtTrackMetadata{
		{
			ID:       "1440783617",
			Name:     "Nevermind",
			Artists:  "Nirvana",
			ItemType: "album",
		},
	}

	best := bestAlbumTrack(tracks, "Nevermind", "Nirvana")
	if best == nil {
		t.Fatal("expected album collection item to match")
	}
	if url := resolveCollectionShareURL(ext, "album", best); url != "https://music.apple.com/us/album/1440783617" {
		t.Fatalf("album share URL = %q", url)
	}
}

func TestCrossExtensionShareUsesArtistCollectionItems(t *testing.T) {
	ext := &loadedExtension{
		Manifest: &ExtensionManifest{
			Capabilities: map[string]interface{}{
				"shareUrlTemplates": map[string]interface{}{
					"artist": "https://music.youtube.com/browse/{id}",
				},
			},
		},
	}
	tracks := []ExtTrackMetadata{
		{
			ID:       "UCrPe3hLA51968GwxHSZ1llw",
			Name:     "Nirvana",
			ItemType: "artist",
		},
	}

	best := bestArtistTrack(tracks, "Nirvana")
	if best == nil {
		t.Fatal("expected artist collection item to match")
	}
	if url := resolveCollectionShareURL(ext, "artist", best); url != "https://music.youtube.com/browse/UCrPe3hLA51968GwxHSZ1llw" {
		t.Fatalf("artist share URL = %q", url)
	}
}
