package weatherstation

type Subject interface {
	registerObserver(o Observer)
	removeObserver(o Observer)
	notifyObservers()
}
