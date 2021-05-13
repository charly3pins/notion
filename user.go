package notion

// User represents a user in a Notion workspace.
// Users include guests, full workspace members, and bots.
type User struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	People
	Bot
}

// People contains the information below when User that represents people have the type property set to "person".
type People struct {
	Person      interface{} `json:"person,omitempty"`
	PersonEmail string      `json:"person.email,omitempty"`
}

// Bot contains the information below when User that represents bots have the type property set to "bot".
type Bot struct {
	Bot interface{} `json:"bot,omitempty"`
}
