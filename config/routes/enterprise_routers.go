// +build enterprise

package routes

import "github.com/sunwukonga/qor-scbn/config/admin"

func init() {
	Router()
	WildcardRouter.AddHandler(admin.MicroSite)
}
