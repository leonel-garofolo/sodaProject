package reports

import (
	"log"
	"strconv"
	"time"

	"github.com/leonel-garofolo/soda/app/repositories"
	gr "github.com/mikeshimura/goreport"
)

//cell sizes
var cellSizeY = 7.8

//Lines cell
const (
	//format
	pageSize    = "A4"
	unit        = "mm"
	orientation = "P" //P: , L: landscape

	//letter
	fontTitleSize    = 14
	fontSize         = 10
	xLineBold        = 2.0
	colNumberSize    = 10.0
	colCellEmptySize = 13.0

	VAR_DELIVERY_NAME = "deliveryName"
	VAR_DELIVERY_CODE = "deliveryRootCode"
)

var xColPosition = []float64{
	colNumberSize,    //cod
	12,               //P/S
	12,               //P/J
	40,               //Direccion
	12,               //Numero
	15,               //Deuda
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
	"Num.",
	"Deuda",
}

type Reports struct {
	Dao *repositories.Dao
}

func New(report *Reports) *Reports {
	return report
}

func (r *Reports) BuildSheetRoot(idRoot string) {
	goReport := gr.CreateGoReport()

	goReport.PageTotal = true
	goReport.SumWork["amountcum="] = 0.0
	font1 := gr.FontMap{
		FontName: "IPAexG",
		FileName: ".//app//reports//ttf//ipaexg.ttf",
	}
	fonts := []*gr.FontMap{&font1}
	goReport.SetFonts(fonts)
	d := new(S1Detail)
	clients, err := r.Dao.GetClientsRoot(idRoot)
	if err != nil {
		log.Panic(err)
	}

	deliveryName, deliveryCode, errDelivery := r.Dao.GetNameCodeRoot(idRoot)
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
	goReport.SetFooterY(295)
	goReport.Execute("simple1.pdf")
}

type S1Header struct{}

func (h S1Header) GetHeight(report gr.GoReport) float64 {
	return 20
}
func (h S1Header) Execute(report gr.GoReport) {
	title := report.Vars[VAR_DELIVERY_NAME] + " Nro. " + report.Vars[VAR_DELIVERY_CODE]

	titleLineY := 7.0
	report.Font("IPAexG", fontTitleSize, "")
	report.Cell(10, titleLineY, title)
	report.Cell(100, titleLineY, time.Now().Format("02/01/2006"))
	report.Cell(178, titleLineY, "pag.")
	report.CellRight(189, titleLineY, 5, strconv.Itoa(report.Page))
	report.Cell(194, titleLineY, "/")
	report.CellRight(198, titleLineY, 3, "{#TotalPage#}")

	report.Font("IPAexG", fontSize, "")
	x := xLineBold
	headerColTitle := 15.0
	xIndex := 0
	for i := 0; i < len(xColPosition); i++ {
		if i < len(headersTitle) {
			report.Cell(x, headerColTitle, headersTitle[i])
		}
		x = x + xColPosition[xIndex]
		xIndex = xIndex + 1
	}
	report.LineH(xLineBold, headerColTitle+5, x-2.6)
}

type S1Detail struct {
}

func (h S1Detail) GetHeight(report gr.GoReport) float64 {
	return 8
}
func (h S1Detail) Execute(report gr.GoReport) {
	cols := report.Records[report.DataPos].([]string)
	report.Font("IPAexG", fontSize, "")

	x := 2.0
	y := 2.0
	yIndex := 0
	report.LineType("straight", 0.3)
	for i := 0; i < len(xColPosition); i++ {
		if i < len(cols) {
			if cols[i] != "0.00" {
				if i == 3 {
					report.Cell(x+2, y, cols[i])
				} else {
					report.CellRight(x-0.5, y, xColPosition[i], cols[i])
				}
			}
			yIndex = yIndex + 1
		} else {
			report.Cell(x, y, "")
		}
		report.LineV(x, 0, cellSizeY)
		x = x + xColPosition[i]
	}
	report.LineH(xLineBold, cellSizeY, x)
	report.LineV(x, 0, cellSizeY)
	amt := ParseFloatNoError(cols[yIndex-1])
	report.SumWork["amountcum="] += amt
}

func ParseFloatNoError(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
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
