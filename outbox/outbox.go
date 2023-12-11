package outbox

import "fmt"

var Outbox = make(chan OutboundRequest, 100)

func ProcessOubox() {
	go func() {
		for {
			select {
			case outbound := <-Outbox:
				makeOutboundRequest(outbound)
			}
		}
	}()
}

func makeOutboundRequest(outbound OutboundRequest) {
	// send request back to slack

	fmt.Printf(`outbound request: [caller:%s] [response:%s]`, outbound.Caller, outbound.Response)

}
