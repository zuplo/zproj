package names

import (
	"fmt"
	"math/rand"
)

var adjectives = []string{
	"bold", "calm", "cool", "dark", "deep", "dry", "fair", "fast",
	"firm", "flat", "full", "gold", "gray", "keen", "kind", "lean",
	"loud", "mild", "neat", "new", "old", "pale", "pure", "raw",
	"red", "rich", "ripe", "safe", "slim", "soft", "tall", "thin",
	"warm", "wide", "wild", "wise", "young", "blue", "green", "swift",
}

var nouns = []string{
	"ash", "bay", "birch", "brook", "cedar", "cliff", "cloud", "cove",
	"creek", "dawn", "dew", "dune", "elm", "fern", "field", "fjord",
	"fog", "frost", "glade", "grove", "hawk", "heath", "hill", "ivy",
	"lake", "leaf", "marsh", "mesa", "mist", "moon", "moss", "oak",
	"palm", "peak", "pine", "pond", "rain", "reef", "ridge", "river",
	"sage", "shade", "shore", "sky", "snow", "spruce", "star", "stone",
	"storm", "sun", "thorn", "tide", "vale", "wave", "willow", "wind",
}

// Generate returns a random name like "bold-cedar".
func Generate() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%s-%s", adj, noun)
}
