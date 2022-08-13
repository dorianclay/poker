package codegenerator

import (
	"math/rand"
	"time"
)

// The first word: n=54
var left = [...]string{
	"baggy", "beady", "boozy", "bumpy", "chewy", "dorky", "flaky", "raspy", "shiny",
	"foggy", "fuzzy", "geeky", "gimpy", "goopy", "gummy", "hasty", "gusty", "dusty",
	"hilly", "homey", "jumpy", "lanky", "leaky", "spiky", "lumpy", "empty", "jazzy",
	"messy", "muggy", "muzzy", "nerdy", "nippy", "pasty", "pokey", "randy", "spicy",
	"ready", "scaly", "seedy", "shaky", "silly", "slimy", "sunny", "surly", "misty",
	"tacky", "tasty", "whiny", "wiggy", "wimpy", "woozy", "zippy", "nifty", "windy",
}

// The second word: n=40
var center = [...]string{
	"amber", "azure", "beige", "black", "brass", "brown", "tense", "alert",
	"coral", "cream", "denim", "green", "ivory", "khaki", "lemon", "lilac",
	"linen", "mauve", "ochre", "olive", "peach", "sepia", "smalt", "taupe",
	"wheat", "white", "civil", "clear", "fixed", "kaput", "elfin", "bored",
	"elite", "rigid", "super", "royal", "slate", "naive", "level", "tough",
}

// The third word: n=72
var right = [...]string{
	"akita", "bison", "bongo", "camel", "civet", "event", "salad", "music",
	"coati", "honey", "corgi", "crane", "entry", "party", "bread", "quack",
	"dhole", "dingo", "eagle", "fossa", "frise", "gecko", "video", "heart",
	"goose", "guppy", "heron", "horse", "hound", "husky", "hyena", "topic",
	"hyrax", "indri", "koala", "lemur", "liger", "story", "hotel", "thing",
	"llama", "macaw", "molly", "moose", "mouse", "okapi", "otter", "panda",
	"quail", "quoll", "robin", "saola", "shark", "truth", "scene", "apple",
	"sheep", "shrew", "skunk", "sloth", "snail", "snake", "spitz", "fruit",
	"squid", "stoat", "tapir", "tetra", "tiger", "whale", "zebra", "zorse",
}

func Code() string {
	rand.Seed(time.Now().UnixNano())
	var l = left[rand.Intn(len(left))]
	var c = center[rand.Intn(len(center))]
	var r = right[rand.Intn(len(right))]

	return l + "-" + c + "-" + r
}
