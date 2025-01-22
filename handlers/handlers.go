package handlers

import (
	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func Init(sm *scs.SessionManager) {
	sessionManager = sm
}
