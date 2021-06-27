package main

var server = createServer()

type ServeraddClient func(server2 *Server) int

type Server struct {
	clients    map[int]*Client
	add_client ServeraddClient
}

func createServer() Server {
	return Server{
		clients: make(map[int]*Client),
		add_client: func(server2 *Server) int {
			id := len(server2.clients) + 1
			for {
				if _, ok := server2.clients[id]; ok {
					id += 1
				} else {
					break
				}
			}
			newClient := Client{
				Id:    id,
				World: nil, // Is going to be set when the socket get the word 'start'
				Config: Configurations{
					RunType: 3,
					Pram:    false,
					Ubers:   make([]UberFormat, 1),
					Request: make([]RequestFormat, 1),
				},
			}
			server2.clients[id] = &newClient
			return id
		},
	}
}
