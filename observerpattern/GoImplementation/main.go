//              Subject
//                 │
//                 │ []Observer
//                 │
//       ┌─────────┼─────────┐
//       ↓         ↓         ↓
//    UserA      UserB     UserC
//       │         │         │
//       ↓         ↓         ↓
//    Update()   Update()   Update()

package main

func main() {
	subject := &Subject{}

	// UserA := &UserA{}
	// UserB := &UserB{}

	// subject.RegisterObserver(UserA)
	// subject.RegisterObserver(UserB)
	user1 := &User{UserID: 101}
	user2 := &User{UserID: 102}
	user3 := &User{UserID: 103}

	subject.RegisterObserver(user1)
	subject.RegisterObserver(user2)
	subject.RegisterObserver(user3)

	subject.NotifyObservers("New message")

	subject.UnregisterObserver(user1)

	subject.NotifyObservers("New New message")
}
