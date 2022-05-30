package z_either_test

import (
	"fmt"
	"github.com/MickStanciu/SC-Generics/generics/z_either"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go/types"
	"strings"
	"testing"
)

func TestEither_MapWhenNoErrors(t *testing.T) {
	e := z_either.NewRight[types.Nil, string]("Hello")

	appendX := func(val string) z_either.Either[types.Nil, string] {
		return z_either.NewRight[types.Nil, string](val + "X")
	}

	appendY := func(val string) z_either.Either[types.Nil, string] {
		return z_either.NewRight[types.Nil, string](val + "Y")
	}

	appendZ := func(val string) z_either.Either[types.Nil, string] {
		return z_either.NewRight[types.Nil, string](val + "Z")
	}

	maybeResult := e.
		Map(appendX).
		Map(appendY).
		Map(appendZ)

	require.False(t, maybeResult.IsLeft())
	result := maybeResult.GetRightOrPanic()

	assert.EqualValues(t, "HelloXYZ", result)
}

func TestEither_MapWhenErrors(t *testing.T) {
	e := z_either.NewRight[error, string]("Hello")

	appendX := func(val string) z_either.Either[error, string] {
		return z_either.NewRight[error, string](val + "X")
	}

	appendY := func(val string) z_either.Either[error, string] {
		return z_either.NewLeft[error, string](fmt.Errorf("something blew up"))
	}

	appendZ := func(val string) z_either.Either[error, string] {
		return z_either.NewRight[error, string](val + "Z")
	}

	maybeResult := e.
		Map(appendX).
		Map(appendY).
		Map(appendZ)

	require.True(t, maybeResult.IsLeft())
	result := maybeResult.GetLeftOrPanic()
	assert.EqualValues(t, "something blew up", result.Error())
}

// FAKES
type user struct {
	name             string
	email            string
	subscribed       bool
	welcomeEmailSent bool
}

func fakeCreateUser(name string, email string) z_either.Either[error, user] {
	fmt.Println("> fakeCreateUser called")

	//Other services called here ...
	return z_either.NewRight[error, user](user{
		name:  name,
		email: email,
	})
}

func fakeSanitizeEmail(e z_either.Either[error, user]) z_either.Either[error, user] {
	fmt.Println("> fakeSanitizeEmail called")
	if e.IsLeft() {
		return e
	}

	value := e.GetRightOrPanic()
	value.email = strings.ToLower(value.email)
	return z_either.NewRight[error, user](value)
}

func fakeSendWelcomeEmail(e z_either.Either[error, user]) z_either.Either[error, user] {
	fmt.Println("> fakeSendWelcomeEmail called")
	if e.IsLeft() {
		return e
	}

	value := e.GetRightOrPanic()
	value.welcomeEmailSent = true
	return z_either.NewRight[error, user](value)
}

func fakeSendWelcomeEmailFail(e z_either.Either[error, user]) z_either.Either[error, user] {
	fmt.Println("> fakeSendWelcomeEmailFail called")
	return z_either.NewLeft[error, user](fmt.Errorf("something in DB is screwed up"))
}

func TestScenario_BindWhenFailing(t *testing.T) {
	maybeResult := fakeCreateUser("George", "George@email.com").
		Bind(fakeSanitizeEmail).
		Bind(fakeSanitizeEmail).
		Bind(fakeSendWelcomeEmail).
		Bind(fakeSendWelcomeEmailFail).
		Bind(fakeSendWelcomeEmail)

	require.NotNil(t, maybeResult)
	assert.True(t, maybeResult.IsLeft())
	assert.EqualValues(t, "something in DB is screwed up", maybeResult.GetLeftOrPanic().Error())
}

func TestScenario_BindWhenNotFailing(t *testing.T) {
	maybeResult := fakeCreateUser("George", "George@email.com").
		Bind(fakeSanitizeEmail).
		Bind(fakeSanitizeEmail).
		Bind(fakeSendWelcomeEmail).
		Bind(fakeSendWelcomeEmail)

	require.NotNil(t, maybeResult)
	assert.False(t, maybeResult.IsLeft())

	result := maybeResult.GetRightOrPanic()
	assert.EqualValues(t, "george@email.com", result.email)
	assert.EqualValues(t, true, result.welcomeEmailSent)
}
