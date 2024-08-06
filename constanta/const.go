package constanta

type Info string

const (
	c_DARK_RED = "\u001b[91m"
	c_YELLOW   = "\u001b[93m"
	c_GREEN    = "\u001b[92m"
	c_RESET    = "\u001b[0m"

	INFO  Info = c_GREEN + "[INFO]" + c_RESET
	DEBUG Info = c_YELLOW + "[DEBUG]" + c_RESET
	ERROR Info = c_DARK_RED + "[ERROR]" + c_RESET
	FATAL Info = c_DARK_RED + "[FATAL]" + c_RESET
	PANIC Info = c_DARK_RED + "[PANIC]" + c_RESET
)
