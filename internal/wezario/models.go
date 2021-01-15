package wezario

type weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}

type result struct {
	Weather []weather `json:"weather"`
	Main    main      `json:"main"`
}
