package db

import "fmt"

type EventReceiver struct {
}

func (e EventReceiver) Event(eventName string) {
	fmt.Println("Event: " + eventName)
}
func (e EventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Println("Event: " + eventName)
	fmt.Println(kvs)
}
func (e EventReceiver) EventErr(eventName string, err error) error {
	fmt.Println("Event: " + eventName)
	fmt.Println(err)
	return nil
}
func (e EventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Println("Event: " + eventName)
	fmt.Println(err)
	fmt.Println(kvs)
	return nil
}
func (e EventReceiver) Timing(eventName string, nanoseconds int64) {
	fmt.Println("Event: " + eventName)
	fmt.Println(nanoseconds)
}
func (e EventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {
	fmt.Println("Event: " + eventName)
	fmt.Println(nanoseconds)
	fmt.Println(kvs)
}
