package meta

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

// sessionListReq is to serialize structured data to the expected msgpack format
// _msgpack => forces it to be returned as an array of elements
type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

// SessionListRes is the session return from the rpc server.
// ,omitempty => data is optional so that encoding and decoding aren't expected, which will flatten the data so that it is not a nested map
type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort string `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploidUUID string `msgpack:"exploit_uuid"`
}

type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

// B/c GO dynamically serializes the response it is possible to have both a successful login and unsuccessful in one structure
// i.e. GO only fills in what is needed
type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        string `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

// Metasploit is to log into te server.
type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

// New returns a new metasploit session info
func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}
	if err := msf.Login(); err != nil {
		return nil, err
	}
	return msf, nil
}

// To be used in every RPC function built
// Interface => generic enough to allow for not needing to define every req and res struct
func (msf *Metasploit) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	dest := fmt.Sprintf("http://%s/api", msf.host)
	r, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}

// To log into the running metasploit server
func (msf *Metasploit) Login() error {

	ctx := &loginReq{
		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}

	var res loginRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	msf.token = res.Token
	return nil
}

func (msf *Metasploit) Logout() error {
	ctx := &logoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}

	var res logoutRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	// set token to empty
	msf.token = ""
	return nil
}

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {
	req := &sessionListReq{Method: "session.list", Token: msf.token}
	res := make(map[uint32]SessionListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	// flattening the map of maps
	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil
}
