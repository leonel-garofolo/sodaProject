package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leonel-garofolo/soda/app/enviroment"
	"github.com/leonel-garofolo/soda/app/model"
	gr "github.com/mikeshimura/goreport"
)

//format
var pageSize = "A4"
var unit = "mm"
var orientation = "P" //P: , L: landscape

//letter
var fontTitleSize = 10
var fontSize = 8

//cell sizes
var cellSizeY = 7.8

var VAR_DELIVERY_NAME = "deliveryName"
var VAR_DELIVERY_CODE = "deliveryRootCode"

var colNumberSize = 10.0
var colCellEmptySize = 12.0
var xColPosition = []float64{
	colNumberSize,    //cod
	colNumberSize,    //P/S
	colNumberSize,    //P/J
	40,               //Direccion
	12,               //Numero
	13,               //Deuda
	colCellEmptySize, //EmptyCell1
	colCellEmptySize, //EmptyCell2
	colCellEmptySize, //EmptyCell3
	colCellEmptySize, //EmptyCell4
	colCellEmptySize, //EmptyCell5
	colCellEmptySize, //EmptyCell6
	colCellEmptySize, //EmptyCell7
	colCellEmptySize, //EmptyCell8
}

var headersTitle = []string{
	"Cod",
	"P/S",
	"P/J",
	"Dirección",
	"Numero",
	"Deuda",
}

type Report struct {
	Context *enviroment.Context
}

func (r *Report) GenerateReport(c *fiber.Ctx) error {
	idRoot := c.Query("id", "-1")
	log.Println("generateReport root-> ", idRoot)

	goReport := gr.CreateGoReport()

	goReport.PageTotal = true
	goReport.SumWork["amountcum="] = 0.0
	font1 := gr.FontMap{
		FontName: "IPAexG",
		FileName: "ttf//ipaexg.ttf",
	}
	fonts := []*gr.FontMap{&font1}
	goReport.SetFonts(fonts)
	d := new(S1Detail)
	clients, err := r.getClientsRoot(idRoot)
	if err != nil {
		log.Panic(err)
	}

	deliveryName, deliveryCode, errDelivery := r.getNameCodeRoot(idRoot)
	if errDelivery != nil {
		log.Panic(errDelivery)
	}
	varGlobales := make(map[string]string)
	varGlobales[VAR_DELIVERY_NAME] = deliveryName
	varGlobales[VAR_DELIVERY_CODE] = deliveryCode
	goReport.Vars = varGlobales

	goReport.RegisterBand(gr.Band(*d), gr.Detail)
	h := new(S1Header)
	goReport.RegisterBand(gr.Band(*h), gr.PageHeader)
	s := new(S1Summary)
	goReport.RegisterBand(gr.Band(*s), gr.Summary)
	goReport.Records = clients
	goReport.SetPage(pageSize, unit, orientation)
	goReport.SetFooterY(400)
	goReport.Execute("simple1.pdf")

	return c.JSON(idRoot)
}

type S1Detail struct {
}

func (h S1Detail) GetHeight(report gr.GoReport) float64 {
	return 8
}
func (h S1Detail) Execute(report gr.GoReport) {
	cols := report.Records[report.DataPos].([]string)
	report.Font("IPAexG", fontSize, "")

	x := 5.0
	y := 2.0
	yIndex := 0
	report.LineType("straight", 0.3)
	for i := 0; i < len(xColPosition); i++ {
		if i < len(cols)-1 {
			report.Cell(x+1, y, cols[yIndex])
			yIndex = yIndex + 1
		} else {
			report.Cell(x, y, "")
		}
		report.LineV(x, 0, cellSizeY)
		x = x + xColPosition[i]
	}
	report.LineH(5.0, cellSizeY, x)
	report.LineV(x, 0, cellSizeY)
	amt := ParseFloatNoError(cols[yIndex])
	report.SumWork["amountcum="] += amt
}

func ParseFloatNoError(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

type S1Header struct {
}

func (h S1Header) GetHeight(report gr.GoReport) float64 {
	return 30
}
func (h S1Header) Execute(report gr.GoReport) {
	title := report.Vars[VAR_DELIVERY_NAME] + " Nro. " + report.Vars[VAR_DELIVERY_CODE]

	titleLineY := 15.0
	report.Font("IPAexG", fontTitleSize, "")
	report.Cell(10, titleLineY, title)
	report.Cell(100, titleLineY, time.Now().Format("02/01/2006"))
	report.Cell(180, titleLineY, "pag.")
	report.CellRight(187, titleLineY, 5, strconv.Itoa(report.Page))
	report.Cell(194, 15, "/")
	report.CellRight(196, titleLineY, 3, "{#TotalPage#}")

	report.Font("IPAexG", fontSize, "")
	x := 5.0
	y := 23.0
	xIndex := 0
	for i := 0; i < len(xColPosition); i++ {
		if i < len(headersTitle) {
			report.Cell(x, y, headersTitle[i])
		}
		x = x + xColPosition[xIndex]
		xIndex = xIndex + 1
	}
	report.LineH(5.0, y+6.8, x)
}

type S1Summary struct {
}

func (h S1Summary) GetHeight(report gr.GoReport) float64 {
	return 2
}
func (h S1Summary) Execute(report gr.GoReport) {
	report.Cell(10, 2, "Deuda Total:")
	report.CellRight(30, 2, 30, strconv.FormatFloat(
		report.SumWork["amountcum="], 'f', 2, 64))
}

func (r *Report) getClientsRoot(idRoot string) ([]interface{}, error) {
	idRootInt, errorParser := strconv.Atoi(idRoot)
	if errorParser != nil {
		panic(errorParser.Error())
	}

	db := r.Context.Database.Connection
	rows, err := db.Query("select c.num_order, c.address, c.`number`, c.price_per_soda, c.price_per_box, debt from client c where c.id_root= ? order by c.num_order asc", idRootInt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var clients []interface{}
	for rows.Next() {
		var client model.Client
		if err := rows.Scan(
			&client.Order,
			&client.Address,
			&client.NumAddress,
			&client.PricePerSoda,
			&client.PricePerBox,
			&client.Debt,
		); err != nil {
			return nil, err
		}

		clients = append(clients, []string{
			strconv.Itoa(client.Order),
			fmt.Sprintf("%.2f", client.PricePerSoda),
			fmt.Sprintf("%.2f", client.PricePerBox),
			client.Address,
			strconv.Itoa(client.NumAddress),
			fmt.Sprintf("%.2f", client.Debt),
		})
	}
	return clients, nil
}

func (r *Report) getNameCodeRoot(idRoot string) (string, string, error) {
	idRootInt, errorParser := strconv.Atoi(idRoot)
	if errorParser != nil {
		panic(errorParser.Error())
	}

	db := r.Context.Database.Connection
	rows, err := db.Query("select d.name, dr.code from soda.delivery_root dr inner join soda.delivery d on d.id_delivery = dr.id_delivery where dr.id_root= ?", idRootInt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var name string
	var code string
	if rows.Next() {
		var codeInt int
		if err := rows.Scan(
			&name,
			&codeInt,
		); err != nil {
			return "", "", err
		}
		code = strconv.Itoa(codeInt)
	}
	return name, code, nil
}
