package bq

import (
	"cloud.google.com/go/bigquery"
)

type WatchID struct {
	watch string
}

type GoalID struct {
	goal bigquery.NullString
}

type GoalsSerialNumber struct {
	serialNumber bigquery.NullInt64
}

type GoalsDateTime struct {
	goalDateTime bigquery.NullDateTime
}

type GoalsPrice struct {
	goalPrice bigquery.NullInt64
}

type GoalsOrder struct {
	goalOrder bigquery.NullString
}

type GoalsCurrency struct {
	goalCurrency bigquery.NullString
}

type PurchaseID struct {
	ID bigquery.NullString
}

type PurchaseDateTime struct {
	purchaseDateTime bigquery.NullDateTime
}

type PurchaseAffiliation struct {
	PurchaseAffiliation bigquery.NullString
}

type PurchaseRevenue struct {
	PurchaseRevenue bigquery.NullFloat64
}

type PurchaseTax struct {
	PurchaseTax bigquery.NullFloat64
}

type PurchaseShipping struct {
	PurchaseShipping bigquery.NullFloat64
}

type PurchaseCoupon struct {
	PurchaseCoupon bigquery.NullString
}

type PurchaseCurrency struct {
	PurchaseCurrency bigquery.NullString
}

type PurchaseProductQuantity struct {
	PurchaseProductQuantity bigquery.NullInt64
}

type ProductsPurchaseID struct {
	ProductsPurchaseID bigquery.NullString
}

type ProductsID struct {
	ProductsID bigquery.NullString
}

type ProductsName struct {
	ProductsName bigquery.NullString
}
type ProductsBrand struct {
	ProductsBrand bigquery.NullString
}
type ProductsCategory struct {
	ProductsCategory bigquery.NullString
}
type ProductsCategory1 struct {
	ProductsCategory1 bigquery.NullString
}
type ProductsCategory2 struct {
	ProductsCategory2 bigquery.NullString
}
type ProductsCategory3 struct {
	ProductsCategory3 bigquery.NullString
}
type ProductsCategory4 struct {
	ProductsCategory4 bigquery.NullString
}
type ProductsCategory5 struct {
	ProductsCategory5 bigquery.NullString
}

type ProductsVariant struct {
	ProductsVariant bigquery.NullString
}

type ProductsPosition struct {
	ProductsPosition bigquery.NullInt64
}

type ProductsPrice struct {
	ProductsPrice bigquery.NullFloat64
}

type ProductsCurrency struct {
	ProductsCurrency bigquery.NullString
}

type ProductsCoupon struct {
	ProductsCoupon bigquery.NullString
}

type ProductsQuantity struct {
	ProductsQuantity bigquery.NullInt64
}

type ImpressionsURL struct {
	ImpressionsURL bigquery.NullString
}

type ImpressionsDateTime struct {
	ImpressionsDateTime bigquery.NullString
}

type ImpressionsProductID struct {
	ImpressionsProductID bigquery.NullString
}

type ImpressionsProductName struct {
	ImpressionsProductName bigquery.NullString
}

type ImpressionsProductBrand struct {
	ImpressionsProductBrand bigquery.NullString
}

type ImpressionsProductCategory struct {
	ImpressionsProductCategory bigquery.NullString
}

type ImpressionsProductCategory1 struct {
	ImpressionsProductCategory1 bigquery.NullString
}

type ImpressionsProductCategory2 struct {
	ImpressionsProductCategory2 bigquery.NullString
}

type ImpressionsProductCategory3 struct {
	ImpressionsProductCategory3 bigquery.NullString
}

type ImpressionsProductCategory4 struct {
	ImpressionsProductCategory4 bigquery.NullString
}

type ImpressionsProductCategory5 struct {
	ImpressionsProductCategory5 bigquery.NullString
}

type ImpressionsProductVariant struct {
	ImpressionsProductVariant bigquery.NullString
}

type ImpressionsProductPrice struct {
	ImpressionsProductPrice bigquery.NullString
}

type ImpressionsProductCurrency struct {
	ImpressionsProductCurrency bigquery.NullString
}

type ImpressionsProductCoupon struct {
	ImpressionsProductCoupon bigquery.NullString
}

