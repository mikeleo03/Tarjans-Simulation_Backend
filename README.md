# 🕸 Tarjans Algorithm Simulation
> Backend side of SCC and Bridges Finder using Tarjans Algorithm with React Typescript and Golang

## General Information
This program is created to find s Strongly Connected Component (SCC) on a directed Graph and Bridges on a undirected graph based on user input. The program will proceed a.txt file which containt directed adjacency lists to interpret a graph, then find the SCC and Bridges using popular SCC algorithm, Tarjans. User can also modify the graph by adding and deleting edges by filling the form. The program also provide a good graph and edge coloring to make a better undersanding about the results. Furthermore, the project information is also provided for future improvements.

## Project Structure
```bash
.
├─── algorithm
│   └─── algorithm.go
├─── middleware
│   └─── handlers.go
├─── router
│   └─── router.go
├─── go.mod
├─── go.sum
├─── main.go
└─── README.md
```

## Prerequisites
- gorilla/mux (v 1.8.0) to handle routing
- time to hanlde processing time
- io/ioutil to unmarshal request body
- net/http to handle request and responseWriter
- encoding/json to parse the json POST body

## Algorithms
Thus section will explain what is Tarjans Algorithm and how it used to detect Strongly Connected Components (SCC) on directed graph and bridges on undirected graph. 

## How to Compile and Run the Program
Clone this repository from terminal with this command
``` bash
$ git clone https://github.com/mikeleo03/Tarjans-Algorithm-Simulation_Backend.git
```
### Run the application on development server
Compile by running the following *command*
``` bash
$ go run main.go
```
If you do it correctly, the pogram should be running on localhost:8080.

## Contributors
<a href = "https://github.com/mikeleo03/markdown-editor/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=mikeleo03/Tarjans-Algorithm-Simulation_Backend"/>
</a>