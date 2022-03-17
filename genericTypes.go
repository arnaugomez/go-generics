package main

type GenericList[E any] []E

type WhatsappMessage[E any] struct {
	Id string
	content E
}