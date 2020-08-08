package media

import "strings"

// Catalogable interface
type Catalogable interface {
	NewMovie(title string, rating Rating, boxOffice float32)
	GetTitle() string
	GetRating() string
	GetBoxOffice() float32
	SetTitle(newTitle string)
	SetRating(newRating Rating)
	SetBoxOffice(newBoxOffice float32)
}

// Movie class
type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

// Rating type
type Rating string

// Rating Constant Values
const (
	R    = "R (Restricted)"
	G    = "G (General Audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

// NewMovie constructs a new Movie object
func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

// GetTitle gets the title
func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

// GetRating gets the rating
func (m *Movie) GetRating() string {
	return string(m.rating)
}

// GetBoxOffice gets the box office
func (m *Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

// SetTitle sets the title
func (m *Movie) SetTitle(newTitle string) {
	m.title = newTitle
}

// SetRating sets the rating
func (m *Movie) SetRating(newRating Rating) {
	m.rating = newRating
}

// SetBoxOffice sets the box office
func (m *Movie) SetBoxOffice(newBoxOffice float32) {
	m.boxOffice = newBoxOffice
}
