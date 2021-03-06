package blockstorage

import (
	"context"
)

// Provider abstracts actions on underlying storage
type Provider interface {
	// Type returns the underlying storage type
	Type() Type
	// Volume operations
	VolumeCreate(context.Context, Volume) (*Volume, error)
	VolumeCreateFromSnapshot(ctx context.Context, snapshot Snapshot, tags map[string]string) (*Volume, error)
	VolumeDelete(context.Context, *Volume) error
	VolumeGet(ctx context.Context, id string, zone string) (*Volume, error)
	// Snapshot operations
	SnapshotCopy(ctx context.Context, from Snapshot, to Snapshot) (*Snapshot, error)
	SnapshotCreate(ctx context.Context, volume Volume, tags map[string]string) (*Snapshot, error)
	SnapshotCreateWaitForCompletion(context.Context, *Snapshot) error
	SnapshotDelete(context.Context, *Snapshot) error
	SnapshotGet(ctx context.Context, id string) (*Snapshot, error)
	// Others
	SetTags(ctx context.Context, resource interface{}, tags map[string]string) error
	VolumesList(ctx context.Context, tags map[string]string, zone string) ([]*Volume, error)
	SnapshotsList(ctx context.Context, tags map[string]string) ([]*Snapshot, error)
}

// RestoreTargeter implements the SnapshotRestoreTargets method
type RestoreTargeter interface {
	// SnapshotRestoreTargets returns whether a snapshot can be restored globally.
	// If not globally restorable, returns a map of the regions and zones to which snapshot can be restored.
	SnapshotRestoreTargets(context.Context, *Snapshot) (global bool, regionsAndZones map[string][]string, err error)
}
