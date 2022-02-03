package apex_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"go.reefassistant.com/apex"
	v1 "go.reefassistant.com/apex/client/public/v1"
	"go.reefassistant.com/apex/test/util"
)

func TestV1BaseStatus(t *testing.T) {
	srv := util.NewMockServerFromFixture("public", "v1-status-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetStatus(context.Background())
	status, _, err := client.PublicV1().DefaultApi.GetStatusExecute(req)
	require.NoError(t, err)

	expected := v1.StatusResponse{
		Istat: v1.StatusContainer{
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
			Feed: &v1.StatusFeed{
				Name:   int32(1),
				Active: int32(2),
			},
			Power: &v1.StatusPower{
				Failed:   v1.PtrInt64(1635155614),
				Restored: v1.PtrInt64(1635155697),
			},
			Inputs: []v1.StatusInput{
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
			Outputs: []v1.StatusOutput{
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
					Intensity: v1.PtrFloat32(27),
					Name:      "Name_I2",
					Gid:       "",
					Type:      "type2|stuff",
					ID:        1,
					Did:       "did_2",
				},
			},
			Link: &v1.StatusLink{
				LinkState: 5,
				LinkKey:   "S0meK3Y",
				Link:      true,
			},
		},
	}
	require.Equal(t, expected, status)
}

func TestBaseDatalog(t *testing.T) {
	srv := util.NewMockServerFromFixture("public", "v1-datalog-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetDatalog(context.Background())
	datalog, _, err := client.PublicV1().DefaultApi.GetDatalogExecute(req)
	require.NoError(t, err)

	expected := v1.DatalogResponse{
		Ilog: v1.DatalogContainer{
			Hostname: "the_Hostname",
			Software: "the_Software",
			Hardware: "the_Hardware",
			Serial:   "the_Serial",
			Type:     "the_Type",
			Extra: &map[string]string{
				"key1": "foo,bar",
			},
			Timezone: "the_Timezone",
			Record: []v1.DatalogRecord{
				{
					Date: int64(1635663600),
					Data: []v1.DatalogData{
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
					Data: []v1.DatalogData{
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
	srv := util.NewMockServerFromFixture("public", "v1-outlog-base.json")
	defer srv.Close()

	client, err := apex.NewDefaultClient(srv.URL)
	require.NoError(t, err)

	req := client.PublicV1().DefaultApi.GetOutlog(context.Background())
	outlog, _, err := client.PublicV1().DefaultApi.GetOutlogExecute(req)
	require.NoError(t, err)

	expected := v1.OutlogResponse{
		Olog: v1.OutlogContainer{
			Hostname: "the_Hostname",
			Software: "the_Software",
			Hardware: "the_Hardware",
			Serial:   "the_Serial",
			Type:     "the_Type",
			Extra: &map[string]string{
				"key1": "foo,bar",
			},
			Timezone: "the_Timezone",
			Record: []v1.OutlogRecord{
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
