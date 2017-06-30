package useCases

import (
	"log"

	"github.com/err0r500/cleanArchitectureGolang/src/domain"
)

// User : a more complete user with all app-specific details
// it should contains all the field that may be sent to a presentationnal layer (http response for example)
// a simple Format() can then be implemented to fileter out or alter the properties name/values
type User struct {
	domain.User
	ID   int
	name string
}

// NewFlowUser : constructor to be sure everything is well initialized
func NewFlowUser(dU domain.User, id int, name string) User {
	return User{dU, id, name}
}

//UserInteractor : all the interfaces needed in order to execute all the useCases
type UserInteractor struct {
	WeakCheck      domain.UserAddressChecker
	StrongCheck    domain.UserAddressChecker
	UserReadWriter UserReadWriter
}

//NewUserInteractor : constructor to be sure everything is well initialized
func NewUserInteractor(wC, sC domain.UserAddressChecker, uRW UserReadWriter) UserInteractor {
	return UserInteractor{wC, sC, uRW}
}

//UserReadWriter interface that will typîcally be defined at "interfaces" level
// NB : a dummy version for test is in useCases/dummy
type UserReadWriter interface {
	GetDetails(interface{}) User
}

//
// the useCases !
//

//UseCase1 : a usecase representing a complete "scenario"
func (t UserInteractor) UseCase1(params interface{}) {
	log.Print("start of UC 1")
	domainUser := domain.NewUser("hehehe", t.WeakCheck)
	fU := NewFlowUser(domainUser, 112345678, "userName")
	fU.User.CheckAddress(fU.User, nil)
	log.Print("end of UC 1")
}

//UseCase2 : another use case
func (t UserInteractor) UseCase2(params interface{}) {
	log.Print("start of UC 2")
	//GetDetails receives params potentially unkown at "useCases" level
	// and return a usecases.User according to that
	fU := t.UserReadWriter.GetDetails(params)
	fU.User.UserAddressChecker = t.StrongCheck
	fU.User.CheckAddress(fU.User, nil)

	log.Print("end of UC 2")
}
