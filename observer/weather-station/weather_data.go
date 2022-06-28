package weatherstation

type WeatherData struct {
	observers   []Observer
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) registerObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) removeObserver(o Observer) {
	w.observers[w.findObserver(o)] = w.observers[len(w.observers)-1]
	w.observers = w.observers[:len(w.observers)-1]
}

func (w *WeatherData) findObserver(o Observer) int {
	for i, observer := range w.observers {
		if observer == o {
			return i
		}
	}
	return -1
}

func (w *WeatherData) notifyObservers() {
	for _, observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) measurementsChanged() {
	w.notifyObservers()
}

func (w *WeatherData) setMeasurements(temperature, humidity, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}