type OfflineCallTalkDuration struct {
	OfflineCallTalkDuration bigquery.NullInt64
}

type OfflineCallHoldDuration struct {
	OfflineCallHoldDuration bigquery.NullInt64
}
type OfflineCallMissed struct {
	OfflineCallMissed bigquery.NullInt64
}

type OfflineCallTag struct {
	OfflineCallTag bigquery.NullString
}
type OfflineCallFirstTimeCaller struct {
	OfflineCallFirstTimeCaller bigquery.NullInt64
}
type OfflineCallURL struct {
	OfflineCallURL bigquery.NullString
}
type ParsedParamsKey1 struct {
	ParsedParamsKey1 bigquery.NullString
}
type ParsedParamsKey2 struct {
	ParsedParamsKey2 bigquery.NullString
}
type ParsedParamsKey3 struct {
	ParsedParamsKey3 bigquery.NullString
}
type ParsedParamsKey4 struct {
	ParsedParamsKey4 bigquery.NullString
}
type ParsedParamsKey5 struct {
	ParsedParamsKey5 bigquery.NullString
}
type ParsedParamsKey6 struct {
	ParsedParamsKey6 bigquery.NullString
}
type ParsedParamsKey7 struct {
	ParsedParamsKey7 bigquery.NullString
}
type ParsedParamsKey8 struct {
	ParsedParamsKey8 bigquery.NullString
}
type ParsedParamsKey9 struct {
	ParsedParamsKey9 bigquery.NullString
}

type ParsedParamsKey10 struct {
	ParsedParamsKey10 bigquery.NullString
}

