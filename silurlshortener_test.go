package silurlshortener

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jarcoal/httpmock"
	"github.com/savannahghi/serverutils"
)

func TestURLShortener_ShortenURL(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *ShortenURLPayload
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: shorten URL",
			args: args{
				ctx: context.Background(),
				payload: &ShortenURLPayload{
					LongURL:         gofakeit.URL(),
					Tags:            []string{"clinical"},
					Domain:          serverutils.MustGetEnvVar("URL_SHORTENER_DOMAIN"),
					ShortCodeLength: 5,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to shorten URL",
			args: args{
				ctx: context.Background(),
				payload: &ShortenURLPayload{
					LongURL:         gofakeit.URL(),
					Tags:            []string{"clinical"},
					Domain:          serverutils.MustGetEnvVar("URL_SHORTENER_DOMAIN"),
					ShortCodeLength: 5,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL := fmt.Sprintf("https://%s/rest/v3", serverutils.MustGetEnvVar("URL_SHORTENER_DOMAIN"))

			if tt.name == "Happy case: shorten URL" {
				httpmock.RegisterResponder(http.MethodPost, baseURL+"/short-urls", func(r *http.Request) (*http.Response, error) {
					resp := &ShortenURLResponse{
						ShortCode:   "123",
						ShortURL:    "http://example/short-urls",
						LongURL:     tt.args.payload.LongURL,
						DateCreated: time.Time{},
						Crawlable:   false,
					}

					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			}
			if tt.name == "Sad case: unable to shorten URL" {
				httpmock.RegisterResponder(http.MethodPost, baseURL+"/short-urls", func(r *http.Request) (*http.Response, error) {
					resp := &ShortenURLResponse{
						ShortCode:   "123",
						ShortURL:    "http://example/short-urls",
						LongURL:     tt.args.payload.LongURL,
						DateCreated: time.Time{},
						Crawlable:   false,
					}

					return httpmock.NewJsonResponse(http.StatusBadRequest, resp)
				})
			}

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			cl, err := NewURLShortener()
			if err != nil {
				t.Errorf("unable to initialize sdk: %s", err)
				return
			}

			_, err = cl.ShortenURL(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortener.ShortenURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
