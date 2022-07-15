package weatherkit

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	config     *weatherKitConfig
}

func NewClient(httpClient *http.Client, config weatherKitConfig) *Client {
	return &Client{
		httpClient: httpClient,
		config:     &config,
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
	token := c.createJwt(*c.config)
	req.Header.Add("Authorization", "Bearer "+token)

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) DoRequest(r *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if v == nil {
		return nil
	}

	var buf bytes.Buffer
	dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))

	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
	}

	return nil
}

func (c *Client) loadPrivateKey(keyPath string) *ecdsa.PrivateKey {
	// Read, decode, and parse the private key
	fileBytes, _ := ioutil.ReadFile(keyPath)
	x509Encoded, _ := pem.Decode(fileBytes)
	parsedKey, _ := x509.ParsePKCS8PrivateKey(x509Encoded.Bytes)
	ecdsaPrivateKey, _ := parsedKey.(*ecdsa.PrivateKey)
	return ecdsaPrivateKey
}

func (c *Client) createJwt(config weatherKitConfig) string {

	// Define standard claims
	claims := jwt.StandardClaims{
		Issuer:    *config.TeamId,
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 5).UTC().Unix(),
		Subject:   *config.ServiceId,
	}

	// Create the JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Add header information
	token.Header = map[string]interface{}{
		"alg": "ES256",
		"kid": config.KeyId,
		"id":  claims.Issuer + "." + claims.Subject,
	}

	// Sign and get the complete encoded token as a string using the secret
	ecdsaPrivateKey := c.loadPrivateKey(*config.PrivateKeyPath)
	tokenString, _ := token.SignedString(ecdsaPrivateKey)

	return tokenString
}

func (c *Client) Availability(ctx context.Context, latitude string, longitude string) ([]string, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "availability", latitude, longitude}, "/"),
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

func (c *Client) CurrentWeather(ctx context.Context, latitude string, longitude string) (Weather, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, latitude, longitude}, "/"),
	}
	u := requestUrl.Query()
	u.Set("dataSets", "currentWeather")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}

func (c *Client) DailyForecast(ctx context.Context, latitude string, longitude string) (Weather, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, latitude, longitude}, "/"),
	}
	u := requestUrl.Query()
	u.Set("dataSets", "forecastDaily")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}

func (c *Client) HourlyForecast(ctx context.Context, latitude string, longitude string) (Weather, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, latitude, longitude}, "/"),
	}
	u := requestUrl.Query()
	u.Set("dataSets", "forecastHourly")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}

func (c *Client) NextHourForecast(ctx context.Context, latitude string, longitude string) (Weather, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, latitude, longitude}, "/"),
	}
	u := requestUrl.Query()
	u.Set("dataSets", "forecastNextHour")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}

func (c *Client) WeatherAlerts(ctx context.Context, latitude string, longitude string) (Weather, error) {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   baseUrl,
		Path:   strings.Join([]string{"api", "v1", "weather", language, latitude, longitude}, "/"),
	}
	u := requestUrl.Query()
	u.Set("country", "US")
	u.Set("dataSets", "weatherAlerts")
	requestUrl.RawQuery = u.Encode()

	//Response object
	var weather Weather

	err := c.Get(ctx, requestUrl.String(), &weather)
	if err != nil {
		log.Fatalf("Request failed: [%s]", err)
	}
	return weather, nil
}
