package infra

type App struct {
	BaseController Controller
}

func NewApp(baseCtl Controller) *App {
	return &App{
		BaseController: baseCtl,
	}
}
