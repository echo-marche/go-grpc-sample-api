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

func InitDepositApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("DEPOSIT_API_HOST") + ":" + GetEnv("DEPOSIT_API_PORT"),
	}
}

func InitNotificationApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("NOTIFICATION_API_HOST") + ":" + GetEnv("NOTIFICATION_API_PORT"),
	}
}

func InitPopularityApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("POPULARITY_API_HOST") + ":" + GetEnv("POPULARITY_API_PORT"),
	}
}

func InitInquiryApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("INQUIRY_API_HOST") + ":" + GetEnv("INQUIRY_API_PORT"),
	}
}

func InitSendmailApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("SENDMAIL_API_HOST") + ":" + GetEnv("SENDMAIL_API_PORT"),
	}
}
