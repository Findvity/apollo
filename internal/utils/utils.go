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
package utils

import (
	"fmt"

	"github.com/BRO3886/apollo/internal/styles"
	"github.com/ttacon/chalk"
)

// Warn prints a formatted and colored WARN message
func Warn(t string, b string) {
	fmt.Println(chalk.White.Color("apollo"), styles.WarnStyle.Style("WARN"), chalk.Magenta.Color(t), chalk.Yellow.Color(b))
}

// Info prints a formatted and colored INFO message
func Info(t string, b string) {
	fmt.Println(chalk.White.Color("apollo"), styles.InfoStyle.Style("INFO"), chalk.Magenta.Color(t), chalk.White.Color(b))
}

// Err prints a formatted and colored ERROR message
func Err(t string, b string) {
	fmt.Println(chalk.White.Color("apollo"), styles.ErrorStyle.Style("ERROR"), chalk.Magenta.Color(t), chalk.White.Color(b))
}
