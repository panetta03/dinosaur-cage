package models

// DinosaurFactory is an interface for creating dinosaurs.
type DinosaurFactory interface {
	CreateDinosaur(name string) Dinosaur
}
