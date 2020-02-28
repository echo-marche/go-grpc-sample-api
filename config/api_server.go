package config

type ApiConfig struct {
	Url string
}

func InitPresenceApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("PRESENCE_API_HOST") + ":" + GetEnv("PRESENCE_API_PORT"),
	}
}

func InitArticleApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("ARTICLE_API_HOST") + ":" + GetEnv("ARTICLE_API_PORT"),
	}
}
