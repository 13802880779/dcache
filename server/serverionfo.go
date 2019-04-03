package server

type Node struct {
	Host string
	Port string
}

func GetServerNodes() []*Node {
	return []*Node{&Node{Host: "localhost", Port: ":8082"}}
}
