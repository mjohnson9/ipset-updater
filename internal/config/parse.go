package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Parse(r io.Reader) (*Config, error) {
	config := new(Config)

	_, err := toml.DecodeReader(r, config)

	return config, err

	// fullbogon_url, err := url.Parse("https://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// config := &Config{
	// 	CacheDir: "/Users/michael/Downloads/ipset-cache",
	// 	Lists: []List{
	// 		List{
	// 			Name:             "Fullbogons",
	// 			URL:              fullbogon_url,
	// 			CommentCharacter: "#",
	// 		},
	// 	},
	// }

	// return config, nil
}
