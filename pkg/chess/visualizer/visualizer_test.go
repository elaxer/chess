package visualizer

import (
	"bytes"
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/variant/standard/board"
)

func TestVisualizer_Visualize(t *testing.T) {
	type fields struct {
		Options Options
	}
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantWriter string
	}{
		{
			"no_options",
			fields{Options: Options{}},
			args{board.NewFactory().CreateFilled()},
			`r n b q k b n r 
p p p p p p p p 
. . . . . . . . 
. . . . . . . . 
. . . . . . . . 
. . . . . . . . 
P P P P P P P P 
R N B Q K B N R 

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Visualizer{
				Options: tt.fields.Options,
			}
			writer := &bytes.Buffer{}
			v.Visualize(tt.args.board, writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Visualizer.Visualize() = \n%v, want\n%v", gotWriter, tt.wantWriter)
			}
		})
	}
}
