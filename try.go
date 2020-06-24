package main

import (
	"awesomeProject/dataserver/logger"
	"encoding/json"
	"fmt"

	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var Patterns TypePattern

func main() {
	fives := [5]string{"fibonaccipatterns", "chartpatterns", "bigmovementpatterns", "keylevelspattern", "consecutivecandlespatterns"}
	for _, p := range fives {
		five(p)
	}
	fmt.Println("FIBONACHI", Patterns.FibonacciPattern[0])
	fmt.Println("BIGMOVEMENT", Patterns.BigMovementPattern[0])
	fmt.Println("CHART", Patterns.ChartPattern[0])
	fmt.Println("COSECUUTIVE", Patterns.ConsecutiveCandelsPattern[0])
	fmt.Println("KEYLEVEL", Patterns.KeyLevelPattern[0])
}

func five(pattern string) {
	nc, _ := nats.Connect("nats://TeamKiyan:AntonKiyan!6543210@10.1.133.167:4222")
	// Flush connection to server, returns when all messages have been processed.
	sub, err := nc.SubscribeSync(pattern)
	if err != nil {
		log.Print(err)
	}

	defer func() {
		err = sub.Unsubscribe()
		if err != nil {
			logger.Warningf("Unsubscribe returned error: %s", err)
		}
	}()
	m, err := sub.NextMsg(time.Second * 15)
	if err != nil {
		log.Print(err)
	}

	switch pattern {
	case "fibonaccipatterns":
		var masF []FibonacciPattern
		err = json.Unmarshal(m.Data, &masF)
		for _, m := range masF {
			Patterns.FibonacciPattern = append(Patterns.FibonacciPattern, m)
		}

	case "chartpatterns":
		var masF []ChartPattern
		err = json.Unmarshal(m.Data, &masF)
		for _, m := range masF {
			Patterns.ChartPattern = append(Patterns.ChartPattern, m)
		}

	case "bigmovementpatterns":
		var masF []BigMovementPattern
		err = json.Unmarshal(m.Data, &masF)
		for _, m := range masF {
			Patterns.BigMovementPattern = append(Patterns.BigMovementPattern, m)
		}

	case "keylevelspattern":
		var masF []KeyLevelPattern
		err = json.Unmarshal(m.Data, &masF)
		for _, m := range masF {
			Patterns.KeyLevelPattern = append(Patterns.KeyLevelPattern, m)
		}

	case "consecutivecandlespatterns":
		var masF []ConsecutiveCandelsPattern
		err = json.Unmarshal(m.Data, &masF)
		for _, m := range masF {
			Patterns.ConsecutiveCandelsPattern = append(Patterns.ConsecutiveCandelsPattern, m)
		}

	}

}

type TypePattern struct {
	ChartPattern              []ChartPattern
	FibonacciPattern          []FibonacciPattern
	KeyLevelPattern           []KeyLevelPattern
	BigMovementPattern        []BigMovementPattern
	ConsecutiveCandelsPattern []ConsecutiveCandelsPattern
}

//func SelectPattern(str string)  []type{
//	var tmp []type
//	switch str {
//	case "fibonaccipatterns":tmp=[]FibonacciPattern{}
//		return tmp
//	case "chartpatterns":
//		return &ChartPattern{}
//	case "bigmovementpatterns":
//		return &BigMovementPattern{}
//	case "keylevelspattern":
//		return &KeyLevelPattern{}
//	case "consecutivecandlespatterns":
//		return &ConsecutiveCandelsPattern{}
//	}
//	return ""
//}
type ConsecutiveCandelsPattern struct {
	ResultUid        int     `json:"resultUid"`
	Symbol           string  `json:"symbol"`
	SymbolCode       string  `json:"symbolCode"`
	PatternName      string  `json:"patternName"`
	Exchange         string  `json:"exchange"`
	Direction        int     `json:"direction"`
	Interval         int     `json:"interval"`
	PatternEndTime   int     `json:"patternEndTime"`
	PatternStartTime int     `json:"patternStartTime"`
	Length           int     `json:"length"`
	Relevant         int     `json:"relevant"`
	Age              int     `json:"age"`
	PatternImageUrl  string  `json:"patternImageUrl"`
	ClickThroughUrl  string  `json:"clickThroughUrl"`
	TimezoneOffset   float32 `json:"timezoneOffset"`
	DemoCandleDelay  int     `json:"demoCandleDelay"`
	DemoMinuteDelay  int     `json:"demoMinuteDelay"`
}
type BigMovementPattern struct {
	ResultUid        int     `json:"resultUid"`
	Symbol           string  `json:"symbol"`
	SymbolCode       string  `json:"symbolCode"`
	PatternName      string  `json:"patternName"`
	Exchange         string  `json:"exchange"`
	Direction        int     `json:"direction"`
	Interval         int     `json:"interval"`
	PatternEndTime   int     `json:"patternEndTime"`
	PatternStartTime int     `json:"patternStartTime"`
	Length           int     `json:"length"`
	Relevant         int     `json:"relevant"`
	Age              int     `json:"age"`
	PatternImageUrl  string  `json:"patternImageUrl"`
	ClickThroughUrl  string  `json:"clickThroughUrl"`
	TimezoneOffset   float32 `json:"timezoneOffset"`
	DemoCandleDelay  int     `json:"demoCandleDelay"`
	DemoMinuteDelay  int     `json:"demoMinuteDelay"`
}
type KeyLevelPattern struct {
	ResultUid           int     `json:"resultUid"`
	Pattern             string  `json:"pattern"`
	Symbol              string  `json:"symbol"`
	SymbolCode          string  `json:"symbolCode"`
	Exchange            string  `json:"exchange"`
	Completed           int     `json:"completed"`
	Quality             int     `json:"quality"`
	Interval            int     `json:"interval"`
	Direction           int     `json:"direction"`
	PatternEndTime      int     `json:"patternEndTime"`
	Price               float32 `json:"price"`
	Point1              int     `json:"point1"`
	Point2              int     `json:"point2"`
	Point3              int     `json:"point3"`
	Point4              int     `json:"point4"`
	Point5              int     `json:"point5"`
	Point6              int     `json:"point6"`
	Point7              int     `json:"point7"`
	Point8              int     `json:"point8"`
	Point9              int     `json:"point9"`
	Point10             int     `json:"point10"`
	Length              int     `json:"length"`
	StopLoss            float32 `json:"stopLoss"`
	PredictionPriceFrom float32 `json:"predictionPriceFrom"`
	PredictionPriceTo   float32 `json:"predictionPriceTo"`
	PredictionTimeTo    int     `json:"predictionTimeTo"`
	PredictionTimeFrom  int     `json:"predictionTimeFrom"`
	Relevant            int     `json:"relevant"`
	Age                 int     `json:"age"`
	PatternImageUrl     string  `json:"patternImageUrl"`
	ClickThroughUrl     string  `json:"clickThroughUrl"`
	TimezoneOffset      float32 `json:"timezoneOffset"`
	DemoCandleDelay     int     `json:"demoCandleDelay"`
	DemoMinuteDelay     int     `json:"demoMinuteDelay"`
	SessionId           string  `json:"sessionId"`
}
type ChartPattern struct {
	ResultUid           int     `json:"resultUid"`
	Pattern             string  `json:"pattern"`
	Symbol              string  `json:"symbol"`
	SymbolCode          string  `json:"symbolCode"`
	Exchange            string  `json:"exchange"`
	Direction           int     `json:"direction"`
	Initialtrend        int     `json:"initialtrend"`
	Quality             int     `json:"quality"`
	Uniformity          int     `json:"uniformity"`
	Breakout            int     `json:"breakout"`
	Clarity             int     `json:"clarity"`
	Interval            int     `json:"interval"`
	Completed           int     `json:"completed"`
	PatternEndTime      int     `json:"patternendTime"`
	ResistanceX0        int     `json:"resistanceX0"`
	ResistanceX1        int     `json:"resistanceX1"`
	ResistanceY0        float32 `json:"resistanceY0"`
	ResistanceY1        float32 `json:"resistanceY1"`
	SupportX0           int     `json:"supportX0"`
	SupportX1           int     `json:"supportX1"`
	SupportY0           float32 `json:"supportY0"`
	SupportY1           float32 `json:"supportY1"`
	VolumeIncrease      int     `json:"volumeIncrease"`
	StopLoss            float32 `json:"stopLoss"`
	PredictionPriceFrom float32 `json:"predictionPriceFrom"`
	PredictionPriceTo   float32 `json:"predictionPriceTo"`
	PredictionTimeTo    int     `json:"predictionTimeTo"`
	PredictionTimeFrom  int     `json:"predictionTimeFrom"`
	Trend               string  `json:"trend"`
	Length              int     `json:"length"`
	Relevant            int     `json:"relevant"`
	Age                 int     `json:"age"`
	PatternImageUrl     string  `json:"patternImageUrl"`
	ClickThroughUrl     string  `json:"clickThroughUrl"`
	TimezoneOffset      float32 `json:"timezoneOffset"`
	DemoCandleDelay     int     `json:"demoCandleDelay"`
	DemoMinuteDelay     int     `json:"demoMinuteDelay"`
	SessionId           string  `json:"sessionId"`
}
type FibonacciPattern struct {
	ResultUid       int     `json:"resultUid"`
	Pattern         string  `json:"pattern"`
	Symbol          string  `json:"symbol"`
	SymbolCode      string  `json:"symbolCode"`
	Exchange        string  `json:"exchange"`
	Completed       int     `json:"completed"`
	Quality         int     `json:"quality"`
	Interval        int     `json:"interval"`
	Uniformity      int     `json:"uniformity"`
	Initialtrend    int     `json:"initialtrend"`
	Clarity         int     `json:"clarity"`
	PriceSymmetry   int     `json:"priceSymmetry"`
	TimeSymmetry    int     `json:"timeSymmetry"`
	PatternEndTime  int     `json:"patternEndTime"`
	DirectionX      string  `json:"directionX"`
	DirectionA      string  `json:"directionA"`
	DirectionB      string  `json:"directionB"`
	DirectionC      string  `json:"directionC"`
	DirectionD      string  `json:"directionD"`
	PriceA          float32 `json:"priceA"`
	PriceB          float32 `json:"priceB"`
	PriceC          float32 `json:"priceC"`
	PriceD          float32 `json:"priceD"`
	PriceX          float32 `json:"priceX"`
	TimeA           int     `json:"timeA"`
	TimeB           int     `json:"timeB"`
	TimeC           int     `json:"timeC"`
	TimeD           int     `json:"timeD"`
	TimeX           int     `json:"timeX"`
	Target0         float32 `json:"target0"`
	Target3         float32 `json:"target3"`
	Target5         float32 `json:"target5"`
	Target6         float32 `json:"target6"`
	Target7         float32 `json:"target7"`
	Target10        float32 `json:"target10"`
	Target12        float32 `json:"target12"`
	Target16        float32 `json:"target16"`
	Length          int     `json:"length"`
	Relevant        int     `json:"relevant"`
	Age             int     `json:"age"`
	PatternImageUrl string  `json:"patternImageUrl"`
	ClickThroughUrl string  `json:"clickThroughUrl"`
	TimezoneOffset  float32 `json:"timezoneOffset"`
	DemoCandleDelay int     `json:"demoCandleDelay"`
	DemoMinuteDelay int     `json:"demoMinuteDelay"`
	SessionId       string  `json:"sessionId"`
}
