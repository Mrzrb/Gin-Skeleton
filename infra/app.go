package infra

type App struct {
	Ctls []Controller
}

func NewApp(baseCtl Controller) *App {
	return &App{
		Ctls: []Controller{baseCtl},
	}
}
