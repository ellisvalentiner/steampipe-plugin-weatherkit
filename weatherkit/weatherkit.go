package weatherkit

type Weather struct {
	CurrentWeather   CurrentWeatherData         `json:"currentWeather,omitempty"`
	DailyForecast    DailyForecastData          `json:"forecastDaily,omitempty"`
	HourlyForecast   HourlyForecastData         `json:"forecastHourly,omitempty"`
	NextHourForecast NextHourForecastData       `json:"forecastNextHour,omitempty"`
	WeatherAlerts    WeatherAlertCollectionData `json:"weatherAlerts,omitempty"`
}

type CurrentWeatherData struct {
	AsOf                   *string         `json:"asOf,omitempty"`
	CloudCover             *float32        `json:"cloudCover,omitempty"`
	ConditionCode          *string         `json:"conditionCode,omitempty"`
	Daylight               *bool           `json:"daylight,omitempty"`
	Humidity               *float32        `json:"humidity,omitempty"`
	PrecipitationIntensity *float32        `json:"precipitationIntensity,omitempty"`
	Pressure               *float32        `json:"pressure,omitempty"`
	PressureTrend          *string         `json:"pressureTrend,omitempty"`
	Temperature            *float32        `json:"temperature,omitempty"`
	TemperatureApparent    *float32        `json:"temperatureApparent,omitempty"`
	TemperatureDewPoint    *float32        `json:"temperatureDewPoint,omitempty"`
	UvIndex                *int            `json:"uvIndex,omitempty"`
	Visibility             *float32        `json:"visibility,omitempty"`
	WindDirection          *int            `json:"windDirection,omitempty"`
	WindGust               *float32        `json:"windGust,omitempty"`
	WindSpeed              *float32        `json:"windSpeed,omitempty"`
	Metadata               WeatherMetadata `json:"metadata,omitempty"`
}

type DailyForecastData struct {
	Days     []DayWeatherConditions `json:"days,omitempty"`
	Metadata WeatherMetadata        `json:"metadata,omitempty"`
}

type HourlyForecastData struct {
	Hours    []HourWeatherConditions `json:"hours,omitempty"`
	Metadata WeatherMetadata         `json:"metadata,omitempty"`
}

type DayWeatherConditions struct {
	ConditionCode       *string          `json:"conditionCode,omitempty"`
	DaytimeForecast     *DayPartForecast `json:"daytimeForecast,omitempty"`
	ForecastEnd         *string          `json:"forecastEnd,omitempty"`
	ForecastStart       *string          `json:"forecastStart,omitempty"`
	MaxUvIndex          *int             `json:"maxUvIndex,omitempty"`
	MoonPhase           *string          `json:"moonPhase,omitempty"`
	Moonrise            *string          `json:"moonrise,omitempty"`
	Moonset             *string          `json:"moonset,omitempty"`
	OvernightForecast   *DayPartForecast `json:"overnightForecast,omitempty"`
	PrecipitationAmount *float32         `json:"precipitationAmount,omitempty"`
	PrecipitationChance *float32         `json:"precipitationChance,omitempty"`
	PrecipitationType   *string          `json:"precipitationType,omitempty"`
	SnowfallAmount      *float32         `json:"snowfallAmount,omitempty"`
	SolarMidnight       *string          `json:"solarMidnight,omitempty"`
	SolarNoon           *string          `json:"solarNoon,omitempty"`
	Sunrise             *string          `json:"sunrise,omitempty"`
	SunriseAstronomical *string          `json:"sunriseAstronomical,omitempty"`
	SunriseCivil        *string          `json:"sunriseCivil,omitempty"`
	SunriseNautical     *string          `json:"sunriseNautical,omitempty"`
	Sunset              *string          `json:"sunset,omitempty"`
	SunsetAstronomical  *string          `json:"sunsetAstronomical,omitempty"`
	SunsetCivil         *string          `json:"sunsetCivil,omitempty"`
	SunsetNautical      *string          `json:"sunsetNautical,omitempty"`
	TemperatureMax      *float32         `json:"temperatureMax,omitempty"`
	TemperatureMin      *float32         `json:"temperatureMin,omitempty"`
}

type DayPartForecast struct {
	CloudCover          *float32 `json:"cloudCover,omitempty"`
	ConditionCode       *string  `json:"conditionCode,omitempty"`
	ForecastEnd         *string  `json:"forecastEnd,omitempty"`
	ForecastStart       *string  `json:"forecastStart,omitempty"`
	Humidity            *float32 `json:"humidity,omitempty"`
	PrecipitationAmount *float32 `json:"precipitationAmount,omitempty"`
	PrecipitationChance *float32 `json:"precipitationChance,omitempty"`
	PrecipitationType   *string  `json:"precipitationType,omitempty"`
	SnowfallAmount      *float32 `json:"snowfallAmount,omitempty"`
	WindDirection       *int     `json:"windDirection,omitempty"`
	WindSpeed           *float32 `json:"windSpeed,omitempty"`
}

