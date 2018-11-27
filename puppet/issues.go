package puppet

import "github.com/puppetlabs/go-issues/issue"

const (
	WF_FIELD_TYPE_MISMATCH          = `WF_FIELD_TYPE_MISMATCH`
	WF_ELEMENT_NOT_PARAMETER        = `WF_ELEMENT_NOT_PARAMETER`
	WF_NO_DEFINITION                = `WF_NO_DEFINITION`
	WF_NO_SERVER_BUILDER_IN_CONTEXT = `WF_NO_SERVER_BUILDER_IN_CONTEXT`
	WF_NOT_ACTIVITY                 = `WF_NOT_ACTIVITY`
	WF_INVALID_FUNCTION             = `WF_INVALID_FUNCTION`
	WF_MISSING_REQUIRED_FUNCION     = `WF_MISSING_REQUIRED_FUNCION`
)

func init() {
	issue.Hard(WF_FIELD_TYPE_MISMATCH, `expected activity %{field} to be a %{expected}, got %{actual}`)
	issue.Hard(WF_ELEMENT_NOT_PARAMETER, `expected activity %{field} element to be a Parameter, got %{type}`)
	issue.Hard(WF_NO_DEFINITION, `expected activity to contain a definition block`)
	issue.Hard(WF_NO_SERVER_BUILDER_IN_CONTEXT, `no ServerBuilder has been registered with the evaluation context`)
	issue.Hard(WF_INVALID_FUNCTION, `invalid function '%{function}'. Expected one of 'create', 'read', 'update', or 'delete'`)
	issue.Hard(WF_MISSING_REQUIRED_FUNCION, `missing required '%{function}'`)
	issue.Hard2(WF_NOT_ACTIVITY, `block may only contain workflow activites. %{actual} is not supported here`,
		issue.HF{`actual`: issue.A_anUc})
}