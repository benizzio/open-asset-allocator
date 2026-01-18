package validation

// FieldSizeConstraints defines the maximum lengths for string fields across the application.
// These values correspond to the database VARCHAR constraints defined in
// migration-[13]-field_size_constraints.sql and should be used for both
// front-end HTML maxlength attributes and back-end validation tags.
const (
	// Asset field constraints
	AssetTickerMaxLength = 20
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
