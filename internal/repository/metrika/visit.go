package metrika

import (
	"context"
	"fmt"
	sdk "github.com/mg-realcom/metrika-sdk"
	"github.com/rs/zerolog"
)

type visitRepository struct {
	client        *sdk.Client
	fields        string
	source        string
	attachmentDir string
	logger        *zerolog.Logger
}

func NewVisitRepository(client *sdk.Client, attachmentDir, fields, source string, logger *zerolog.Logger) *visitRepository {
	repoLogger := logger.With().Str("metrikaRepository", "visit").Logger()
	return &visitRepository{
		client:        client,
		attachmentDir: attachmentDir,
		fields:        fields,
		source:        source,
		logger:        &repoLogger,
	}
}

type DateRange struct {
	DateTo   string
	DateFrom string
}

func (l *visitRepository) PushLog(ctx context.Context, dateFrom, dateTo string) ([]string, error) {
	l.logger.Trace().Msg("GetVisits")
	reqID, err := l.client.CreateLog(ctx, dateFrom, dateTo, l.fields, l.source)
	if err != nil {
		return nil, err
	}
	if reqID == 0 {
		return nil, fmt.Errorf("Visits: GetRequestID Error")
	}
	parts, err := l.client.GetParts(ctx, reqID)
	if err != nil {
		return nil, err
	}
	files, err := l.client.CollectAllParts(ctx, reqID, parts, l.attachmentDir)
	if err != nil {
		return nil, err
	}
	l.logger.Info().Msg("Visits have been downloaded")
	_, err = l.client.DeleteLog(ctx, int(l.client.CounterId), reqID)
	if err != nil {
		return nil, err
	}
	return files, nil
}

const VisitsFields = "ym:s:visitID,ym:s:counterID,ym:s:watchIDs,ym:s:date,ym:s:dateTime,ym:s:dateTimeUTC,ym:s:isNewUser,ym:s:startURL,ym:s:endURL,ym:s:pageViews,ym:s:visitDuration,ym:s:bounce,ym:s:ipAddress,ym:s:regionCountry,ym:s:regionCity,ym:s:regionCountryID,ym:s:regionCityID,ym:s:clientID,ym:s:counterUserIDHash,ym:s:networkType,ym:s:goalsID,ym:s:goalsSerialNumber,ym:s:goalsDateTime,ym:s:goalsPrice,ym:s:goalsOrder,ym:s:goalsCurrency,ym:s:lastTrafficSource,ym:s:lastAdvEngine,ym:s:lastReferalSource,ym:s:lastSearchEngineRoot,ym:s:lastSearchEngine,ym:s:lastSocialNetwork,ym:s:lastSocialNetworkProfile,ym:s:referer,ym:s:lastDirectClickOrder,ym:s:lastDirectBannerGroup,ym:s:lastDirectClickBanner,ym:s:lastDirectClickOrderName,ym:s:lastClickBannerGroupName,ym:s:lastDirectClickBannerName,ym:s:lastDirectPhraseOrCond,ym:s:lastDirectPlatformType,ym:s:lastDirectPlatform,ym:s:lastDirectConditionType,ym:s:lastCurrencyID,ym:s:from,ym:s:UTMCampaign,ym:s:UTMContent,ym:s:UTMMedium,ym:s:UTMSource,ym:s:UTMTerm,ym:s:openstatAd,ym:s:openstatCampaign,ym:s:openstatService,ym:s:openstatSource,ym:s:lastGCLID,ym:s:firstGCLID,ym:s:lastSignificantGCLID,ym:s:clientTimeZone,ym:s:deviceCategory,ym:s:mobilePhone,ym:s:mobilePhoneModel,ym:s:operatingSystemRoot,ym:s:operatingSystem,ym:s:browser,ym:s:cookieEnabled,ym:s:javascriptEnabled,ym:s:screenFormat,ym:s:physicalScreenWidth,ym:s:physicalScreenHeight,ym:s:purchaseID,ym:s:purchaseDateTime,ym:s:purchaseAffiliation,ym:s:purchaseRevenue,ym:s:purchaseTax,ym:s:purchaseShipping,ym:s:purchaseCoupon,ym:s:purchaseCurrency,ym:s:purchaseProductQuantity,ym:s:productsPurchaseID,ym:s:productsID,ym:s:productsName,ym:s:productsBrand,ym:s:productsCategory,ym:s:productsCategory1,ym:s:productsCategory2,ym:s:productsCategory3,ym:s:productsCategory4,ym:s:productsCategory5,ym:s:productsVariant,ym:s:productsPosition,ym:s:productsPrice,ym:s:productsCurrency,ym:s:productsCoupon,ym:s:productsQuantity,ym:s:impressionsURL,ym:s:impressionsDateTime,ym:s:impressionsProductID,ym:s:impressionsProductName,ym:s:impressionsProductBrand,ym:s:impressionsProductCategory,ym:s:impressionsProductCategory1,ym:s:impressionsProductCategory2,ym:s:impressionsProductCategory3,ym:s:impressionsProductCategory4,ym:s:impressionsProductCategory5,ym:s:impressionsProductVariant,ym:s:impressionsProductPrice,ym:s:impressionsProductCurrency,ym:s:impressionsProductCoupon,ym:s:offlineCallTalkDuration,ym:s:offlineCallHoldDuration,ym:s:offlineCallMissed,ym:s:offlineCallTag,ym:s:offlineCallFirstTimeCaller,ym:s:offlineCallURL,ym:s:parsedParamsKey1,ym:s:parsedParamsKey2,ym:s:parsedParamsKey3,ym:s:parsedParamsKey4,ym:s:parsedParamsKey5,ym:s:parsedParamsKey6,ym:s:parsedParamsKey7,ym:s:parsedParamsKey8,ym:s:parsedParamsKey9,ym:s:parsedParamsKey10"
const VisitsSource = "visits"
