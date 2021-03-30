package main

type DataStore interface {
	GetName() string
	GetProfile() Profile
}

type Profile struct {
	Username string
	Dob      string
	Age      string
	Email    string
	Phone    string
}

type Details struct {
	name    string
	profile Profile
}

func (d *Details) GetName() string {
	return d.name
}

func (d *Details) GetProfile() Profile {
	return d.profile
}
