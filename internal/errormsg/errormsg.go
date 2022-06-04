package errmsg

import (
	"sched/pkg/meta"
)

var (
	// 1000 - 1999: system error
	ErrorInternalServer = meta.MetaErrorInternalServer.AppendMessage(-1000, "The server encountered an internal error or misconfiguration and was unable to complete your request.")
	ErrorForbidden      = meta.MetaErrorForbidden.AppendMessage(-1001, "You do not have permission to access this resource.")

	// 2000 - 2999: new cart, add item and remove item error
	// ErrorOpenOrderNotFound         = meta.MetaErrorBadRequest.AppendMessage(-2001, "open order not found.")
	// ErrorOrderSymbolNotMatch       = meta.MetaErrorBadRequest.AppendMessage(-2002, "symbol not match.")
	// ErrorTransactionStatusNotAllow = meta.MetaErrorBadRequest.AppendMessage(-2003, "transaction status not allow.")
	// ErrorTransactionNotFound       = meta.MetaErrorBadRequest.AppendMessage(-2004, "transaction not found.")

	// Schedule, 2000
)

func ErrorInvalidRequest(msg string) *meta.MetaError {
	return meta.MetaErrorBadRequest.AppendMessage(-1002, msg)
}
