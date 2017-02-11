// +build enterprise

package migrations

import "github.com/sunwukonga/qor-scbn/config/admin"

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
