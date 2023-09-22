package metrika

import (
	"context"
	sdk "github.com/mg-realcom/metrika-sdk"
	"github.com/rs/zerolog"
)

type hitRepository struct {
	client        *sdk.Client
	fields        string
	source        string
	attachmentDir string
	logger        *zerolog.Logger
}

func NewHitRepository(client *sdk.Client, attachmentDir, fields, source string, logger *zerolog.Logger) *hitRepository {
	return &hitRepository{
		client:        client,
		attachmentDir: attachmentDir,
		fields:        fields,
		source:        source,
		logger:        logger,
	}
}

func (l *hitRepository) PushHits(ctx context.Context, dateFrom, dateTo string) ([]string, error) {
	reqID, err := l.client.CreateLog(ctx, dateFrom, dateTo, l.fields, l.source)
	if err != nil {
		return nil, err
	}
	parts, err := l.client.GetParts(ctx, reqID)
	if err != nil {
		return nil, err
	}
	files, err := l.client.CollectAllParts(ctx, reqID, parts, l.attachmentDir)
	if err != nil {
		return nil, err
	}
	_, err = l.client.DeleteLog(ctx, int(l.client.CounterId), reqID)
	if err != nil {
		return nil, err
	}
	return files, nil
}

const HitsFields = "ym:pv:watchID,ym:pv:counterID,ym:pv:date,ym:pv:dateTime,ym:pv:title,ym:pv:URL,ym:pv:referer,ym:pv:UTMCampaign,ym:pv:UTMContent,ym:pv:UTMMedium,ym:pv:UTMSource,ym:pv:UTMTerm,ym:pv:browser,ym:pv:browserMajorVersion,ym:pv:browserMinorVersion,ym:pv:browserCountry,ym:pv:browserEngine,ym:pv:browserEngineVersion1,ym:pv:browserEngineVersion2,ym:pv:browserEngineVersion3,ym:pv:browserEngineVersion4,ym:pv:browserLanguage,ym:pv:clientTimeZone,ym:pv:cookieEnabled,ym:pv:deviceCategory,ym:pv:from,ym:pv:hasGCLID,ym:pv:GCLID,ym:pv:ipAddress,ym:pv:javascriptEnabled,ym:pv:mobilePhone,ym:pv:mobilePhoneModel,ym:pv:openstatAd,ym:pv:openstatCampaign,ym:pv:openstatService,ym:pv:openstatSource,ym:pv:operatingSystem,ym:pv:operatingSystemRoot,ym:pv:physicalScreenHeight,ym:pv:physicalScreenWidth,ym:pv:regionCity,ym:pv:regionCountry,ym:pv:regionCityID,ym:pv:regionCountryID,ym:pv:screenColors,ym:pv:screenFormat,ym:pv:screenHeight,ym:pv:screenOrientation,ym:pv:screenWidth,ym:pv:windowClientHeight,ym:pv:windowClientWidth,ym:pv:lastTrafficSource,ym:pv:lastSearchEngine,ym:pv:lastSearchEngineRoot,ym:pv:lastAdvEngine,ym:pv:artificial,ym:pv:pageCharset,ym:pv:isPageView,ym:pv:isTurboPage,ym:pv:link,ym:pv:download,ym:pv:notBounce,ym:pv:lastSocialNetwork,ym:pv:httpError,ym:pv:clientID,ym:pv:counterUserIDHash,ym:pv:networkType,ym:pv:lastSocialNetworkProfile,ym:pv:goalsID,ym:pv:shareService,ym:pv:shareURL,ym:pv:shareTitle,ym:pv:iFrame,ym:pv:recommendationSystem,ym:pv:messenger,ym:pv:parsedParamsKey1,ym:pv:parsedParamsKey2,ym:pv:parsedParamsKey3,ym:pv:parsedParamsKey4,ym:pv:parsedParamsKey5,ym:pv:parsedParamsKey6,ym:pv:parsedParamsKey7,ym:pv:parsedParamsKey8,ym:pv:parsedParamsKey9,ym:pv:parsedParamsKey10"
const HitsSource = "hits"
