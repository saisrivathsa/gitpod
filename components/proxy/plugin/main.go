package plugin

import (
	"bytes"
	"io"

	"github.com/caddyserver/caddy/v2"
	"github.com/mholt/caddy-l4/layer4"
)

func init() {
	caddy.RegisterModule(MatchSSH{})
}

// MatchSSH is able to match SSH connections.
type MatchSSH struct{}

// CaddyModule returns the Caddy module information.
func (MatchSSH) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "layer4.matchers.ssh",
		New: func() caddy.Module { return new(MatchSSH) },
	}
}

func (h *MatchSSH) Provision(ctx caddy.Context) error {
	return nil
}

// Handle handles the connections.
func (t *MatchSSH) Handle(cx *layer4.Connection, next layer4.Handler) error {
	/*
		// terminate TLS by performing the handshake
		tlsConn := tls.Server(cx.Conn, tlsCfg)
		err := tlsConn.Handshake()
		if err != nil {
			return err
		}
		t.logger.Debug("terminated TLS", zap.String("server_name", clientHello.ServerName))
	*/

	return next.Handle(cx)
}

// Match returns true if the connection looks like SSH.
func (m MatchSSH) Match(cx *layer4.Connection) (bool, error) {
	p := make([]byte, len(sshPrefix))
	n, err := io.ReadFull(cx, p)
	if err != nil || n < len(sshPrefix) {
		return false, nil
	}
	return bytes.Equal(p, sshPrefix), nil
}

var sshPrefix = []byte("SSH-")

// Interface guards
var (
	_ caddy.Provisioner  = (*MatchSSH)(nil)
	_ layer4.NextHandler = (*MatchSSH)(nil)
	_ layer4.ConnMatcher = (*MatchSSH)(nil)
)
