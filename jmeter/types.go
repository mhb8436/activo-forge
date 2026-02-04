package jmeter

import "encoding/xml"

// JMeterTestPlan 루트 요소
type JMeterTestPlan struct {
	XMLName    xml.Name `xml:"jmeterTestPlan"`
	Version    string   `xml:"version,attr"`
	Properties string   `xml:"properties,attr"`
	Jmeter     string   `xml:"jmeter,attr"`
	HashTree   HashTree `xml:"hashTree"`
}

// HashTree JMeter 해시트리
type HashTree struct {
	XMLName xml.Name      `xml:"hashTree"`
	Items   []interface{} `xml:",any"`
}

// TestPlan 테스트 플랜
type TestPlan struct {
	XMLName     xml.Name     `xml:"TestPlan"`
	GuiClass    string       `xml:"guiclass,attr"`
	TestClass   string       `xml:"testclass,attr"`
	TestName    string       `xml:"testname,attr"`
	Enabled     string       `xml:"enabled,attr"`
	BoolProps   []BoolProp   `xml:"boolProp"`
	StringProps []StringProp `xml:"stringProp"`
	ElementProp ElementProp  `xml:"elementProp"`
}

// ThreadGroup 스레드 그룹
type ThreadGroup struct {
	XMLName     xml.Name           `xml:"ThreadGroup"`
	GuiClass    string             `xml:"guiclass,attr"`
	TestClass   string             `xml:"testclass,attr"`
	TestName    string             `xml:"testname,attr"`
	Enabled     string             `xml:"enabled,attr"`
	StringProps []StringProp       `xml:"stringProp"`
	BoolProps   []BoolProp         `xml:"boolProp"`
	ElementProp LoopControllerProp `xml:"elementProp"`
}

// HTTPSamplerProxy HTTP 샘플러
type HTTPSamplerProxy struct {
	XMLName     xml.Name         `xml:"HTTPSamplerProxy"`
	GuiClass    string           `xml:"guiclass,attr"`
	TestClass   string           `xml:"testclass,attr"`
	TestName    string           `xml:"testname,attr"`
	Enabled     string           `xml:"enabled,attr"`
	BoolProps   []BoolProp       `xml:"boolProp"`
	StringProps []StringProp     `xml:"stringProp"`
	ElementProp *ElementPropArgs `xml:"elementProp,omitempty"`
}

// StringProp 문자열 속성
type StringProp struct {
	XMLName xml.Name `xml:"stringProp"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

// BoolProp 불린 속성
type BoolProp struct {
	XMLName xml.Name `xml:"boolProp"`
	Name    string   `xml:"name,attr"`
	Value   bool     `xml:",chardata"`
}

// IntProp 정수 속성
type IntProp struct {
	XMLName xml.Name `xml:"intProp"`
	Name    string   `xml:"name,attr"`
	Value   int      `xml:",chardata"`
}

// ElementProp 요소 속성
type ElementProp struct {
	XMLName        xml.Name       `xml:"elementProp"`
	Name           string         `xml:"name,attr"`
	ElementType    string         `xml:"elementType,attr"`
	GuiClass       string         `xml:"guiclass,attr"`
	TestClass      string         `xml:"testclass,attr"`
	CollectionProp CollectionProp `xml:"collectionProp"`
}

// ElementPropArgs Arguments용 요소 속성
type ElementPropArgs struct {
	XMLName        xml.Name           `xml:"elementProp"`
	Name           string             `xml:"name,attr"`
	ElementType    string             `xml:"elementType,attr"`
	GuiClass       string             `xml:"guiclass,attr"`
	TestClass      string             `xml:"testclass,attr"`
	CollectionProp CollectionPropArgs `xml:"collectionProp"`
}

// LoopControllerProp 루프 컨트롤러 속성
type LoopControllerProp struct {
	XMLName     xml.Name     `xml:"elementProp"`
	Name        string       `xml:"name,attr"`
	ElementType string       `xml:"elementType,attr"`
	GuiClass    string       `xml:"guiclass,attr"`
	TestClass   string       `xml:"testclass,attr"`
	BoolProps   []BoolProp   `xml:"boolProp"`
	StringProps []StringProp `xml:"stringProp"`
}

// CollectionProp 컬렉션 속성
type CollectionProp struct {
	XMLName xml.Name `xml:"collectionProp"`
	Name    string   `xml:"name,attr"`
}

// CollectionPropArgs Arguments용 컬렉션 속성
type CollectionPropArgs struct {
	XMLName xml.Name       `xml:"collectionProp"`
	Name    string         `xml:"name,attr"`
	Items   []HTTPArgument `xml:"elementProp"`
}

// HTTPArgument HTTP 인자
type HTTPArgument struct {
	XMLName     xml.Name     `xml:"elementProp"`
	Name        string       `xml:"name,attr"`
	ElementType string       `xml:"elementType,attr,omitempty"`
	StringProps []StringProp `xml:"stringProp"`
	BoolProps   []BoolProp   `xml:"boolProp"`
}
