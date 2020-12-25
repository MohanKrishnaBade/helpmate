package cmd

import (
	"github.com/helpmate/src"
	"github.com/spf13/cobra"
)

// logParserCmd represents the logParser command
var (
	filerName, queryString string
	limit                  int64
	show                   bool
)
var logParserCmd = &cobra.Command{
	Use:   "logParser",
	Short: "command helps to parse the log file.",
	Run: func(cmd *cobra.Command, args []string) {
		src.LogParser(
			filerName,
			queryString,
			limit,
			show,
		)
	},
}

func init() {
	rootCmd.AddCommand(logParserCmd)

	// flags
	logParserCmd.PersistentFlags().StringVarP(&filerName, "filepath", "f", "./test.log", "logger filepath")
	logParserCmd.PersistentFlags().StringVarP(&queryString, "querystring", "q", "", "search term to retrieve the data from a log file")
	logParserCmd.PersistentFlags().Int64VarP(&limit, "limit", "l", 0, `* positive limit will read from the beginning of the file and limit the results
* 0 will read from the beginning of the file but there is not limit to stdout the results
* negative will read from the end of the file  and limit the results`)
	logParserCmd.PersistentFlags().BoolVarP(&show, "show", "s", true, "stdout the results if the flag is true.")

}
