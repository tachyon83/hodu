package hodu

import "encoding/json"
import "fmt"
import "net/http"
import "time"

type ServerTokenClaim struct {
	ExpiresAt int64 `json:"exp"`
	IssuedAt int64 `json:"iat"`
}

type json_out_token struct {
	AccessToken string `json:"access-token"`
	RefreshToken string `json:"refresh-token,omitempty"`
}

type json_out_server_conn struct {
	Id ConnId `json:"id"`
	ServerAddr string `json:"server-addr"`
	ClientAddr string `json:"client-addr"`
	ClientToken string `json:"client-token"`
	Routes []json_out_server_route `json:"routes"`
}

type json_out_server_route struct {
	Id RouteId `json:"id"`
	ClientPeerAddr string `json:"client-peer-addr"`
	ClientPeerName string `json:"client-peer-name"`
	ServerPeerOption string `json:"server-peer-option"`
	ServerPeerServiceAddr string `json:"server-peer-service-addr"` // actual listening address
	ServerPeerServiceNet string `json:"server-peer-service-net"`
}

type json_out_server_peer struct {
	Id PeerId `json:"id"`
	ServerPeerAddr string `json:"server-peer-addr"`
	ServerLocalAddr string `json:"server-local-addr"`
	ClientPeerAddr string `json:"client-peer-addr"`
	ClientLocalAddr string `json:"client-local-addr"`
}

type json_out_server_stats struct {
	json_out_go_stats

	ServerConns int64 `json:"server-conns"`
	ServerRoutes int64 `json:"server-routes"`
	ServerPeers int64 `json:"server-peers"`

	SshProxySessions int64 `json:"pxy-ssh-sessions"`
}

// ------------------------------------

type server_ctl struct {
	s *Server
	id string
	noauth bool // override the auth configuration if true
}

type server_ctl_token struct {
	server_ctl
}

type server_ctl_server_conns struct {
	server_ctl
}

type server_ctl_server_conns_id struct {
	server_ctl
}

type server_ctl_server_conns_id_routes struct {
	server_ctl
}

type server_ctl_server_conns_id_routes_id struct {
	server_ctl
}

type server_ctl_server_conns_id_routes_id_peers struct {
	server_ctl
}

type server_ctl_server_conns_id_routes_id_peers_id struct {
	server_ctl
}

type server_ctl_notices struct {
	server_ctl
}

type server_ctl_notices_id struct {
	server_ctl
}

type server_ctl_stats struct {
	server_ctl
}

// ------------------------------------

func (ctl *server_ctl) Id() string {
	return ctl.id
}

func (ctl *server_ctl) Cors(req *http.Request) bool {
	return ctl.s.Cfg.CtlCors
}

func (ctl *server_ctl) Authenticate(req *http.Request) (int, string) {
	if ctl.noauth || ctl.s.Cfg.CtlAuth == nil { return http.StatusOK, "" }
	return ctl.s.Cfg.CtlAuth.Authenticate(req)
}

// ------------------------------------

