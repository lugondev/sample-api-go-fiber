package notifier

type Notifier interface {
	Notify(message Message) error
}

type Message struct {
	Content string
}
