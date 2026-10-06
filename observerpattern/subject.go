package main

type Subject struct{
	observer []Observer
}

func (s * Subject) RegisterObserver(o Observer){
	s.observer = append(s.observer,o)
}

func (s *Subject) UnregisterObserver(o Observer) {
    for i, observer := range s.observer {
        if observer == o {
            s.observer = append(s.observer[:i],s.observer[i+1:]...,)
            return
        }
    }
}

func (s * Subject) NotifyObservers(msg string){
	for _,obs := range s.observer{
		obs.Update(msg)
	}
}