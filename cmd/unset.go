/*
Copyright © 2026 Andreas Burg <andreas.burg@gmx.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/aburg/tjann/util"
	"github.com/spf13/cobra"
)

// unsetCmd represents the unset command
var unsetCmd = &cobra.Command{
	Use:   "unset <kind> <key>",
	Short: "Delete a value from the annotation store",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		from := args[0]
		key := args[1]

		// read data via timew export
		times, err := util.ReadTimes(from)
		if err != nil {
			return err
		}

		for _, time := range times {
			_, ok := time.Annotation[key]
			if ok {
				delete(time.Annotation, key)
				util.WriteTime(time)
			} else {
				return fmt.Errorf("key not found: %s", key)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}
