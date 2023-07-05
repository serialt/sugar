/*
 * @Description   : Blue Planet
 * @Author        : serialt
 * @Email         : tserialt@gmail.com
 * @Created Time  : 2023-04-18 08:31:58
 * @Last modified : 2023-07-05 13:07:37
 * @FilePath      : /sugar/exmaple/main.go
 * @Other         :
 * @              :
 *
 *
 *
 */
package main

import (
	"log/slog"

	sugar "github.com/serialt/sugar/v2"
)

func init() {
	slog.SetDefault(sugar.New(
		sugar.WithLevel("debug"),
		sugar.WithShort(false),
		sugar.WithType("json")))
}
func main() {

	slog.Debug("debug msg")
	slog.Info("info msg")
}
