package main

type NotificationAPIWrapper interface {
	Send(to string, message string) error
}
