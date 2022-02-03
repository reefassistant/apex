package apex_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go.reefassistant.com/apex"
	v1alpha1 "go.reefassistant.com/apex/client/private/v1alpha1"
	"go.reefassistant.com/apex/test/simulator"
)

func TestAuth(t *testing.T) {
	ctrl := simulator.NewController("userx", "secret")
	srv := httptest.NewServer(ctrl)
	defer srv.Close()

	c, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	ctx := context.Background()

	// Not yet logged in, GetUser should return empty
	resp1, _, err := c.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "", resp1.Username)

	// Login attempt with valid credentials
	resp2, _, err := c.PrivateV1().DefaultApi.LoginUser(ctx).LoginRequest(v1alpha1.LoginRequest{
		Login:    "userx",
		Password: "secret",
	}).Execute()
	require.NoError(t, err)

	// Test new sid
	sid2 := resp2.ConnectSid
	require.Equal(t, 25, len(resp2.ConnectSid))

	// Validate we are who we think we are, getUser needs to return our username
	resp3, _, err := c.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "userx", resp3.Username)

	// Login again with valid credentials
	resp4, _, err := c.PrivateV1().DefaultApi.LoginUser(ctx).LoginRequest(v1alpha1.LoginRequest{
		Login:    "userx",
		Password: "secret",
	}).Execute()
	require.NoError(t, err)

	// Test new sid
	sid3 := resp4.ConnectSid
	require.Equal(t, 25, len(resp4.ConnectSid))
	require.NotEqual(t, sid2, sid3)

	// Invalidate previous sid and check we are still connected using newest sid
	ctrl.InvalidateSession(sid2)
	resp5, _, err := c.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "userx", resp5.Username)

	// Login with invalid credentials should result in HTTP 401
	_, _, err = c.PrivateV1().DefaultApi.LoginUser(ctx).LoginRequest(v1alpha1.LoginRequest{
		Login:    "userx",
		Password: "bobo",
	}).Execute()
	require.Contains(t, err.Error(), http.StatusText(http.StatusUnauthorized))

	// Verify our previous successfull authenticted sid is still valid
	resp6, _, err := c.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "userx", resp6.Username)

	// Destroy newest session and confirm we are no longer authenticted
	ctrl.InvalidateSession(sid3)
	resp7, _, err := c.PrivateV1().DefaultApi.GetUser(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, "", resp7.Username)
}
