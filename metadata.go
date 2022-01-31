package main

// metadata stores the metadata associated with
// a single migration in the database.
type MigrationMetadata struct {
	ID           string // ID of the migration.
	Depth        int    // Depth of this migration in the tree
	GitShortHash string // Short hash of the current commit.
	PreviousID   string // Previous migration ID.
	IsActive     bool   // Has this migration been run?
	IsCurrent    bool   // Is this migration the current one?
}
