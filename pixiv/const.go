package pixiv

const (
	userAgent = "PixivAndroidApp/5.0.235 (Android 5.1.1; MI 6 )"

	contentType    = "application/x-www-form-urlencoded"
	acceptLanguage = "zh_CN"
	appOS          = "android"
	appVersion     = "5.0.235"

	clientId      = "MOBrBDS8blbauoSck0ZfDbtuzpyT"
	clientSecret  = "lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"
	hashSecret    = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	timeFormatter = "2006-01-02T15:04:05-07:00"

	apiHost          = "https://app-api.pixiv.net"
	oauthUrl         = "https://oauth.secure.pixiv.net/auth/token"
	illustDetailUrl  = apiHost + "/v1/illust/detail"
	illustRelatedUrl = apiHost + "/v2/illust/related"
	userDetailUrl    = apiHost + "/v1/user/detail"
	userIllustUrl    = apiHost + "/v1/user/illusts"
	rankUrl          = apiHost + "/v1/illust/ranking"
	searchUrl        = apiHost + "/v1/search/illust"
)
