func HandleScrapeArticleTask(db *gorm.DB) asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var p ScrapePayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("failed to parse payload: %w", err)
		}

		var article models.Article
		if err := db.First(&article, p.ArticleID).Error; err != nil {
			return fmt.Errorf("article not found: %w", err)
		}

		resp, err := http.Get(article.URL)
		if err != nil {
			return fmt.Errorf("failed to fetch article URL: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad response status: %d", resp.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to parse HTML: %w", err)
		}

		var content string
		doc.Find("p").Each(func(i int, s *goquery.Selection) {
			content += s.Text() + "\n"
		})

		article.Title = doc.Find("title").Text()
		article.Content = content
		article.Status = "scraped"

		return db.Save(&article).Error
	}
}

