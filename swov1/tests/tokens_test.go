// THESE TESTS ARE MANUALLY CREATED TO TEST CREATION OF TOKENS VIA THE SDK.

package tests

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/solarwinds/swo-sdk-go/swov1/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSDK_CreateTokens(t *testing.T) {
	ctx := context.Background()
	s := CreateTestClient("token-test")

	createRes, err := s.Tokens.CreateToken(ctx, components.TokensCreateTokenRequest{
		Name: "swo-sdk-e2e-test-token",
		Tags: components.Tags{
			Server:          "swo-sdk-e2e-test-server",
			TagWithoutValue: "swo-sdk-e2e-test-tag",
		},
		Type: components.TokensCreateTokenRequestTypeIngestion,
	})
	require.NoError(t, err)
	assert.Equal(t, 201, createRes.HTTPMeta.Response.StatusCode)
	assert.NotEmpty(t, createRes.TokensCreateTokenResponse.Token)

	// --- CLEANUP: Delete the created token via internal endpoint ---

	token := createRes.TokensCreateTokenResponse.Token
	if token == "" {
		t.Log("Skipping cleanup: token is empty")
		return
	}
	sig := token[len(token)-43:]

	// Use SWO_STAGE_API_TOKEN for cleanup as well
	cleanupJWT := os.Getenv("SWO_STAGE_API_TOKEN")
	if cleanupJWT == "" {
		t.Log("Skipping token cleanup: SWO_STAGE_API_TOKEN not set")
		return
	}
	baseURL := os.Getenv("PUBLIC_SWO_API_STAGE_URL")
	if baseURL == "" {
		t.Log("Skipping token cleanup: PUBLIC_SWO_API_STAGE_URL not set")
		return
	}
	// Ensure no trailing slash
	baseURL = strings.TrimRight(baseURL, "/")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/tokens/%s", baseURL, sig), nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cleanupJWT))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Failed to cleanup token %s: %v", sig, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Logf("Cleanup token DELETE returned status %d", resp.StatusCode)
	}
}
