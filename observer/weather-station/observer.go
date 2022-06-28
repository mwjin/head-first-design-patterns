package weatherstation

type Observer interface {
	update(temp, humidity, pressure float32)
}