type VisitSchema struct {
	VisitID                     bigquery.NullInt64    `bigquery:"visitID"`
	CounterID                   bigquery.NullInt64    `bigquery:"counterID"`
	WatchIDS                    bigquery.NullString   `bigquery:"watchIDs"`
	Date                        bigquery.NullDate     `bigquery:"date"`
	DateTime                    bigquery.NullDateTime `bigquery:"dateTime"`
	DateTimeUTC                 bigquery.NullDateTime `bigquery:"dateTimeUTC"`
	IsNewUser                   bigquery.NullBool     `bigquery:"isNewUser"`
	StartURL                    bigquery.NullString   `bigquery:"startURL"`
	EndURL                      bigquery.NullString   `bigquery:"endURL"`
	PageViews                   bigquery.NullInt64    `bigquery:"pageViews"`
	VisitDuration               bigquery.NullInt64    `bigquery:"visitDuration"`
	Bounce                      bigquery.NullBool     `bigquery:"bounce"`
	IPAddress                   bigquery.NullString   `bigquery:"ipAddress"`
	RegionCountry               bigquery.NullString   `bigquery:"regionCountry"`
	RegionCity                  bigquery.NullString   `bigquery:"regionCity"`
	RegionCountryID             bigquery.NullInt64    `bigquery:"regionCountryID"`
	RegionCityID                bigquery.NullInt64    `bigquery:"regionCityID"`
	ClientID                    bigquery.NullString   `bigquery:"clientID"`
	CounterUserIDHash           bigquery.NullString   `bigquery:"counterUserIDHash"`
	NetworkType                 bigquery.NullString   `bigquery:"networkType"`
	GoalsID                     bigquery.NullString   `bigquery:"goalsID"`
	GoalsSerialNumber           bigquery.NullString   `bigquery:"goalsSerialNumber"`
	GoalsDateTime               bigquery.NullString   `bigquery:"goalsDateTime"`
	GoalsPrice                  bigquery.NullString   `bigquery:"goalsPrice"`
	GoalsOrder                  bigquery.NullString   `bigquery:"goalsOrder"`
	GoalsCurrency               bigquery.NullString   `bigquery:"goalsCurrency"`
	LastTrafficSource           bigquery.NullString   `bigquery:"lastTrafficSource"`
	LastAdvEngine               bigquery.NullString   `bigquery:"lastAdvEngine"`
	LastReferalSource           bigquery.NullString   `bigquery:"lastReferalSource"`
	LastSearchEngineRoot        bigquery.NullString   `bigquery:"lastSearchEngineRoot"`
	LastSearchEngine            bigquery.NullString   `bigquery:"lastSearchEngine"`
	LastSocialNetwork           bigquery.NullString   `bigquery:"lastSocialNetwork"`
	LastSocialNetworkProfile    bigquery.NullString   `bigquery:"lastSocialNetworkProfile"`
	Referer                     bigquery.NullString   `bigquery:"referer"`
	LastDirectClickOrder        bigquery.NullInt64    `bigquery:"lastDirectClickOrder"`
	LastDirectBannerGroup       bigquery.NullInt64    `bigquery:"lastDirectBannerGroup"`
	LastDirectClickBanner       bigquery.NullString   `bigquery:"lastDirectClickBanner"`
	LastDirectClickOrderName    bigquery.NullString   `bigquery:"lastDirectClickOrderName"`
	LastClickBannerGroupName    bigquery.NullString   `bigquery:"lastClickBannerGroupName"`
	LastDirectClickBannerName   bigquery.NullString   `bigquery:"lastDirectClickBannerName"`
	LastDirectPhraseOrCond      bigquery.NullString   `bigquery:"lastDirectPhraseOrCond"`
	LastDirectPlatformType      bigquery.NullString   `bigquery:"lastDirectPlatformType"`
	LastDirectPlatform          bigquery.NullString   `bigquery:"lastDirectPlatform"`
	LastDirectConditionType     bigquery.NullString   `bigquery:"lastDirectConditionType"`
	LastCurrencyID              bigquery.NullString   `bigquery:"lastCurrencyID"`
	From                        bigquery.NullString   `bigquery:"from"`
	UTMCampaign                 bigquery.NullString   `bigquery:"UTMCampaign"`
	UTMContent                  bigquery.NullString   `bigquery:"UTMContent"`
	UTMMedium                   bigquery.NullString   `bigquery:"UTMMedium"`
	UTMSource                   bigquery.NullString   `bigquery:"UTMSource"`
	UTMTerm                     bigquery.NullString   `bigquery:"UTMTerm"`
	OpenstatAd                  bigquery.NullString   `bigquery:"openstatAd"`
	OpenstatCampaign            bigquery.NullString   `bigquery:"openstatCampaign"`
	OpenstatService             bigquery.NullString   `bigquery:"openstatService"`
	OpenstatSource              bigquery.NullString   `bigquery:"openstatSource"`
	LastGCLID                   bigquery.NullString   `bigquery:"lastGCLID"`
	FirstGCLID                  bigquery.NullString   `bigquery:"firstGCLID"`
	LastSignificantGCLID        bigquery.NullString   `bigquery:"lastSignificantGCLID"`
	ClientTimeZone              bigquery.NullInt64    `bigquery:"clientTimeZone"`
	DeviceCategory              bigquery.NullString   `bigquery:"deviceCategory"`
	MobilePhone                 bigquery.NullString   `bigquery:"mobilePhone"`
	MobilePhoneModel            bigquery.NullString   `bigquery:"mobilePhoneModel"`
	OperatingSystemRoot         bigquery.NullString   `bigquery:"operatingSystemRoot"`
	OperatingSystem             bigquery.NullString   `bigquery:"operatingSystem"`
	Browser                     bigquery.NullString   `bigquery:"browser"`
	CookieEnabled               bigquery.NullBool     `bigquery:"cookieEnabled"`
	JavascriptEnabled           bigquery.NullBool     `bigquery:"javascriptEnabled"`
	ScreenFormat                bigquery.NullString   `bigquery:"screenFormat"`
	PhysicalScreenWidth         bigquery.NullInt64    `bigquery:"physicalScreenWidth"`
	PhysicalScreenHeight        bigquery.NullInt64    `bigquery:"physicalScreenHeight"`
	PurchaseID                  bigquery.NullString   `bigquery:"purchaseID"`
	PurchaseDateTime            bigquery.NullString   `bigquery:"purchaseDateTime"`
	PurchaseAffiliation         bigquery.NullString   `bigquery:"purchaseAffiliation"`
	PurchaseRevenue             bigquery.NullString   `bigquery:"purchaseRevenue"`
	PurchaseTax                 bigquery.NullString   `bigquery:"purchaseTax"`
	PurchaseShipping            bigquery.NullString   `bigquery:"purchaseShipping"`
	PurchaseCoupon              bigquery.NullString   `bigquery:"purchaseCoupon"`
	PurchaseCurrency            bigquery.NullString   `bigquery:"purchaseCurrency"`
	PurchaseProductQuantity     bigquery.NullString   `bigquery:"purchaseProductQuantity"`
	ProductsPurchaseID          bigquery.NullString   `bigquery:"productsPurchaseID"`
	ProductsID                  bigquery.NullString   `bigquery:"productsID"`
	ProductsName                bigquery.NullString   `bigquery:"productsName"`
	ProductsBrand               bigquery.NullString   `bigquery:"productsBrand"`
	ProductsCategory            bigquery.NullString   `bigquery:"productsCategory"`
	ProductsCategory1           bigquery.NullString   `bigquery:"productsCategory1"`
	ProductsCategory2           bigquery.NullString   `bigquery:"productsCategory2"`
	ProductsCategory3           bigquery.NullString   `bigquery:"productsCategory3"`
	ProductsCategory4           bigquery.NullString   `bigquery:"productsCategory4"`
	ProductsCategory5           bigquery.NullString   `bigquery:"productsCategory5"`
	ProductsVariant             bigquery.NullString   `bigquery:"productsVariant"`
	ProductsPosition            bigquery.NullString   `bigquery:"productsPosition"`
	ProductsPrice               bigquery.NullString   `bigquery:"productsPrice"`
	ProductsCurrency            bigquery.NullString   `bigquery:"productsCurrency"`
	ProductsCoupon              bigquery.NullString   `bigquery:"productsCoupon"`
	ProductsQuantity            bigquery.NullString   `bigquery:"productsQuantity"`
	ImpressionsURL              bigquery.NullString   `bigquery:"impressionsURL"`
	ImpressionsDateTime         bigquery.NullString   `bigquery:"impressionsDateTime"`
	ImpressionsProductID        bigquery.NullString   `bigquery:"impressionsProductID"`
	ImpressionsProductName      bigquery.NullString   `bigquery:"impressionsProductName"`
	ImpressionsProductBrand     bigquery.NullString   `bigquery:"impressionsProductBrand"`
	ImpressionsProductCategory  bigquery.NullString   `bigquery:"impressionsProductCategory"`
	ImpressionsProductCategory1 bigquery.NullString   `bigquery:"impressionsProductCategory1"`
	ImpressionsProductCategory2 bigquery.NullString   `bigquery:"impressionsProductCategory2"`
	ImpressionsProductCategory3 bigquery.NullString   `bigquery:"impressionsProductCategory3"`
	ImpressionsProductCategory4 bigquery.NullString   `bigquery:"impressionsProductCategory4"`
	ImpressionsProductCategory5 bigquery.NullString   `bigquery:"impressionsProductCategory5"`
	ImpressionsProductVariant   bigquery.NullString   `bigquery:"impressionsProductVariant"`
	ImpressionsProductPrice     bigquery.NullString   `bigquery:"impressionsProductPrice"`
	ImpressionsProductCurrency  bigquery.NullString   `bigquery:"impressionsProductCurrency"`
	ImpressionsProductCoupon    bigquery.NullString   `bigquery:"impressionsProductCoupon"`
	OfflineCallTalkDuration     bigquery.NullString   `bigquery:"offlineCallTalkDuration"`
	OfflineCallHoldDuration     bigquery.NullString   `bigquery:"offlineCallHoldDuration"`
	OfflineCallMissed           bigquery.NullString   `bigquery:"offlineCallMissed"`
	OfflineCallTag              bigquery.NullString   `bigquery:"offlineCallTag"`
	OfflineCallFirstTimeCaller  bigquery.NullString   `bigquery:"offlineCallFirstTimeCaller"`
	OfflineCallURL              bigquery.NullString   `bigquery:"offlineCallURL"`
	ParsedParamsKey1            bigquery.NullString   `bigquery:"parsedParamsKey1"`
	ParsedParamsKey2            bigquery.NullString   `bigquery:"parsedParamsKey2"`
	ParsedParamsKey3            bigquery.NullString   `bigquery:"parsedParamsKey3"`
	ParsedParamsKey4            bigquery.NullString   `bigquery:"parsedParamsKey4"`
	ParsedParamsKey5            bigquery.NullString   `bigquery:"parsedParamsKey5"`
	ParsedParamsKey6            bigquery.NullString   `bigquery:"parsedParamsKey6"`
	ParsedParamsKey7            bigquery.NullString   `bigquery:"parsedParamsKey7"`
	ParsedParamsKey8            bigquery.NullString   `bigquery:"parsedParamsKey8"`
	ParsedParamsKey9            bigquery.NullString   `bigquery:"parsedParamsKey9"`
	ParsedParamsKey10           bigquery.NullString   `bigquery:"parsedParamsKey10"`
}

