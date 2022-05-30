package y_maybe_test

import (
	"fmt"
	"github.com/MickStanciu/SC-Generics/generics/y_maybe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSomething_Exists(t *testing.T) {
	maybeName := y_maybe.NewSomething("George")
	assert.True(t, maybeName.Exists())
}

func TestSomething_Get(t *testing.T) {
	name, err := y_maybe.
		NewSomething("George").
		Get()

	require.Nil(t, err)
	assert.EqualValues(t, "George", name)
}

func TestSomething_GetOrElse(t *testing.T) {
	name := y_maybe.
		NewSomething("George").
		GetOrElse("")

	assert.EqualValues(t, "George", name)
}

func TestNothing_Exists(t *testing.T) {
	maybeNothing := y_maybe.NewNothing("George")
	assert.False(t, maybeNothing.Exists())
}

func TestNothing_Get(t *testing.T) {
	_, err := y_maybe.
		NewNothing("George").
		Get()

	require.NotNil(t, err)
	assert.EqualValues(t, "expected something but got nothing", err.Error())
}

func TestNothing_GetOrElse(t *testing.T) {
	name := y_maybe.
		NewNothing("George").
		GetOrElse("Horace")

	require.NotNil(t, name)
	assert.EqualValues(t, "Horace", name)
}

// FAKES
type user struct {
	name       string
	email      string
	subscribed bool
}

func fakeCreateUserTruthy(name string, email string) y_maybe.Maybe[user] {
	fmt.Println("> fakeCreateUserTruthy called")
	//Other services called here ...
	return y_maybe.NewSomething(user{
		name:  name,
		email: email,
	})
}

func fakeSignUpForMarketingEmailsFalsy(y_maybe.Maybe[user]) y_maybe.Maybe[user] {
	fmt.Println("> fakeSignUpForMarketingEmailsFalsy called")
	// Fakes error
	return y_maybe.NewNothing(user{})
}

func fakeSendWelcomeEmailTruthy(maybeUser y_maybe.Maybe[user]) y_maybe.Maybe[user] {
	fmt.Println("> fakeSendWelcomeEmailTruthy called")
	panic("not implemented")
}

func TestScenario(t *testing.T) {
	maybeResult := fakeCreateUserTruthy("George", "george@email.com").
		Bind(fakeSignUpForMarketingEmailsFalsy).
		Bind(fakeSendWelcomeEmailTruthy).
		Bind(fakeSendWelcomeEmailTruthy).
		Bind(fakeSendWelcomeEmailTruthy).
		Bind(fakeSendWelcomeEmailTruthy).
		Bind(fakeSendWelcomeEmailTruthy)

	require.NotNil(t, maybeResult)
	assert.False(t, maybeResult.Exists())

	_, err := maybeResult.Get()
	assert.EqualValues(t, "expected something but got nothing", err.Error())
}
