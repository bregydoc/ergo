package ergo

// Translater is a struct to describe a translater worker, this search all the db and translate all user messages
type Translater struct {
	repo   Repository
	stream chan int
}
