package main

import "sockets/server"

// var wg sync.WaitGroup

// func worker(inputChan <-chan *message.UserInput, wg *sync.WaitGroup) {
// 	for input := range inputChan {
// 		println(input.IsShooting)
// 		wg.Done()
// 	}
// }

func main() {

	serv := server.New()
	serv.Start()

	// v := types.Vector2D(10, 20)
	// // u := types.Vector2D(20, 30)

	// far := types.Vector2D(10, 5)

	// println(int(v.Distance(far)))

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "websockets.html")
// })

//Netork thread

//Simulation thread;

// eventQueue := events.NewEventQ()

// simulate := simulation.NewEngine()
// eventQueue.RegisterInput(simulate)
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: true,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: true,
// // })
// simulate.Start()

// simulate.FPS = 1

// eventQueue.Start()
