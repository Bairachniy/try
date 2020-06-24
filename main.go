package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect("nats://TeamKiyan:AntonKiyan!6543210@10.1.133.167:4222")
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer nc.Close()

	recvCh := make(chan *FibonacciPattern)
	ec.BindRecvChan("fibonaccipatterns", recvCh)
	//var fib []FibonacciPattern

	who := <-recvCh
	//err:=json.Unmarshal(*who,fib);if err!=nil{
	//	log.Print(err)
	//}
	fmt.Println(who)
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

//
//type FibonacciPattern struct {
//	ResultUid       int     `json:"resultUid"`
//	Pattern         string  `json:"pattern"`
//	Symbol          string  `json:"symbol"`
//	SymbolCode      string  `json:"symbolCode"`
//	Exchange        string  `json:"exchange"`
//	Completed       int     `json:"completed"`
//	Quality         int     `json:"quality"`
//	Interval        int     `json:"interval"`
//	Uniformity      int     `json:"uniformity"`
//	Initialtrend    int     `json:"initialtrend"`
//	Clarity         int     `json:"clarity"`
//	PriceSymmetry   int     `json:"priceSymmetry"`
//	TimeSymmetry    int     `json:"timeSymmetry"`
//	PatternEndTime  int     `json:"patternEndTime"`
//	DirectionX      string  `json:"directionX"`
//	DirectionA      string  `json:"directionA"`
//	DirectionB      string  `json:"directionB"`
//	DirectionC      string  `json:"directionC"`
//	DirectionD      string  `json:"directionD"`
//	PriceA          int     `json:"priceA"`
//	PriceB          int     `json:"priceB"`
//	PriceC          int     `json:"priceC"`
//	PriceD          int     `json:"priceD"`
//	PriceX          int     `json:"priceX"`
//	TimeA           int     `json:"timeA"`
//	TimeB           int     `json:"timeB"`
//	TimeC           int     `json:"timeC"`
//	TimeD           int     `json:"timeD"`
//	TimeX           int     `json:"timeX"`
//	Target0         float32 `json:"target0"`
//	Target3         float32 `json:"target3"`
//	Target5         float32 `json:"target5"`
//	Target6         float32 `json:"target6"`
//	Target7         float32 `json:"target7"`
//	Target10        float32 `json:"target10"`
//	Target12        float32 `json:"target12"`
//	Target16        float32 `json:"target16"`
//	Length          int     `json:"length"`
//	Relevant        int     `json:"relevant"`
//	Age             int     `json:"age"`
//	PatternImageUrl string  `json:"patternImageUrl"`
//	ClickThroughUrl string  `json:"clickThroughUrl"`
//	TimezoneOffset  int     `json:"timezoneOffset"`
//	DemoCandleDelay int     `json:"demoCandleDelay"`
//	DemoMinuteDelay int     `json:"demoMinuteDelay"`
//	SessionId       string  `json:"sessionId"`
//}
type AuthResponse struct {
	Code     int     `json:"code"`
	Response string  `json:"response"`
	Payload  Payload `json:"payload"`
}

type Payload struct {
	CrmToken      string `json:"crmToken"`
	SessionId     string `json:"sessionId"`
	StompUser     string `json:"stompUser"`
	StompPassword string `json:"stompPassword"`
}

//
//type Cook struct{
//	UmSession string `json:"um_session"`
//	Path string `json:"Path"`
//	Expires string `json:"Expires"`
//}
