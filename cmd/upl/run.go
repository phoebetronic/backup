package upl

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/phoebetron/trades/typ/key"
	"github.com/phoebetron/trades/typ/trades"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/framer"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type run struct {
	client  *apicliaws.AWS
	flags   *flags
	frames  framer.Frames
	key     *key.Key
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.key = r.newkey()
	}

	{
		r.client = apicliaws.Default()
	}

	{
		r.storage = r.newsto()
	}

	{
		r.frames = r.newfra()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	var end time.Time
	{
		sta = r.frames.Min().Sta
		end = r.frames.Max().End
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("starting backup between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	tra := &trades.Trades{}
	{
		tra.EX = r.storage.Market().Exc()
		tra.AS = r.storage.Market().Ass()
		tra.ST = timestamppb.New(sta)
		tra.EN = timestamppb.New(end)
		tra.TR = r.tra()
	}

	var byt []byte
	{
		byt, err = proto.Marshal(tra)
		if err != nil {
			panic(err)
		}
	}

	var rea bytes.Reader
	{
		rea = *bytes.NewReader(byt)
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("buffered %s\n", r.siz(rea.Size()))
	}

	var buc string
	{
		buc = "xh3b4sd-phoebe-backup"
	}

	var pre string
	{
		pre = fmt.Sprintf("tra-raw.exc-%s.ass-%s", r.storage.Market().Exc(), r.storage.Market().Ass())
	}

	var suf string
	{
		suf = fmt.Sprintf("%s.pb.raw", bacfmt(sta))
	}

	{
		err := r.client.Upload(buc, filepath.Join(pre, suf), rea)
		if err != nil {
			panic(err)
		}
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("removing trades between %s and %s\n", scrfmt(sta), scrfmt(end))
	}

	{
		r.rem()
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("finished backup\n")
	}
}
