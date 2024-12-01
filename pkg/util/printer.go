package util

import (
	"fmt"
	"hankquan.top/ecron/pkg/store/crontab"
	"log"
	"os"
	"text/tabwriter"
)

var header = "INDEX\tCRON_EXPR\tCMD\tNEXT_SCHEDULED\tSTATE"

func PrintTable(cronEntries []crontab.CronEntry) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	defer func(w *tabwriter.Writer) {
		err := w.Flush()
		if err != nil {
			log.Fatal("Failed to flush TabWriter", err)
		}
	}(w)
	_, _ = fmt.Fprintln(w, header)
	for _, cronEntry := range cronEntries {
		_, _ = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			cronEntry.Index, cronEntry.Cron, cronEntry.Cmd, cronEntry.Next, cronEntry.State)
	}
}
