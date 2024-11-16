package util

import (
	"fmt"
	"hankquan.top/ecron/pkg/crontab"
	"log"
	"os"
	"text/tabwriter"
)

var header = "INDEX\tCRON_EXPR\tCMD\tNEXT_SCHEDULED\tSTATE"

func PrintTable(cronJobRows []crontab.CronJobRow) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	defer func(w *tabwriter.Writer) {
		err := w.Flush()
		if err != nil {
			log.Fatal("Failed to flush TabWriter", err)
		}
	}(w)
	_, _ = fmt.Fprintln(w, header)
	for _, cronLine := range cronJobRows {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			cronLine.Index, cronLine.Cron, cronLine.Cmd, cronLine.Next, cronLine.State)
	}
}
