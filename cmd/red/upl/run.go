package upl

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/backup"
	"github.com/xh3b4sd/tracer"

	"github.com/phoebetron/backup/pkg/cli/apicliaws"
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
		r.log.Log(r.ctx, "level", "info", "message", "creating redis backup")
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
		r.log.Log(r.ctx, "level", "info", "message", "uploading redis backup", "siz", r.siz(inf.Size()))
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

	return nil
}
