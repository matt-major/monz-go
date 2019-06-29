package monzgo

type FeedItemOptions struct {
	Type            string
	URL             string
	Title           string
	ImageURL        string
	BackgroundColor string
	BodyColor       string
	TitleColor      string
	Body            string
}

func (m *Monzgo) AddFeedItem(accountID string, options FeedItemOptions) error {
	requestData := make(map[string]string)
	requestData["account_id"] = accountID
	requestData["type"] = options.Type
	requestData["url"] = options.URL
	requestData["params[title]"] = options.Title
	requestData["params[image_url]"] = options.ImageURL
	requestData["params[background_color]"] = options.BackgroundColor
	requestData["params[body_color]"] = options.BodyColor
	requestData["params[title_color]"] = options.TitleColor
	requestData["params[body]"] = options.Body

	emptyRespHolder := &map[string]string{}

	err := m.request("POST", "feed", emptyRespHolder, requestData)
	return err
}
