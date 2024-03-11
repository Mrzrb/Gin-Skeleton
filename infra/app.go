package infra

type App struct {
	Ctls []Controller `wire:"-"`
}
