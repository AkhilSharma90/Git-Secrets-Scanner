package report

import "github.com/AkhilSharma90/Git-Secrets-Scanner/internal/report/secret"

type ReportWriter interface {
	WriteAll(s []*secret.Secret) error
	Close() error
}
