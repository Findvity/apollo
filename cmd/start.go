/*
Copyright © 2020 Siddhartha Varma & Vasudha Tapriya

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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/BRO3886/apollo/internal/styles"
	"github.com/manifoldco/promptui"

	"github.com/BRO3886/apollo/internal/utils"
	"github.com/ttacon/chalk"

	"github.com/BRO3886/apollo/pkg/apollo"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start execution of commands",
	Long: `
	use this command to begin execution step by step
	`,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := ioutil.ReadFile("/home/sidv/go_projs/cli_projs/ge_healthcare/data/sample.json")
		if err != nil {
			utils.Err("fatal", "unable to read data")
			os.Exit(1)
		}
		cmds, err := apollo.UnmarshalCommands(d)
		if err != nil {
			utils.Err("decode", "unable to decode data")
			os.Exit(1)
		}

		utils.Warn("sill", "starting execution")
		for _, c := range cmds {
			utils.Info("sill", fmt.Sprintf("%s %s", "now executing:", chalk.Yellow.Color("Step "+c.ID)))
			fmt.Println(strings.Repeat("-", len(c.Info)))
			fmt.Println(c.Info)
			fmt.Println(strings.Repeat("-", len(c.Info)))
			utils.Info("cmd", fmt.Sprintf("%s %s", "command:", chalk.Green.Color(c.Input)))
			l := strings.Split(c.Input, " ")

			if l[0] == "vim" || l[0] == "nano" {
				utils.Info("sill", "execute the above command in a new terminal instance")
				utils.Info("sill", "continue after following the above steps and saving the file")
			} else {

				cmdPath, err := exec.LookPath(l[0])
				if err != nil {
					utils.Err("notsup", "could not find "+l[0])
					os.Exit(1)
				}
				sub := l[1:]
				sub = append([]string{""}, sub...)
				cmd := exec.Cmd{Path: cmdPath, Args: sub}

				var outb, errb bytes.Buffer
				cmd.Stdout = &outb
				cmd.Stderr = &errb

				err = cmd.Run()
				if err != nil {
					utils.Err("fatal", err.Error())
					os.Exit(1)
				}
				fmt.Println()
				fmt.Println(styles.InfoStyle.Style("OUTPUT"))
				fmt.Println(outb.String())
			}
			fmt.Println(styles.InfoStyle.Style("EXPECTED OUTPUT"))
			fmt.Println(strings.Repeat("-", len(c.Info)))
			fmt.Println(strings.Join(c.Expected, "\n"))
			fmt.Println(strings.Repeat("-", len(c.Info)))

			prompt := promptui.Select{
				Label: "Is the output same as expected output?",
				Items: []string{"Yes", "No", "I don't know"},
			}

			_, result, err := prompt.Run()

			if err != nil {
				utils.Err("prompt", fmt.Sprintf("Prompt failed: %v", err))
				return
			}

			fmt.Printf("You choose %q. Continuing...\n", result)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
