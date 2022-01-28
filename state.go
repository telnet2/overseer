package overseer

const (
	INITIAL = iota
	IDLE
	STARTING
	RUNNING
	STOPPING
	INTERRUPT // final state (used when stopped or signaled)
	FINISHED  // final state (used then was a natural exit)
	FATAL     // final state (used when was an error while starting)
	// UNKNOWN ?
)

const (
	IDLE_STATE     = "idle"
	DONE_STATE     = "done"
	FATAL_STATE    = "fatal"
	RUNNING_STATE  = "running"
	STOPPING_STATE = "stopping"
	UNKNOWN_STATE  = "unknown"
)

// CmdState represents a Cmd state
type CmdState uint8

func (p CmdState) String() string {
	switch p {
	case INITIAL:
		return "initial"
	case STARTING:
		return "starting"
	case RUNNING:
		return "running"
	case STOPPING:
		return "stopping"
	case INTERRUPT:
		return "interrupted"
	case FINISHED:
		return "finished"
	case FATAL:
		return "fatal"
	default:
		return "unknown"
	}
}

// OvrState represents a Overseer state
type OvrState uint8

func (p OvrState) String() string {
	switch p {
	case IDLE:
		return IDLE_STATE
	case RUNNING:
		return RUNNING_STATE
	case STOPPING:
		return STOPPING_STATE
	case FATAL:
		return FATAL_STATE
	default:
		return UNKNOWN_STATE
	}
}
