package integrations

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WebgrRepository struct {
	DB *gorm.DB
}

type Clientes struct {
	Cncliente uint   `json:"cncliente" gorm:"primaryKey"`
	Nmrazao   string `json:"nmrazao"`
	Nmnick    string `json:"nmnick"`
	Cdcnpj    string `json:"cdcnpj"`
}

type Produtos struct {
	Cnproduto uint   `json:"cnproduto" gorm:"primaryKey"`
	Descricao string `json:"descricao"`
}

type Peditems struct {
	Cnpeditem       uint     `json:"cnpeditem" gorm:"primaryKey"`
	Cnpedido        uint     `json:"cnpedido"`
	Qtd             float64  `json:"qtd"`
	Preco           float64  `json:"preco"`
	Vrdescontofinal float64  `json:"vrdescontofinal"`
	Cnproduto       uint     `json:"cnproduto"`
	Produto         Produtos `json:"produto" gorm:"foreignKey:Cnproduto;references:Cnproduto"`
}

type Pedidos struct {
	Cnempresa       uint       `json:"cnempresa"`
	Cnpedido        uint       `json:"cnpedido" gorm:"primaryKey"`
	Cdpedido        string     `json:"cdpedido"`
	Vrtotal         float64    `json:"vrtotal"`
	Vrdescontofinal float64    `json:"vrdescontofinal"`
	Cncliente       uint       `json:"cncliente"`
	Cliente         Clientes   `json:"cliente" gorm:"foreignKey:Cncliente;references:Cncliente"`
	Peditems        []Peditems `json:"peditems" gorm:"foreignKey:Cnpedido;references:Cnpedido"`
}

func (r *WebgrRepository) GetPedidos(context *fiber.Ctx) error {
	pedidos := []Pedidos{}

	err := r.DB.Preload("Cliente").Preload("Peditems.Produto").Where("pedidos.cnempresa = ?", 71).Find(&pedidos).Error

	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Não foi possível buscar os pedidos",
		})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Pedidos encontrados com sucesso",
		"data":    pedidos,
	})
}

func (r *WebgrRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/webgr/pedidos", r.GetPedidos)
}
