
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["management/controllers:UserController"] = append(beego.GlobalControllerRouter["management/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
