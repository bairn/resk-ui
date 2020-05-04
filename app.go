package resk

import (
	_ "github.com/bairn/account/core/accounts"
	"github.com/bairn/infra"
	"github.com/bairn/infra/base"
	_ "github.com/bairn/resk/core/envelopes"
	_ "resk-ui/views"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
