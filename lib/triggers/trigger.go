package trigger

type Trigger interface {
	Trigger() bool
	Triggerd() bool
	ClearTrigger()
}
