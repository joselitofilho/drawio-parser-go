package xml

import (
	"encoding/xml"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseXML(t *testing.T) {
	type args struct {
		fileName string
	}

	tests := []struct {
		name      string
		args      args
		want      *MxFile
		targetErr error
	}{
		{
			name: "happy path",
			args: args{
				fileName: "testdata/diagram.xml",
			},
			want: &MxFile{
				XMLName: xml.Name{Local: "mxfile"},
				Diagram: Diagram{
					XMLName: xml.Name{Local: "diagram"},
					MxGraphModel: MxGraphModel{
						XMLName: xml.Name{Local: "mxGraphModel"},
						Root: Root{
							XMLName: xml.Name{Local: "root"},
							MxCells: []MxCell{{
								XMLName: xml.Name{Local: "mxCell"},
								ID:      "kVijt7gfVD9ZtySMmpSK-1",
								Value:   "myReceiver",
								Style: "outlineConnect=0;dashed=0;verticalLabelPosition=bottom;verticalAlign=top;" +
									"align=center;html=1;shape=mxgraph.aws3.lambda;fillColor=#F58534;" +
									"gradientColor=none;",
								Parent: "1",
								Vertex: "1",
								Geometry: &Geometry{
									XMLName: xml.Name{Local: "mxGeometry"},
									X:       850,
									Y:       -1121.5,
									Width:   76.5,
									Height:  93,
									As:      "geometry",
								},
							}},
						},
					},
				},
			},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.args.fileName)

			if tc.targetErr == nil {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.ErrorIs(t, err, tc.targetErr)
			}
		})
	}
}

func TestParseXML_OsOpenErrors(t *testing.T) {
	type args struct {
		fileName string
	}

	errDummy := errors.New("dummy error")

	tests := []struct {
		name      string
		args      args
		setup     func() (tearDown func())
		targetErr error
	}{
		{
			name: "when os.Open fails should return an error",
			args: args{
				fileName: "notfound",
			},
			setup: func() (tearDown func()) {
				osOpen = func(name string) (*os.File, error) {
					require.Equal(t, name, "notfound")
					return nil, errDummy
				}

				return func() {
					osOpen = os.Open
				}
			},
			targetErr: errDummy,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			tearDown := tc.setup()
			defer tearDown()

			got, err := Parse(tc.args.fileName)

			require.Nil(t, got)
			require.ErrorIs(t, err, tc.targetErr)
		})
	}
}

func TestParseXML_XMLDecodeErrors(t *testing.T) {
	type args struct {
		fileName string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when xml.Decode fails should return an error",
			args: args{
				fileName: "testdata/invalid.xml",
			},
			wantErr: true,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.args.fileName)

			require.Nil(t, got)
			require.Error(t, err)
		})
	}
}
