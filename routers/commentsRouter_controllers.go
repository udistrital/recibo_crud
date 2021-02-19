package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:EstadoReciboController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:PagoReciboController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:ReciboController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoPagoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"] = append(beego.GlobalControllerRouter["github.com/udistrital/recibo_crud/controllers:TipoReciboController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