func (ctl *server_ctl_token) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var je *json.Encoder
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	switch req.Method {
		case http.MethodGet:
			var jwt *JWT[ServerTokenClaim]
			var claim ServerTokenClaim
			var tok string
			var now time.Time

			if s.Cfg.CtlAuth == nil || !s.Cfg.CtlAuth.Enabled || s.Cfg.CtlAuth.TokenRsaKey == nil {
				status_code = WriteJsonRespHeader(w, http.StatusForbidden)
				err = fmt.Errorf("auth not enabled or token rsa key not set")
				je.Encode(JsonErrmsg{Text: err.Error()})
				goto oops
			}

			now = time.Now()
			claim.IssuedAt = now.Unix()
			claim.ExpiresAt = now.Add(s.Cfg.CtlAuth.TokenTtl).Unix()
			jwt = NewJWT(s.Cfg.CtlAuth.TokenRsaKey, &claim)
			tok, err = jwt.SignRS512()
			if err != nil {
				status_code = WriteJsonRespHeader(w, http.StatusInternalServerError)
				je.Encode(JsonErrmsg{Text: err.Error()})
				goto oops
			}

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			err = je.Encode(json_out_token{ AccessToken: tok }) // TODO: refresh token
			if err != nil { goto oops }

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var je *json.Encoder
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	switch req.Method {
		case http.MethodGet:
			var cts *ServerConn
			var js []json_out_server_conn

			js = make([]json_out_server_conn, 0)
			s.cts_mtx.Lock()
			for _, cts = range s.cts_map {
				var r *ServerRoute
				var jsp []json_out_server_route

				jsp = make([]json_out_server_route, 0)
				cts.route_mtx.Lock()
				for _, r = range cts.route_map {
					jsp = append(jsp, json_out_server_route{
						Id: r.Id,
						ClientPeerAddr: r.PtcAddr,
						ClientPeerName: r.PtcName,
						ServerPeerServiceAddr: r.SvcAddr.String(),
						ServerPeerServiceNet: r.SvcPermNet.String(),
						ServerPeerOption: r.SvcOption.String(),
					})
				}
				js = append(js, json_out_server_conn{
					Id: cts.Id,
					ClientAddr: cts.RemoteAddr.String(),
					ServerAddr: cts.LocalAddr.String(),
					ClientToken: cts.ClientToken,
					Routes: jsp,
				})
				cts.route_mtx.Unlock()
			}
			s.cts_mtx.Unlock()

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(js); err != nil { goto oops }

		case http.MethodDelete:
			s.ReqStopAllServerConns()
			status_code = WriteEmptyRespHeader(w, http.StatusNoContent)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns_id) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var err error
	var je *json.Encoder
	var conn_id string
	var cts *ServerConn

	s = ctl.s
	je = json.NewEncoder(w)

	conn_id = req.PathValue("conn_id")
	cts, err = s.FindServerConnByIdStr(conn_id)
	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodGet:
			var r *ServerRoute
			var jsp []json_out_server_route
			var js *json_out_server_conn

			jsp = make([]json_out_server_route, 0)
			cts.route_mtx.Lock()
			for _, r = range cts.route_map {
				jsp = append(jsp, json_out_server_route{
					Id: r.Id,
					ClientPeerAddr: r.PtcAddr,
					ClientPeerName: r.PtcName,
					ServerPeerServiceAddr: r.SvcAddr.String(),
					ServerPeerServiceNet: r.SvcPermNet.String(),
					ServerPeerOption: r.SvcOption.String(),
				})
			}
			js = &json_out_server_conn{
				Id: cts.Id,
				ClientAddr: cts.RemoteAddr.String(),
				ServerAddr: cts.LocalAddr.String(),
				ClientToken: cts.ClientToken,
				Routes: jsp,
			}
			cts.route_mtx.Unlock()

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(js); err != nil { goto oops }

		case http.MethodDelete:
			//s.RemoveServerConn(cts)
			cts.ReqStop()
			status_code = WriteEmptyRespHeader(w, http.StatusNoContent)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns_id_routes) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var err error
	var conn_id string
	var je *json.Encoder
	var cts *ServerConn

	s = ctl.s
	je = json.NewEncoder(w)

	conn_id = req.PathValue("conn_id")
	cts, err = s.FindServerConnByIdStr(conn_id)
	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodGet:
			var r *ServerRoute
			var jsp []json_out_server_route

			jsp = make([]json_out_server_route, 0)
			cts.route_mtx.Lock()
			for _, r = range cts.route_map {
				jsp = append(jsp, json_out_server_route{
					Id: r.Id,
					ClientPeerAddr: r.PtcAddr,
					ClientPeerName: r.PtcName,
					ServerPeerServiceAddr: r.SvcAddr.String(),
					ServerPeerServiceNet: r.SvcPermNet.String(),
					ServerPeerOption: r.SvcOption.String(),
				})
			}
			cts.route_mtx.Unlock()

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(jsp); err != nil { goto oops }

		case http.MethodDelete:
			// direct removal causes quite some clean-up issues.
			// make stop request to all server routes and let the task stop by themselves
			//cts.RemoveAllServerRoutes()
			cts.ReqStopAllServerRoutes()
			status_code = WriteEmptyRespHeader(w, http.StatusNoContent)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns_id_routes_id) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var conn_id string
	var route_id string
	var je *json.Encoder
	var r *ServerRoute
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	if ctl.id == HS_ID_WPX && req.Method != http.MethodGet {
		// support the get method only, if invoked via the wpx endpoint
		status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
		goto done
	}

	conn_id = req.PathValue("conn_id")
	route_id = req.PathValue("route_id")
	r, err = s.FindServerRouteByIdStr(conn_id, route_id)
	if err != nil {
		/*
		if route_id == PORT_ID_MARKER && ctl.s.wpx_foreign_port_proxy_marker != nil {
			// don't care if the ctl call is from wpx or not. if the request
			// is by the port number(noted by route being PORT_ID_MARKER),
			// check if it's a foreign port
			var pi *ServerRouteProxyInfo
			// currenly, this is invoked via wpx only for ssh from xterm.html
			// ugly, but hard-code the type to "ssh" here for now...
			pi, err = ctl.s.wpx_foreign_port_proxy_maker("ssh", conn_id)
			if err == nil { r = proxy_info_to_server_route(pi) } // fake route
		}
		*/

		if ctl.id == HS_ID_WPX && route_id == PORT_ID_MARKER && ctl.s.wpx_foreign_port_proxy_maker != nil {
			var pi *ServerRouteProxyInfo
			// currenly, this is invoked via wpx only for ssh from xterm.html
			// ugly, but hard-code the type to "ssh" here for now...
			pi, err = ctl.s.wpx_foreign_port_proxy_maker("ssh", conn_id)
			if err == nil { r = proxy_info_to_server_route(pi) } // fake route
		}
	}

	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodGet:
			status_code = WriteJsonRespHeader(w, http.StatusOK)
			err = je.Encode(json_out_server_route{
				Id: r.Id,
				ClientPeerAddr: r.PtcAddr,
				ClientPeerName: r.PtcName,
				ServerPeerServiceAddr: r.SvcAddr.String(),
				ServerPeerServiceNet: r.SvcPermNet.String(),
				ServerPeerOption: r.SvcOption.String(),
			})
			if err != nil { goto oops }

		case http.MethodDelete:
			/*if r is foreign {
				// foreign route
				ctl.s.wpx_foreign_port_proxy_stopper(conn_id)
			} else {*/
				// native route
				r.ReqStop()
			//}
			status_code = WriteEmptyRespHeader(w, http.StatusNoContent)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns_id_routes_id_peers) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var conn_id string
	var route_id string
	var je *json.Encoder
	var r *ServerRoute
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	conn_id = req.PathValue("conn_id")
	route_id = req.PathValue("route_id")
	r, err = s.FindServerRouteByIdStr(conn_id, route_id)
	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodGet:
			var p *ServerPeerConn
			var jcp []json_out_server_peer

			jcp = make([]json_out_server_peer, 0)
			r.pts_mtx.Lock()
			for _, p = range r.pts_map {
				jcp = append(jcp, json_out_server_peer{
					Id: p.conn_id,
					ServerPeerAddr: p.conn.RemoteAddr().String(),
					ServerLocalAddr: p.conn.LocalAddr().String(),
					ClientPeerAddr: p.client_peer_raddr,
					ClientLocalAddr: p.client_peer_laddr,
				})
			}
			r.pts_mtx.Unlock()

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(jcp); err != nil { goto oops }

		case http.MethodDelete:
			r.ReqStopAllServerPeerConns()
			status_code = WriteEmptyRespHeader(w, http.StatusOK)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_server_conns_id_routes_id_peers_id) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var conn_id string
	var route_id string
	var peer_id string
	var je *json.Encoder
	var p *ServerPeerConn
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	conn_id = req.PathValue("conn_id")
	route_id = req.PathValue("route_id")
	peer_id = req.PathValue("peer_id")
	p, err = s.FindServerPeerConnByIdStr(conn_id, route_id, peer_id)
	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodGet:
			var jcp *json_out_server_peer

			jcp = &json_out_server_peer{
				Id: p.conn_id,
				ServerPeerAddr: p.conn.RemoteAddr().String(),
				ServerLocalAddr: p.conn.LocalAddr().String(),
				ClientPeerAddr: p.client_peer_raddr,
				ClientLocalAddr: p.client_peer_laddr,
			}

			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(jcp); err != nil { goto oops }

		case http.MethodDelete:
			p.ReqStop()
			status_code = WriteEmptyRespHeader(w, http.StatusNoContent)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_notices) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var cts *ServerConn
	var err error

	s = ctl.s

	switch req.Method {
		case http.MethodPost:
			var noti json_in_notice

			err = json.NewDecoder(req.Body).Decode(&noti)
			if err != nil {
				status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
				goto oops
			}

			// TODO: what if this loop takes too long? in that case, lock is held for long. think about how to handle this.
			s.cts_mtx.Lock()
			for _, cts = range s.cts_map {
				cts.pss.Send(MakeConnNoticePacket(noti.Text))
				// let's not care about an error when broacasting a notice to all connections
			}
			s.cts_mtx.Unlock()
			status_code = WriteJsonRespHeader(w, http.StatusOK)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_notices_id) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var conn_id string
	var cts *ServerConn
	var je *json.Encoder
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	conn_id = req.PathValue("conn_id") // server connection
	cts, err = s.FindServerConnByIdStr(conn_id)
	if err != nil {
		status_code = WriteJsonRespHeader(w, http.StatusNotFound)
		je.Encode(JsonErrmsg{Text: err.Error()})
		goto oops
	}

	switch req.Method {
		case http.MethodPost:
			var noti json_in_notice

			err = json.NewDecoder(req.Body).Decode(&noti)
			if err != nil {
				status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
				goto oops
			}

			// no check if noti.Text is empty as i want an empty message to be delivered too.
			err = cts.pss.Send(MakeConnNoticePacket(noti.Text))
			if err != nil {
				err = fmt.Errorf("failed to send conn_notice text '%s' to %s - %s", noti.Text, cts.RemoteAddr, err.Error())
				status_code = WriteJsonRespHeader(w, http.StatusInternalServerError)
				je.Encode(JsonErrmsg{Text: err.Error()})
				goto oops
			}

			status_code = WriteJsonRespHeader(w, http.StatusOK)

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusBadRequest)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}

// ------------------------------------

func (ctl *server_ctl_stats) ServeHTTP(w http.ResponseWriter, req *http.Request) (int, error) {
	var s *Server
	var status_code int
	var je *json.Encoder
	var err error

	s = ctl.s
	je = json.NewEncoder(w)

	switch req.Method {
		case http.MethodGet:
			var stats json_out_server_stats
			stats.from_runtime_stats()
			stats.ServerConns = s.stats.conns.Load()
			stats.ServerRoutes = s.stats.routes.Load()
			stats.ServerPeers = s.stats.peers.Load()
			stats.SshProxySessions = s.stats.ssh_proxy_sessions.Load()
			status_code = WriteJsonRespHeader(w, http.StatusOK)
			if err = je.Encode(stats); err != nil { goto oops }

		default:
			status_code = WriteEmptyRespHeader(w, http.StatusMethodNotAllowed)
	}

//done:
	return status_code, nil

oops:
	return status_code, err
}
