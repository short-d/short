package fw

import "time"

type Timer interface {
	Now() time.Time
}