type HourWeatherConditions struct {
	CloudCover          *float32 `json:"cloudCover,omitempty"`
	ConditionCode       *string  `json:"conditionCode,omitempty"`
	Daylight            *bool    `json:"daylight,omitempty"`
	ForecastStart       *string  `json:"forecastStart,omitempty"`
	Humidity            *float32 `json:"humidity,omitempty"`
	PrecipitationChance *float32 `json:"precipitationChance,omitempty"`
	PrecipitationType   *string  `json:"precipitationType,omitempty"`
	Pressure            *float32 `json:"pressure,omitempty"`
	PressureTrend       *string  `json:"pressureTrend,omitempty"`
	SnowfallIntensity   *float32 `json:"snowfallIntensity,omitempty"`
	Temperature         *float32 `json:"temperature,omitempty"`
	TemperatureApparent *float32 `json:"temperatureApparent,omitempty"`
	TemperatureDewPoint *float32 `json:"temperatureDewPoint,omitempty"`
	UvIndex             *int     `json:"uvIndex,omitempty"`
	Visibility          *float32 `json:"visibility,omitempty"`
	WindDirection       *int     `json:"windDirection,omitempty"`
	WindGust            *float32 `json:"windGust,omitempty"`
	WindSpeed           *float32 `json:"windSpeed,omitempty"`
	PrecipitationAmount *float32 `json:"precipitationAmount,omitempty"`
}

type NextHourForecastData struct {
	ForecastEnd   string                  `json:"forecastEnd,omitempty"`
	ForecastStart string                  `json:"forecastStart,omitempty"`
	Minutes       []ForecastMinute        `json:"minutes,omitempty"`
	Summary       []ForecastPeriodSummary `json:"summary,omitempty"`
	Metadata      WeatherMetadata         `json:"metadata,omitempty"`
}

type ForecastMinute struct {
	PrecipitationChance    *float32 `json:"precipitationChance,omitempty"`
	PrecipitationIntensity *float32 `json:"precipitationIntensity,omitempty"`
	StartTime              *string  `json:"startTime,omitempty"`
}

type ForecastPeriodSummary struct {
	Condition              *string  `json:"condition,omitempty"`
	EndTime                *string  `json:"endTime,omitempty"`
	PrecipitationChance    *float32 `json:"precipitationChance,omitempty"`
	PrecipitationIntensity *float32 `json:"precipitationIntensity,omitempty"`
	StartTime              *string  `json:"startTime,omitempty"`
}

type WeatherAlertCollectionData struct {
	Alerts   []WeatherAlertSummary `json:"alerts,omitempty"`
	Metadata WeatherMetadata       `json:"metadata,omitempty"`
}

type WeatherAlertSummary struct {
	AreaId        *string   `json:"areaId,omitempty"`
	AreaName      *string   `json:"areaName,omitempty"`
	Certainty     *string   `json:"certainty,omitempty"`
	CountryCode   *string   `json:"countryCode,omitempty"`
	Description   *string   `json:"description,omitempty"`
	DetailsUrl    *string   `json:"detailsUrl,omitempty"`
	EffectiveTime *string   `json:"effectiveTime,omitempty"`
	EventEndTime  *string   `json:"eventEndTime,omitempty"`
	ExpireTime    *string   `json:"expireTime,omitempty"`
	Id            *string   `json:"id,omitempty"`
	IssuedTime    *string   `json:"issuedTime,omitempty"`
	Responses     *[]string `json:"responses,omitempty"`
	Severity      *string   `json:"severity,omitempty"`
	Source        *string   `json:"source,omitempty"`
	Urgency       *string   `json:"urgency,omitempty"`
}

type WeatherMetadata struct {
	AttributionUrl *string  `json:"attributionUrl,omitempty"`
	ExpireTime     *string  `json:"expireTime,omitempty"`
	Latitude       *float32 `json:"latitude,omitempty"`
	Longitude      *float32 `json:"longitude,omitempty"`
	ReadTime       *string  `json:"readTime,omitempty"`
	ReportedTime   *string  `json:"reportedTime,omitempty"`
	Units          *string  `json:"units,omitempty"`
	Version        *int     `json:"version,omitempty"`
}
