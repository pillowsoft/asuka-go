//go:build darwin && !ios

package darwin

import (
	. "github.com/egoist/asuka/internal/darwin"
)

// commonModes is CoreFoundation's kCFRunLoopCommonModes, the constant
// itself. A run loop knows the common modes by that pointer: a new
// NSString with the same text names an
// ordinary mode nothing runs, so a timer or display link added with it
// never fires.
func commonModes() ID { return Const(CoreFoundation(), "kCFRunLoopCommonModes") }
