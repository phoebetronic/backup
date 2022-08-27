package upl

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
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
	storage trades.Storage
}

func (r *run) run(cmd *cobra.Command, args []string) {
	var err error

	{
		r.flags.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.client = apicliaws.New()
		r.frames = r.newfra()
		r.storage = r.newsto()
	}

	// --------------------------------------------------------------------- //

	var sta time.Time
	{
		sta = r.flags.Time
	}

	var end time.Time
	{
		end = sta.AddDate(0, 1, 0)
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
		tra.TR = r.alltra()
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

	for _, h := range r.frames {
		err := r.storage.Delete(h.Sta)
		if err != nil {
			panic(err)
		}
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("finished backup\n")
	}
}