type HitSchema struct {
	WatchID                  bigquery.NullInt64    `bigquery:"watchID"`
	CounterID                bigquery.NullInt64    `bigquery:"counterID"`
	Date                     bigquery.NullDate     `bigquery:"date"`
	DateTime                 bigquery.NullDateTime `bigquery:"dateTime"`
	Title                    bigquery.NullString   `bigquery:"title"`
	URL                      bigquery.NullString   `bigquery:"URL"`
	Referrer                 bigquery.NullString   `bigquery:"referrer"`
	UTMCampaign              bigquery.NullString   `bigquery:"UTMCampaign"`
	UTMContent               bigquery.NullString   `bigquery:"UTMContent"`
	UTMMedium                bigquery.NullString   `bigquery:"UTMMedium"`
	UTMSource                bigquery.NullString   `bigquery:"UTMSource"`
	UTMTerm                  bigquery.NullString   `bigquery:"UTMTerm"`
	Browser                  bigquery.NullString   `bigquery:"browser"`
	BrowserMajorVersion      bigquery.NullInt64    `bigquery:"browserMajorVersion"`
	BrowserMinorVersion      bigquery.NullInt64    `bigquery:"browserMinorVersion"`
	BrowserCountry           bigquery.NullString   `bigquery:"browserCountry"`
	BrowserEngine            bigquery.NullString   `bigquery:"browserEngine"`
	BrowserEngineVersion1    bigquery.NullInt64    `bigquery:"browserEngineVersion1"`
	BrowserEngineVersion2    bigquery.NullInt64    `bigquery:"browserEngineVersion2"`
	BrowserEngineVersion3    bigquery.NullInt64    `bigquery:"browserEngineVersion3"`
	BrowserEngineVersion4    bigquery.NullInt64    `bigquery:"browserEngineVersion4"`
	BrowserLanguage          bigquery.NullString   `bigquery:"browserLanguage"`
	ClientTimeZone           bigquery.NullInt64    `bigquery:"clientTimeZone"`
	CookieEnabled            bigquery.NullBool     `bigquery:"cookieEnabled"`
	DeviceCategory           bigquery.NullString   `bigquery:"deviceCategory"`
	From                     bigquery.NullString   `bigquery:"from"`
	HasGCLID                 bigquery.NullBool     `bigquery:"hasGCLID"`
	GCLID                    bigquery.NullString   `bigquery:"GCLID"`
	IPAddress                bigquery.NullString   `bigquery:"ipAddress"`
	JavaScriptEnabled        bigquery.NullBool     `bigquery:"javascriptEnabled"`
	MobilePhone              bigquery.NullString   `bigquery:"mobilePhone"`
	MobilePhoneModel         bigquery.NullString   `bigquery:"mobilePhoneModel"`
	OpenstatAd               bigquery.NullString   `bigquery:"openstatAd"`
	OpenstatCampaign         bigquery.NullString   `bigquery:"openstatCampaign"`
	OpenstatService          bigquery.NullString   `bigquery:"openstatService"`
	OpenstatSource           bigquery.NullString   `bigquery:"openstatSource"`
	OperatingSystem          bigquery.NullString   `bigquery:"operatingSystem"`
	OperatingSystemRoot      bigquery.NullString   `bigquery:"operatingSystemRoot"`
	PhysicalScreenHeight     bigquery.NullInt64    `bigquery:"physicalScreenHeight"`
	PhysicalScreenWidth      bigquery.NullInt64    `bigquery:"physicalScreenWidth"`
	RegionCity               bigquery.NullString   `bigquery:"regionCity"`
	RegionCountry            bigquery.NullString   `bigquery:"regionCountry"`
	RegionCityID             bigquery.NullInt64    `bigquery:"regionCityID"`
	RegionCountryID          bigquery.NullInt64    `bigquery:"regionCountryID"`
	ScreenColors             bigquery.NullInt64    `bigquery:"screenColors"`
	ScreenFormat             bigquery.NullString   `bigquery:"screenFormat"`
	ScreenHeight             bigquery.NullInt64    `bigquery:"screenHeight"`
	ScreenOrientation        bigquery.NullInt64    `bigquery:"screenOrientation"`
	ScreenWidth              bigquery.NullInt64    `bigquery:"screenWidth"`
	WindowClientHeight       bigquery.NullInt64    `bigquery:"windowClientHeight"`
	WindowClientWidth        bigquery.NullInt64    `bigquery:"windowClientWidth"`
	LastTrafficSource        bigquery.NullString   `bigquery:"lastTrafficSource"`
	LastSearchEngine         bigquery.NullString   `bigquery:"lastSearchEngine"`
	LastSearchEngineRoot     bigquery.NullString   `bigquery:"lastSearchEngineRoot"`
	LastAdvEngine            bigquery.NullString   `bigquery:"lastAdvEngine"`
	Artificial               bigquery.NullString   `bigquery:"artificial"`
	PageCharset              bigquery.NullString   `bigquery:"pageCharset"`
	IsPageView               bigquery.NullBool     `bigquery:"isPageView"`
	IsTurboPage              bigquery.NullBool     `bigquery:"isTurboPage"`
	Link                     bigquery.NullString   `bigquery:"link"`
	Download                 bigquery.NullString   `bigquery:"download"`
	NotBounce                bigquery.NullString   `bigquery:"notBounce"`
	LastSocialNetwork        bigquery.NullString   `bigquery:"lastSocialNetwork"`
	HTTPError                bigquery.NullString   `bigquery:"httpError"`
	ClientID                 bigquery.NullString   `bigquery:"clientID"`
	CounterUserIDHash        bigquery.NullString   `bigquery:"counterUserIDHash"`
	NetworkType              bigquery.NullString   `bigquery:"networkType"`
	LastSocialNetworkProfile bigquery.NullString   `bigquery:"lastSocialNetworkProfile"`
	GoalsID                  bigquery.NullString   `bigquery:"goalsID"`
	ShareService             bigquery.NullString   `bigquery:"shareService"`
	ShareURL                 bigquery.NullString   `bigquery:"shareURL"`
	ShareTitle               bigquery.NullString   `bigquery:"shareTitle"`
	IFrame                   bigquery.NullBool     `bigquery:"iFrame"`
	RecommendationSystem     bigquery.NullString   `bigquery:"recommendationSystem"`
	Messenger                bigquery.NullString   `bigquery:"messenger"`
	ParsedParamsKey1         bigquery.NullString   `bigquery:"parsedParamsKey1"`
	ParsedParamsKey2         bigquery.NullString   `bigquery:"parsedParamsKey2"`
	ParsedParamsKey3         bigquery.NullString   `bigquery:"parsedParamsKey3"`
	ParsedParamsKey4         bigquery.NullString   `bigquery:"parsedParamsKey4"`
	ParsedParamsKey5         bigquery.NullString   `bigquery:"parsedParamsKey5"`
	ParsedParamsKey6         bigquery.NullString   `bigquery:"parsedParamsKey6"`
	ParsedParamsKey7         bigquery.NullString   `bigquery:"parsedParamsKey7"`
	ParsedParamsKey8         bigquery.NullString   `bigquery:"parsedParamsKey8"`
	ParsedParamsKey9         bigquery.NullString   `bigquery:"parsedParamsKey9"`
	ParsedParamsKey10        bigquery.NullString   `bigquery:"parsedParamsKey10"`
}
