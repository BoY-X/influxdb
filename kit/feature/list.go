// Code generated by the feature package; DO NOT EDIT.

package feature

var appMetrics = MakeBoolFlag(
	"App Metrics",
	"appMetrics",
	"Bucky, Monitoring Team",
	false,
	Permanent,
	true,
)

// AppMetrics - Send UI Telementry to Tools cluster - should always be false in OSS
func AppMetrics() BoolFlag {
	return appMetrics
}

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var communityTemplates = MakeBoolFlag(
	"Community Templates",
	"communityTemplates",
	"Bucky, Johnny Steenbergen (Berg)",
	false,
	Permanent,
	true,
)

// CommunityTemplates - Replace current template uploading functionality with community driven templates
func CommunityTemplates() BoolFlag {
	return communityTemplates
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var pushDownWindowAggregateCount = MakeBoolFlag(
	"Push Down Window Aggregate Count",
	"pushDownWindowAggregateCount",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateCount - Enable Count variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateCount() BoolFlag {
	return pushDownWindowAggregateCount
}

var pushDownWindowAggregateSum = MakeBoolFlag(
	"Push Down Window Aggregate Sum",
	"pushDownWindowAggregateSum",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateSum - Enable Sum variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateSum() BoolFlag {
	return pushDownWindowAggregateSum
}

var pushDownWindowAggregateMin = MakeBoolFlag(
	"Push Down Window Aggregate Min",
	"pushDownWindowAggregateMin",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateMin - Enable Min variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateMin() BoolFlag {
	return pushDownWindowAggregateMin
}

var pushDownWindowAggregateMax = MakeBoolFlag(
	"Push Down Window Aggregate Max",
	"pushDownWindowAggregateMax",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateMax - Enable Max variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateMax() BoolFlag {
	return pushDownWindowAggregateMax
}

var pushDownWindowAggregateMean = MakeBoolFlag(
	"Push Down Window Aggregate Mean",
	"pushDownWindowAggregateMean",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateMean - Enable Mean variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateMean() BoolFlag {
	return pushDownWindowAggregateMean
}

var groupWindowAggregateTranspose = MakeBoolFlag(
	"Group Window Aggregate Transpose",
	"groupWindowAggregateTranspose",
	"Query Team",
	false,
	Temporary,
	false,
)

// GroupWindowAggregateTranspose - Enables the GroupWindowAggregateTransposeRule for all enabled window aggregates
func GroupWindowAggregateTranspose() BoolFlag {
	return groupWindowAggregateTranspose
}

var newAuth = MakeBoolFlag(
	"New Auth Package",
	"newAuth",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewAuthPackage - Enables the refactored authorization api
func NewAuthPackage() BoolFlag {
	return newAuth
}

var newLabels = MakeBoolFlag(
	"New Label Package",
	"newLabels",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewLabelPackage - Enables the refactored labels api
func NewLabelPackage() BoolFlag {
	return newLabels
}

var hydratevars = MakeBoolFlag(
	"New Hydrate Vars Functionality",
	"hydratevars",
	"Ariel Salem / Monitoring Team",
	false,
	Temporary,
	true,
)

// NewHydrateVarsFunctionality - Enables a minimalistic variable hydration
func NewHydrateVarsFunctionality() BoolFlag {
	return hydratevars
}

var memoryOptimizedFill = MakeBoolFlag(
	"Memory Optimized Fill",
	"memoryOptimizedFill",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedFill - Enable the memory optimized fill()
func MemoryOptimizedFill() BoolFlag {
	return memoryOptimizedFill
}

var memoryOptimizedSchemaMutation = MakeBoolFlag(
	"Memory Optimized Schema Mutation",
	"memoryOptimizedSchemaMutation",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedSchemaMutation - Enable the memory optimized schema mutation functions
func MemoryOptimizedSchemaMutation() BoolFlag {
	return memoryOptimizedSchemaMutation
}

var urmFreeTasks = MakeBoolFlag(
	"Urm Free Tasks",
	"urmFreeTasks",
	"Lyon Hill",
	false,
	Temporary,
	false,
)

// UrmFreeTasks - allow task system to operate without creating additional urms
func UrmFreeTasks() BoolFlag {
	return urmFreeTasks
}

var simpleTaskOptionsExtraction = MakeBoolFlag(
	"Simple Task Options Extraction",
	"simpleTaskOptionsExtraction",
	"Brett Buddin",
	false,
	Temporary,
	false,
)

// SimpleTaskOptionsExtraction - Simplified task options extraction to avoid undefined functions when saving tasks
func SimpleTaskOptionsExtraction() BoolFlag {
	return simpleTaskOptionsExtraction
}

var useUserPermission = MakeBoolFlag(
	"Use User Permission",
	"useUserPermission",
	"Lyon Hill",
	false,
	Temporary,
	false,
)

// UseUserPermission - When enabled we will use the new user service permission function
func UseUserPermission() BoolFlag {
	return useUserPermission
}

var mergeFiltersRule = MakeBoolFlag(
	"Merged Filters Rule",
	"mergeFiltersRule",
	"Query Team",
	false,
	Temporary,
	false,
)

// MergedFiltersRule - Create one filter combining multiple single return statements
func MergedFiltersRule() BoolFlag {
	return mergeFiltersRule
}

var notebooks = MakeBoolFlag(
	"Notebooks",
	"notebooks",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// Notebooks - Determine if the notebook feature's route and navbar icon are visible to the user
func Notebooks() BoolFlag {
	return notebooks
}

var all = []Flag{
	appMetrics,
	backendExample,
	communityTemplates,
	frontendExample,
	pushDownWindowAggregateCount,
	pushDownWindowAggregateSum,
	pushDownWindowAggregateMin,
	pushDownWindowAggregateMax,
	pushDownWindowAggregateMean,
	groupWindowAggregateTranspose,
	newAuth,
	newLabels,
	hydratevars,
	memoryOptimizedFill,
	memoryOptimizedSchemaMutation,
	urmFreeTasks,
	simpleTaskOptionsExtraction,
	useUserPermission,
	mergeFiltersRule,
	notebooks,
}

var byKey = map[string]Flag{
	"appMetrics":                    appMetrics,
	"backendExample":                backendExample,
	"communityTemplates":            communityTemplates,
	"frontendExample":               frontendExample,
	"pushDownWindowAggregateCount":  pushDownWindowAggregateCount,
	"pushDownWindowAggregateSum":    pushDownWindowAggregateSum,
	"pushDownWindowAggregateMin":    pushDownWindowAggregateMin,
	"pushDownWindowAggregateMax":    pushDownWindowAggregateMax,
	"pushDownWindowAggregateMean":   pushDownWindowAggregateMean,
	"groupWindowAggregateTranspose": groupWindowAggregateTranspose,
	"newAuth":                       newAuth,
	"newLabels":                     newLabels,
	"hydratevars":                   hydratevars,
	"memoryOptimizedFill":           memoryOptimizedFill,
	"memoryOptimizedSchemaMutation": memoryOptimizedSchemaMutation,
	"urmFreeTasks":                  urmFreeTasks,
	"simpleTaskOptionsExtraction":   simpleTaskOptionsExtraction,
	"useUserPermission":             useUserPermission,
	"mergeFiltersRule":              mergeFiltersRule,
	"notebooks":                     notebooks,
}
