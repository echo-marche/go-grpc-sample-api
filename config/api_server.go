package config

type ApiConfig struct {
	Url string
}

func InitPresenceApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("PRESENCE_API_SERVICE_HOST") + ":" + GetEnv("PRESENCE_API_SERVICE_PORT"),
	}
}

func InitArticleApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("ARTICLE_API_SERVICE_HOST") + ":" + GetEnv("ARTICLE_API_SERVICE_PORT"),
	}
}

func InitDepositApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("DEPOSIT_API_SERVICE_HOST") + ":" + GetEnv("DEPOSIT_API_SERVICE_PORT"),
	}
}

func InitNotificationApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("NOTIFICATION_API_SERVICE_HOST") + ":" + GetEnv("NOTIFICATION_API_SERVICE_PORT"),
	}
}

func InitPopularityApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("POPULARITY_API_SERVICE_HOST") + ":" + GetEnv("POPULARITY_API_SERVICE_PORT"),
	}
}

func InitInquiryApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("INQUIRY_API_SERVICE_HOST") + ":" + GetEnv("INQUIRY_API_SERVICE_PORT"),
	}
}

func InitSendmailApiConfig() ApiConfig {
	return ApiConfig{
		Url: GetEnv("SENDMAIL_API_SERVICE_HOST") + ":" + GetEnv("SENDMAIL_API_SERVICE_PORT"),
	}
}
