package main

import "errors"

type Observable interface {
	Subscribe(o *Observer) (bool, error)
	Unsubscribe(o *Observer) (bool, error)
	SendAll() (bool, error)
}

// Observer Interface
type Observer interface {
	HandleEvent(vacancies []string)
}

// Concrete Observer: StockObserver
type Person struct {
	name string
}

func (s *Person) HandleEvent(vacancies []string) {
	// do something
	println("Observer:", s.name, ", Vacancies has been updated", ":")
	for i := 0; i < len(vacancies); i++ {
		println(vacancies[i])
	}
}

// Concrete Subject: stockMonitor
type Jobsite struct {
	// internal state
	Vacancies   []string
	Subscribers []Observer
}

func (s *Jobsite) Subscribe(o Observer) (bool, error) {

	for _, observer := range s.Subscribers {
		if observer == o {
			return false, errors.New("Observer already exists")
		}
	}
	s.Subscribers = append(s.Subscribers, o)
	return true, nil
}

func (s *Jobsite) Unsubscribe(o Observer) (bool, error) {
	for i, observer := range s.Subscribers {
		if observer == o {
			s.Subscribers = append(s.Subscribers[:i], s.Subscribers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}
func (s *Jobsite) AddVacancy(vac string) (bool, error) {
	s.Vacancies = append(s.Vacancies, vac)
	s.SendAll()
	return true, nil
}

func (s *Jobsite) RemoveVacancy(vac string) (bool, error) {
	for i, str := range s.Vacancies {
		if str == vac {
			s.Vacancies = append(s.Vacancies[:i], s.Vacancies[i+1:]...)
			s.SendAll()
			return true, nil
		}
	}
	return false, errors.New("Vacancy not found")
}
func (s *Jobsite) SendAll() (bool, error) {
	for _, observer := range s.Subscribers {
		observer.HandleEvent(s.Vacancies)
	}
	return true, nil
}
func main() {
	hhKz := Jobsite{}
	bob := Person{"Bob"}
	hhKz.AddVacancy("Senior HTML Developer")
	hhKz.Subscribe(&bob)
	hhKz.AddVacancy("Java Junior Developer")
	hhKz.RemoveVacancy("Senior HTML Developer")
}
