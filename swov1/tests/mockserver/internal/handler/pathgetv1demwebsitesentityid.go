// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathGetV1DemWebsitesEntityID(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "getWebsite[0]":
			dir.HandlerFunc("getWebsite", testGetWebsiteGetWebsite0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testGetWebsiteGetWebsite0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, false, "Bearer"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody := &components.GetWebsiteResponse{
		ID:     "e-1448474379026206720",
		Type:   "Website",
		Status: components.GetWebsiteResponseStatusUp,
		Name:   "solarwinds.com",
		URL:    "https://www.solarwinds.com",
		AvailabilityCheckSettings: &components.GetWebsiteResponseAvailabilityCheckSettings{
			PlatformOptions: &components.GetWebsiteResponsePlatformOptions{
				ProbePlatforms: []components.ProbePlatform{
					components.ProbePlatformAws,
				},
				TestFromAll: types.Bool(true),
			},
			TestFrom: components.TestFrom{
				Type: components.TypeRegion,
				Values: []string{
					"NA",
				},
			},
			TestIntervalInSeconds: 14400,
			OutageConfiguration: &components.GetWebsiteResponseOutageConfiguration{
				FailingTestLocations: components.GetWebsiteResponseFailingTestLocationsAll,
				ConsecutiveForDown:   2,
			},
			CheckForString: &components.GetWebsiteResponseCheckForString{
				Operator: components.GetWebsiteResponseOperatorContains,
				Value:    "string",
			},
			Protocols: []components.WebsiteProtocol{
				components.WebsiteProtocolHTTP,
				components.WebsiteProtocolHTTPS,
			},
			Ssl: &components.GetWebsiteResponseSsl{
				Enabled:                        types.Bool(true),
				DaysPriorToExpiration:          types.Int(7),
				IgnoreIntermediateCertificates: types.Bool(true),
			},
			CustomHeaders: []components.CustomHeaders{
				components.CustomHeaders{
					Name:  "string",
					Value: "string",
				},
			},
			AllowInsecureRenegotiation: types.Bool(true),
			PostData:                   types.String("{\"example\": \"value\"}"),
		},
		Tags: []components.Tag{
			components.Tag{
				Key:   "environment",
				Value: "production",
			},
		},
		Rum: &components.GetWebsiteResponseRum{
			ApdexTimeInSeconds: types.Int(4),
			Snippet:            types.String("string"),
			Spa:                true,
		},
		LastOutageStartTime:          types.MustNewTimeFromString("2025-01-15T14:31:19.735Z"),
		LastOutageEndTime:            types.MustNewTimeFromString("2025-01-15T14:31:19.735Z"),
		LastTestTime:                 types.MustNewTimeFromString("2025-01-15T14:31:19.735Z"),
		LastErrorTime:                types.MustNewTimeFromString("2025-01-15T14:31:19.735Z"),
		LastResponseTime:             types.Int(376),
		NextOnDemandAvailabilityTime: types.Int(0),
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}
