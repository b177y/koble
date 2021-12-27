package sysrq

type SRQCommand byte

const (
	// Reboot will immediately reboot the system without syncing or
	// unmounting your disks.
	Reboot SRQCommand = 'b'

	// Crash will perform a system crash and a crashdump will be
	// taken if configured.
	Crash SRQCommand = 'c'

	// ShowLocks will show all locks that are held.
	ShowLocks SRQCommand = 'd'

	// TerminateAllTasks will send a SIGTERM to all processes, except for init.
	TerminateAllTasks SRQCommand = 'e'

	// MemoryFullOOMKill will call the OOM killer to kill a memory hog process,
	// but doesn't panic if nothing can be killed.
	MemoryFullOOMKill SRQCommand = 'f'

	// Help will display help
	Help SRQCommand = 'h'

	// KillAllTasks will send a SIGKILL to all processes, except for init.
	KillAllTasks SRQCommand = 'i'

	// ThawFilesystems forcibly "Just thaw it" - filesystems frozen by the
	// FIFREEZE ioctl.
	ThawFilesystems SRQCommand = 'j'

	// SAK (Secure Access Key) kills all programs on the current virtual
	// console.
	SAK SRQCommand = 'k'

	// ShowBacktraceAllActiveCPUs shows a stack backtrace for all active
	// CPUs.
	ShowBacktraceAllActiveCPUs SRQCommand = 'l'

	// ShowMemoryUsage dumps current memory info to your console.
	ShowMemoryUsage SRQCommand = 'm'

	// NiceAllRTTasks will make RT tasks nice-able.
	NiceAllRTTasks SRQCommand = 'n'

	// Poweroff shuts your system off (if configured and supported).
	Poweroff SRQCommand = 'o'

	// ShowRegisters dumps the current registers and flags to your console.
	ShowRegisters SRQCommand = 'p'

	// ShowAllTimers dumps per CPU lists of all armed hrtimers (but NOT
	// regular timer_list timers) and detailed information about all
	// clockevent devices.
	ShowAllTimers SRQCommand = 'q'

	// Unraw turns off keyboard raw mode and sets it to XLATE.
	Unraw SRQCommand = 'r'

	// Sync attempts to sync all mounted filesystems.
	Sync SRQCommand = 's'

	// ShowTaskStates dumps a list of current tasks and their information
	// to your console.
	ShowTaskStates SRQCommand = 't'

	// Unmount attempts to remount all mounted filesystems read-only.
	Unmount SRQCommand = 'u'

	// ShowBlockedTasks dumps tasks that are in uninterruptable (blocked)
	// state.
	ShowBlockedTasks SRQCommand = 'w'

	// DumpFtraceBuffer dumps the ftrace buffer.
	DumpFtraceBuffer SRQCommand = 'z'

	// Loglevel0 sets the console log level to 0. In this level, only
	// emergency messages like PANICs or OOPSes would make it to your
	// console.
	Loglevel0 SRQCommand = '0'

	// Loglevel1 sets the console log level to 1.
	Loglevel1 SRQCommand = '1'

	// Loglevel2 sets the console log level to 2.
	Loglevel2 SRQCommand = '2'

	// Loglevel3 sets the console log level to 3.
	Loglevel3 SRQCommand = '3'

	// Loglevel4 sets the console log level to 4.
	Loglevel4 SRQCommand = '4'

	// Loglevel5 sets the console log level to 5.
	Loglevel5 SRQCommand = '5'

	// Loglevel6 sets the console log level to 6.
	Loglevel6 SRQCommand = '6'

	// Loglevel7 sets the console log level to 7.
	Loglevel7 SRQCommand = '7'

	// Loglevel8 sets the console log level to 8.
	Loglevel8 SRQCommand = '8'

	// Loglevel9 sets the console log level to 9, the most verbose level.
	Loglevel9 SRQCommand = '9'
)
