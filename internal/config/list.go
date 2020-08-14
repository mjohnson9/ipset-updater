package config

import "net/url"

type List struct {
	Name             string   `toml:"name"`
	URL              *url.URL `toml:"url"`
	CommentCharacter string   `toml:"comment_character,omitempty"`
}
