package session

type SessionManager interface{
	Init(addr string, options ...string) (err error)
	CreateSession()(session Session, err error)
	Get(sessionId string)(session Session, err error)
}