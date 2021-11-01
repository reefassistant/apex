package apex_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"go.reefassistant.com/apex"
	v1public "go.reefassistant.com/apex/client/public/v1"
)

func newMockServerFromFixture(dir, file string) *httptest.Server {
	data, _ := ioutil.ReadFile(filepath.Join("test", dir, "fixtures", file))
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.Write(data)
		}),
	)
}

func TestHTTPClientTimeout(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte("foobar"))
		}),
	)
	defer srv.Close()

	client, err := apex.NewClient(srv.URL, apex.ClientConfig{
		HTTPClient: &http.Client{
			Timeout: 50 * time.Millisecond,
		},
	})
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetStatus(context.Background())
	_, _, err = client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}

func TestContextClientTimeout(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.Write([]byte("foobar"))
		}),
	)
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req := client.PublicV1().DefaultApi.GetStatus(ctx)
	_, _, err = client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.Contains(t, err.Error(), context.DeadlineExceeded.Error())
}
func TestV1BaseStatus(t *testing.T) {
	srv := newMockServerFromFixture("public", "v1-status-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetStatus(context.Background())
	status, _, err := client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.NoError(t, err)

	expected := v1public.StatusResponse{
		Istat: v1public.StatusContainer{
			Hostname: "the_Hostname",
			Software: "the_Software",
			Hardware: "the_Hardware",
			Serial:   "the_Serial",
			Type:     "the_Type",
			Extra: &map[string]string{
				"key1": "foo,bar",
			},
			Timezone: "the_Timezone",
			Date:     int64(1635672222),
			Feed: &v1public.StatusFeed{
				Name:   int32(1),
				Active: int32(2),
			},
			Power: &v1public.StatusPower{
				Failed:   v1public.PtrInt64(1635155614),
				Restored: v1public.PtrInt64(1635155697),
			},
			Inputs: []v1public.StatusInput{
				{
					Did:   "input1",
					Type:  "foo",
					Name:  "bar",
					Value: float32(1.2345),
				},
				{
					Did:   "INPUT_2",
					Type:  "Stuff",
					Name:  "More_STUFF",
					Value: float32(0),
				},
				{
					Did:   "3_P0",
					Type:  "foz",
					Name:  "baz",
					Value: float32(0),
				},
				{
					Did:   "4_P1",
					Type:  "fooz",
					Name:  "baaz",
					Value: float32(123),
				},
			},
			Outputs: []v1public.StatusOutput{
				{
					Status: []string{
						"",
						"OK",
						"89.8",
						"0.05",
						"197",
						"-339",
					},
					Name: "Name_I1",
					Gid:  "gid1",
					Type: "type1",
					ID:   0,
					Did:  "did_1",
				},
				{
					Status: []string{
						"FOO",
					},
					Intensity: v1public.PtrFloat32(27),
					Name:      "Name_I2",
					Gid:       "",
					Type:      "type2|stuff",
					ID:        1,
					Did:       "did_2",
				},
			},
			Link: &v1public.StatusLink{
				LinkState: 5,
				LinkKey:   "S0meK3Y",
				Link:      true,
			},
		},
	}
	require.Equal(t, expected, status)
}

func TestBaseDatalog(t *testing.T) {
	srv := newMockServerFromFixture("public", "v1-datalog-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetDatalog(context.Background())
	datalog, _, err := client.PublicV1().DefaultApi.GetDatalogExecute(req)
	require.NoError(t, err)

	expected := v1public.DatalogResponse{
		Ilog: v1public.DatalogContainer{
			Hostname: "the_Hostname",
			Software: "the_Software",
			Hardware: "the_Hardware",
			Serial:   "the_Serial",
			Type:     "the_Type",
			Extra: &map[string]string{
				"key1": "foo,bar",
			},
			Timezone: "the_Timezone",
			Record: []v1public.DatalogRecord{
				{
					Date: int64(1635663600),
					Data: []v1public.DatalogData{
						{
							Name:  "Tmp",
							Did:   "base_Temp",
							Type:  "Temp",
							Value: "79.0",
						},
						{
							Name:  "pH",
							Did:   "base_pH",
							Type:  "pH",
							Value: "8.23",
						},
					},
				},
				{
					Date: int64(1635663660),
					Data: []v1public.DatalogData{
						{
							Name:  "Tmp",
							Did:   "base_Temp",
							Type:  "Temp",
							Value: "79.5",
						},
						{
							Name:  "pH",
							Did:   "base_pH",
							Type:  "pH",
							Value: "8.25",
						},
					},
				},
			},
		},
	}
	require.Equal(t, expected, datalog)
}

func TestBaseOutlog(t *testing.T) {
	srv := newMockServerFromFixture("public", "v1-outlog-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetOutlog(context.Background())
	outlog, _, err := client.PublicV1().DefaultApi.GetOutlogExecute(req)
	require.NoError(t, err)

	expected := v1public.OutlogResponse{
		Olog: v1public.OutlogContainer{
			Hostname: "the_Hostname",
			Software: "the_Software",
			Hardware: "the_Hardware",
			Serial:   "the_Serial",
			Type:     "the_Type",
			Extra: &map[string]string{
				"key1": "foo,bar",
			},
			Timezone: "the_Timezone",
			Record: []v1public.OutlogRecord{
				{
					Date:  int64(1635666329),
					Did:   "1_1",
					Name:  "OUTLET_FOO",
					Value: "OFF",
				},
				{
					Date:  int64(1635666779),
					Did:   "2_1",
					Name:  "OUTLET_BAR",
					Value: "ON",
				},
			},
		},
	}
	require.Equal(t, expected, outlog)
}
