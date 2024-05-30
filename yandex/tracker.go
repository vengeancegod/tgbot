// package yandex

// import (
// 	"fmt"

// 	tracker "github.com/dvsnin/yandex-tracker-go"
// )

// func YandexTracker() {
// 	client := tracker.New("YOUR YANDEX.TRACKER TOKEN", "YOUR YANDEX ORG_ID", "TICKET KEY")
// 	ticket, err := client.GetTicket()
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// 	fmt.Printf("%s\n", ticket.Description())
// }