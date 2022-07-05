package upl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/backup"
	"github.com/xh3b4sd/tracer"
)

type run struct {
	aws *apicliaws.AWS
	bac backup.Interface
	ctx context.Context
	fla *fla
	log logger.Interface
}

func (r *run) run(cmd *cobra.Command, args []string) error {
	var err error

	{
		r.fla.Verify()
	}

	// --------------------------------------------------------------------- //

	{
		r.aws = apicliaws.Default()
	}

	{
		r.bac = redigo.Default().Backup()
	}

	{
		r.ctx = context.Background()
	}

	{
		r.log = logger.Default()
	}

	// --------------------------------------------------------------------- //

	{
		fmt.Printf("starting redis backup\n")
	}

	{
		err := r.bac.Create()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var fil *os.File
	{
		fil, err = os.Open(filepath.Join(r.fla.Dat, r.fla.Fil))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var inf os.FileInfo
	{
		inf, err = fil.Stat()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("uploading %s\n", r.siz(inf.Size()))
	}

	var buc string
	{
		buc = "xh3b4sd-phoebe-backup"
	}

	var pre string
	{
		pre = timfmt(time.Now())
	}

	{
		err := r.aws.Upload(buc, filepath.Join(pre, r.fla.Fil), fil)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("\nfinished redis backup\n")
	}

	return nil
}
