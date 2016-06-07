package nero_test

import (
	"log"
	"runtime"
	"strconv"
	"testing"

	"github.com/bmartel/nero"
)

type NewUserAction struct {
	Email string
}

func (NewUserAction) Type() string {
	return "NewUserAction"
}

type NewSubscriptionAction struct {
	Subscription string
}

func (NewSubscriptionAction) Type() string {
	return "NewSubscriptionAction"
}

type NewUserListener struct{}

func (*NewUserListener) Listen(action nero.Action) {
	userAction := action.(NewUserAction)

	log.Println("User signed up: " + userAction.Email)
}

type NewSubscriptionListener struct{}

func (*NewSubscriptionListener) Listen(action nero.Action) {
	subscriptionAction := action.(NewSubscriptionAction)

	log.Println("New subscription was created for: " + subscriptionAction.Subscription)
}

func TestEmitAndListen(t *testing.T) {
	n := nero.New()
	defer n.Close()

	n.Listen(NewUserAction{}, &NewUserListener{})
	n.Listen(NewSubscriptionAction{}, &NewSubscriptionListener{})
	n.Emit(NewUserAction{"somerandomperson@email.com"})
	t.Logf("Number of goroutines used: %s", strconv.Itoa(runtime.NumGoroutine()))
}
