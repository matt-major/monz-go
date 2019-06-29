package monzgo

// AddFeedItem adds an item to the Feed for the specified Account ID
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

	return m.request("POST", "feed", emptyRespHolder, requestData)
}
