/*
 * @Description   : Blue Planet
 * @Author        : serialt
 * @Email         : tserialt@gmail.com
 * @Created Time  : 2023-04-18 08:31:58
 * @Last modified : 2023-04-30 11:14:32
 * @FilePath      : /sugar/exmaple/main.go
 * @Other         :
 * @              :
 *
 *
 *
 */
package main

import (
	sugar "github.com/serialt/sugar/v3"
	"golang.org/x/exp/slog"
)

func init() {
	slog.SetDefault(sugar.New(
		sugar.WithLevel("debug"),
	))
}
func main() {

	slog.Debug("debug msg")
	slog.Info("info msg")
}
