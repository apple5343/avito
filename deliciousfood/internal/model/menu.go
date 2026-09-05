package model

import (
	"github.com/shopspring/decimal"
)

var MenuItems = []MenuItem{
	{
		Name:        "Бургер классический",
		Description: "Сочная говяжья котлета, салат айсберг, помидоры, маринованные огурцы, фирменный соус",
		Price:       decimal.NewFromFloat(450.00),
		Category:    "Бургеры",
		IsAvailable: true,
	},
	{
		Name:        "Чизбургер двойной",
		Description: "Две сочные котлеты, двойной чеддер, бекон, фирменный соус",
		Price:       decimal.NewFromFloat(650.00),
		Category:    "Бургеры",
		IsAvailable: true,
	},
	{
		Name:        "Острый куриный бургер",
		Description: "Хрустящая куриная котлета в панировке, острый соус, халапеньо, салат коул-слоу",
		Price:       decimal.NewFromFloat(420.00),
		Category:    "Бургеры",
		IsAvailable: true,
	},
	{
		Name:        "Пицца Маргарита",
		Description: "Томатный соус, моцарелла, базилик, оливковое масло",
		Price:       decimal.NewFromFloat(590.00),
		Category:    "Пицца",
		IsAvailable: true,
	},
	{
		Name:        "Пицца Пепперони",
		Description: "Томатный соус, моцарелла, пепперони, орегано",
		Price:       decimal.NewFromFloat(690.00),
		Category:    "Пицца",
		IsAvailable: true,
	},
	{
		Name:        "Мясная пицца",
		Description: "Томатный соус, моцарелла, говядина, курица, бекон, колбаски",
		Price:       decimal.NewFromFloat(790.00),
		Category:    "Пицца",
		IsAvailable: true,
	},
	{
		Name:        "Картофель фри большой",
		Description: "Хрустящий картофель фри с морской солью",
		Price:       decimal.NewFromFloat(180.00),
		Category:    "Закуски",
		IsAvailable: true,
	},
	{
		Name:        "Куриные наггетсы (8 шт)",
		Description: "Сочные куриные кусочки в хрустящей панировке, соус на выбор",
		Price:       decimal.NewFromFloat(320.00),
		Category:    "Закуски",
		IsAvailable: true,
	},
	{
		Name:        "Луковые кольца",
		Description: "Хрустящие луковые кольца в темпуре с соусом BBQ",
		Price:       decimal.NewFromFloat(220.00),
		Category:    "Закуски",
		IsAvailable: true,
	},
	{
		Name:        "Coca-Cola 0.5л",
		Description: "Классическая Кока-Кола",
		Price:       decimal.NewFromFloat(120.00),
		Category:    "Напитки",
		IsAvailable: true,
	},
	{
		Name:        "Апельсиновый сок 0.3л",
		Description: "Свежевыжатый апельсиновый сок",
		Price:       decimal.NewFromFloat(180.00),
		Category:    "Напитки",
		IsAvailable: true,
	},
	{
		Name:        "Вода минеральная 0.5л",
		Description: "Минеральная вода без газа",
		Price:       decimal.NewFromFloat(80.00),
		Category:    "Напитки",
		IsAvailable: true,
	},
	{
		Name:        "Чизкейк Нью-Йорк",
		Description: "Классический американский чизкейк с ягодным соусом",
		Price:       decimal.NewFromFloat(280.00),
		Category:    "Десерты",
		IsAvailable: true,
	},
	{
		Name:        "Мороженое ассорти (3 шарика)",
		Description: "Ванильное, шоколадное, клубничное мороженое",
		Price:       decimal.NewFromFloat(220.00),
		Category:    "Десерты",
		IsAvailable: true,
	},
	{
		Name:        "Шоколадный брауни",
		Description: "Тёплый шоколадный брауни с шариком ванильного мороженого",
		Price:       decimal.NewFromFloat(250.00),
		Category:    "Десерты",
		IsAvailable: true,
	},
}

type MenuItem struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Category    string          `json:"category"`
	IsAvailable bool            `json:"is_available"`
}
