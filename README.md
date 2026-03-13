# NumberGoApp 

Simple TUI app made using Go and bubbletea 

Idle game with theme of scaling an application for infinite users 


## Structure

```bash
├── cmd
│   └── numbergoapp
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── engine
│   │   └── tick.go
│   ├── game
│   │   ├── actions.go
│   │   ├── levels.go
│   │   ├── state.go
│   │   ├── tick.go
│   │   └── upgrades.go
│   ├── save
│   │   ├── convert.go
│   │   ├── save.go
│   │   └── types.go
│   └── ui
│       ├── model.go
│       ├── styles.go
│       └── view.go
├── README.md


```

There are 4 main modules
1. engine - mainly for the tick
2. game - game logic
3. save - save file and loading
4. ui - loading TUI and interaction


## To run
Make sure you have go installed

At root directory of project run
```bash
go run cmd/numbergoapp/main.go
```

