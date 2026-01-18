package validation

// FieldSizeConstraints defines the maximum lengths for string fields across the application.
// These values correspond to:
// - Database VARCHAR constraints defined in migration-[13]-field_size_constraints.sql
// - Back-end validation tags in REST model structs (api/rest/model/*.go)
// - Front-end HTML maxlength attributes in web components
//
// IMPORTANT: Go struct tags cannot reference constants, so the values in REST model
// validation tags (e.g., `validate:"max=100"`) must be kept in sync manually.
// When changing these values, update all three locations:
// 1. Database migration
// 2. Go validation tags in model structs
// 3. HTML maxlength attributes
const (
	// Asset field constraints
	AssetTickerMaxLength = 40
	AssetNameMaxLength   = 100

	// Portfolio field constraints
	PortfolioNameMaxLength = 100

	// Allocation plan field constraints
	AllocationPlanNameMaxLength = 100
	AllocationPlanTypeMaxLength = 50

	// Portfolio allocation field constraints
	AllocationClassMaxLength = 100

	// Observation time field constraints
	ObservationTimeTagMaxLength = 100
)
