package proxy

import (
	"errors"
	"github.com/miekg/dns"
	"net"
	"github.com/mageddo/dns-proxy-server/events/local"
	"github.com/mageddo/log"
)

type LocalDnsSolver struct {

}

func (LocalDnsSolver) Solve(question dns.Question) (*dns.Msg, error) {

	key := question.Name[:len(question.Name)-1]
	activeEnv := events.GetConfiguration().GetActiveEnv()
	hostname := activeEnv.GetHostname(key)
	if  hostname != nil {
		rr := &dns.A{
			Hdr: dns.RR_Header{Name: question.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
			A: net.IPv4(hostname.Ip[0], hostname.Ip[1], hostname.Ip[2], hostname.Ip[3]),
		}

		m := new(dns.Msg)
		m.Answer = append(m.Answer, rr)
		log.Logger.Infof("m=solve, status=success, solver=local")
		return m, nil
	}
	return nil, errors.New("hostname not found")
}
