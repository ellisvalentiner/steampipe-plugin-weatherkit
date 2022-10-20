package weatherkit

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	config     *weatherKitConfig
	logger     hclog.Logger
}

func NewClient(ctx context.Context, httpClient *http.Client, config *weatherKitConfig) *Client {
	return &Client{
		httpClient: httpClient,
		config:     config,
		logger:     plugin.Logger(ctx),
	}
}

func (c *Client) Get(ctx context.Context, url string, v interface{}) error {
	req, _ := c.NewRequest(ctx, http.MethodGet, url)
	err := c.DoRequest(req, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) NewRequest(ctx context.Context, method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var token string
	if c.config.Token != nil {
		token = *c.config.Token
	} else {
		token = c.createJwt()
	}

	req.Header.Add("Authorization", "Bearer "+token)

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) checkResponseStatus(r *http.Response) {
	status := r.Status
	switch status {
	case "400 Bad Request":
		c.logger.Error("DoRequest", "status", status, "message", "The server is unable to process the request due to an invalid parameter value.")
		panic("\nthe server is unable to process the request due to an invalid parameter value.\nPlease file an issue at https://github.com/ellisvalentiner/steampipe-plugin-weatherkit")
	case "401 Unauthorized":
		c.logger.Error("DoRequest", "status", status, "message", "The request isn’t authorized or doesn’t include the correct authentication information.")
		panic("\nthe request isn’t authorized or doesn’t include the correct authentication information.\nHint: check the credentials in ~/.steampipe/config/weatherkit.spc")
	default:
		c.logger.Info("DoRequest", "status", status)
	}
}

func (c *Client) DoRequest(r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}

	resp, err := c.httpClient.Do(r)

	if resp == nil || err != nil {
		c.logger.Error("DoRequest", "message", "received empty response")
		return errors.New("an error occurred while doing the request")
	}

	c.checkResponseStatus(resp)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var buf bytes.Buffer
	dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))

	if err := dec.Decode(v); err != nil {
		c.logger.Error("DoRequest", "message", "could not parse response body")
		return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
	}

	return nil
}

func (c *Client) loadPrivateKey() *ecdsa.PrivateKey {
	// Read, decode, and parse the private key
	fileBytes, _ := os.ReadFile(*c.config.PrivateKeyPath)
	x509Encoded, _ := pem.Decode(fileBytes)
	parsedKey, _ := x509.ParsePKCS8PrivateKey(x509Encoded.Bytes)
	ecdsaPrivateKey, _ := parsedKey.(*ecdsa.PrivateKey)
	return ecdsaPrivateKey
}

func (c *Client) createJwt() string {

	// Define standard claims
	claims := jwt.StandardClaims{
		Issuer:    *c.config.TeamId,
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 5).UTC().Unix(),
		Subject:   *c.config.ServiceId,
	}

	// Create the JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Add header information
	token.Header = map[string]interface{}{
		"alg": "ES256",
		"kid": c.config.KeyId,
		"id":  claims.Issuer + "." + claims.Subject,
	}

	// Sign and get the complete encoded token as a string using the secret
	ecdsaPrivateKey := c.loadPrivateKey()
	tokenString, _ := token.SignedString(ecdsaPrivateKey)

	return tokenString
}

func (c *Client) Availability(ctx context.Context, latitude float64, longitude float64) ([]string, error) {
	lat := fmt.Sprintf("%f", latitude)
	lng := fmt.Sprintf("%f", longitude)
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "availability", lat, lng}, "/"),
	}
	u := requestUrl.Query()
	u.Set("country", "US")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var dataSet []string

	err := c.Get(ctx, requestUrl.String(), &dataSet)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return dataSet, nil
}

func (c *Client) Weather(ctx context.Context, latitude float64, longitude float64, datasets []string) (Weather, error) {
	lat := fmt.Sprintf("%f", latitude)
	lng := fmt.Sprintf("%f", longitude)
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, lat, lng}, "/"),
	}
	u := requestUrl.Query()
	u.Set("country", "US")
	u.Set("dataSets", strings.Join(datasets, ","))
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}

func (c *Client) CurrentWeather(ctx context.Context, latitude float64, longitude float64) (Weather, error) {
	return c.Weather(ctx, latitude, longitude, []string{"currentWeather"})
}

func (c *Client) DailyForecast(ctx context.Context, latitude float64, longitude float64) (Weather, error) {
	return c.Weather(ctx, latitude, longitude, []string{"forecastDaily"})
}

func (c *Client) HourlyForecast(ctx context.Context, latitude float64, longitude float64) (Weather, error) {
	return c.Weather(ctx, latitude, longitude, []string{"forecastHourly"})
}

func (c *Client) NextHourForecast(ctx context.Context, latitude float64, longitude float64) (Weather, error) {
	return c.Weather(ctx, latitude, longitude, []string{"forecastNextHour"})
}

func (c *Client) WeatherAlerts(ctx context.Context, latitude float64, longitude float64) (Weather, error) {
	return c.Weather(ctx, latitude, longitude, []string{"weatherAlerts"})
}
