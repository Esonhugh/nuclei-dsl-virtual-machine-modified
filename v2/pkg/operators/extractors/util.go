package extractors

import (
	"github.com/Knetic/govaluate"
	"github.com/projectdiscovery/dsl"
)

// SupportsMap determines if the extractor type requires a map
func SupportsMap(extractor *Extractor) bool {
	return extractor.Type.ExtractorType == KValExtractor || extractor.Type.ExtractorType == DSLExtractor
}

func (e *Extractor) ExtractorSpecificDSLFunctions() map[string]govaluate.ExpressionFunction {
	funcs := dsl.HelperFunctions()
	loadFunc := govaluate.ExpressionFunction(
		func(args ...interface{}) (interface{}, error) {
			return e.loadDSLData(args[0].(string)), nil
		})
	storeFunc := govaluate.ExpressionFunction(
		func(args ...interface{}) (interface{}, error) {
			e.storeDSLData(args[0].(string), args[1])
			return nil, nil
		})
	helperFunctions := make(map[string]govaluate.ExpressionFunction)
	for k, v := range funcs {
		helperFunctions[k] = v
	}
	helperFunctions["load"] = loadFunc
	helperFunctions["store"] = storeFunc
	return helperFunctions
}

func (e *Extractor) loadDSLData(key string) any {
	if e.dslData == nil {
		return nil
	}
	return e.dslData[key]
}

func (e *Extractor) storeDSLData(key string, data interface{}) {
	if e.dslData == nil {
		e.dslData = make(map[string]interface{})
	}
	e.dslData[key] = data
}

func (e *Extractor) mergeOtherData(data map[string]interface{}) map[string]interface{} {
	Merege := make(map[string]interface{})
	for k, v := range e.dslData {
		Merege[k] = v
	}
	for k, v := range data {
		Merege[k] = v
	}
	// e.dslData = Merege
	return Merege
}
