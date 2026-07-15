package types

type SubItem struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// NavGroup models a parent drop-down category block
type NavGroup struct {
	Service     string    `json:"service"` // Acts as the display name (e.g. "Retail & CPG")
	Link        string    `json:"link"`
	SubServices []SubItem `json:"subservices"`
}

// HeaderConfig groups everything nicely so you can pass it to your HTMX component
type HeaderConfig struct {
	Services   []NavGroup `json:"services"`
	Industries []NavGroup `json:"industries"`
	Insights   []NavGroup `json:"insights"`
	About      []NavGroup `json:"about"`
}

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVideo MediaType = "video"
)

type HeroData struct {
	Title       string
	Description string
	ButtonText  string
	ButtonURL   string
	MediaURL    string
	MediaType   MediaType
}

type AdvantageData struct {
	Title   string
	Content string
}
type TestimonialData struct {
	Title string
	Image string
}
