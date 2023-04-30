/*
 * @Description   : Blue Planet
 * @Author        : serialt
 * @Email         : tserialt@gmail.com
 * @Created Time  : 2023-04-18 08:31:58
 * @Last modified : 2023-04-30 10:37:30
 * @FilePath      : /sugar/exmaple/main.go
 * @Other         :
 * @              :
 *
 *
 *
 */
package main

import (
	sugar "github.com/serialt/sugar/v2"
	"golang.org/x/exp/slog"
)

func init() {
	options := []sugar.LogOptions{
		sugar.WithLevel("debug"),
		sugar.WithShort(false),
		sugar.WithType("json"),
	}
	slog.SetDefault(sugar.New(options...))
}
func main() {

	slog.Debug("debug msg")
	slog.Info("info msg")
}
