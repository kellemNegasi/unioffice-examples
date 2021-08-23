// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/spreadsheet"
	"github.com/unidoc/unioffice/spreadsheet/formula"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Currently support", len(formula.SupportedFunctions()), "functions")
	fmt.Println(formula.SupportedFunctions())
	ss := spreadsheet.New()
	defer ss.Close()
	sheet := ss.AddSheet()
	sheet.Cell("A1").SetNumber(1.2)
	sheet.Cell("A2").SetNumber(2.3)
	sheet.Cell("A3").SetNumber(2.3)

	formEv := formula.NewEvaluator()

	// the formula context allows the formula evaluator to pull data from a
	// sheet
	a1Cell := sheet.FormulaContext().Cell("A1", formEv)
	fmt.Println("A1 is", a1Cell.Value())

	// So that when evaluating formulas, live workbook data is used. Formulas
	// can be evaluated directly in the context of a sheet.
	result := formEv.Eval(sheet.FormulaContext(), "SUM(A1:A3)")
	fmt.Println("SUM(A1:A3) is", result.Value())

	// Or, stored in a cell and the cell evaulated.
	sheet.Cell("A4").SetFormulaRaw("SUM(A1:A3)+SUM(A1:A3)")
	a4Value := formEv.Eval(sheet.FormulaContext(), "A4")
	fmt.Println("A4 is", a4Value.Value())

}
