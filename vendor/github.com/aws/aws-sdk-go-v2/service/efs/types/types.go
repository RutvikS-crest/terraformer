// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Provides a description of an EFS file system access point.
type AccessPointDescription struct {

	// The unique Amazon Resource Name (ARN) associated with the access point.
	AccessPointArn *string

	// The ID of the access point, assigned by Amazon EFS.
	AccessPointId *string

	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string

	// The ID of the EFS file system that the access point applies to.
	FileSystemId *string

	// Identifies the lifecycle phase of the access point.
	LifeCycleState LifeCycleState

	// The name of the access point. This is the value of the Name tag.
	Name *string

	// Identified the AWS account that owns the access point resource.
	OwnerId *string

	// The full POSIX identity, including the user ID, group ID, and secondary group
	// IDs on the access point that is used for all file operations by NFS clients
	// using the access point.
	PosixUser *PosixUser

	// The directory on the Amazon EFS file system that the access point exposes as the
	// root directory to NFS clients using the access point.
	RootDirectory *RootDirectory

	// The tags associated with the access point, presented as an array of Tag objects.
	Tags []Tag
}

// The backup policy for the file system used to create automatic daily backups. If
// status has a value of ENABLED, the file system is being automatically backed up.
// For more information, see Automatic backups
// (https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html#automatic-backups).
type BackupPolicy struct {

	// Describes the status of the file system's backup policy.
	//
	// * ENABLED - EFS is
	// automatically backing up the file system.>
	//
	// * ENABLING - EFS is turning on
	// automatic backups for the file system.
	//
	// * DISABLED - automatic back ups are
	// turned off for the file system.
	//
	// * DISABLING - EFS is turning off automatic
	// backups for the file system.
	//
	// This member is required.
	Status Status
}

// Required if the RootDirectory > Path specified does not exist. Specifies the
// POSIX IDs and permissions to apply to the access point's RootDirectory > Path.
// If the access point root directory does not exist, EFS creates it with these
// settings when a client connects to the access point. When specifying
// CreationInfo, you must include values for all properties. Amazon EFS creates a
// root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and
// permissions for the directory. If you do not provide this information, Amazon
// EFS does not create the root directory. If the root directory does not exist,
// attempts to mount using the access point will fail. If you do not provide
// CreationInfo and the specified RootDirectory does not exist, attempts to mount
// the file system using the access point will fail.
type CreationInfo struct {

	// Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from
	// 0 to 2^32 (4294967295).
	//
	// This member is required.
	OwnerGid *int64

	// Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0
	// to 2^32 (4294967295).
	//
	// This member is required.
	OwnerUid *int64

	// Specifies the POSIX permissions to apply to the RootDirectory, in the format of
	// an octal number representing the file's mode bits.
	//
	// This member is required.
	Permissions *string
}

// A description of the file system.
type FileSystemDescription struct {

	// The time that the file system was created, in seconds (since
	// 1970-01-01T00:00:00Z).
	//
	// This member is required.
	CreationTime *time.Time

	// The opaque string specified in the request.
	//
	// This member is required.
	CreationToken *string

	// The ID of the file system, assigned by Amazon EFS.
	//
	// This member is required.
	FileSystemId *string

	// The lifecycle phase of the file system.
	//
	// This member is required.
	LifeCycleState LifeCycleState

	// The current number of mount targets that the file system has. For more
	// information, see CreateMountTarget.
	//
	// This member is required.
	NumberOfMountTargets int32

	// The AWS account that created the file system. If the file system was created by
	// an IAM user, the parent account to which the user belongs is the owner.
	//
	// This member is required.
	OwnerId *string

	// The performance mode of the file system.
	//
	// This member is required.
	PerformanceMode PerformanceMode

	// The latest known metered size (in bytes) of data stored in the file system, in
	// its Value field, and the time at which that size was determined in its Timestamp
	// field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of a
	// consistent snapshot of the file system, but it is eventually consistent when
	// there are no writes to the file system. That is, SizeInBytes represents actual
	// size only if the file system is not modified for a period longer than a couple
	// of hours. Otherwise, the value is not the exact size that the file system was at
	// any point in time.
	//
	// This member is required.
	SizeInBytes *FileSystemSize

	// The tags associated with the file system, presented as an array of Tag objects.
	//
	// This member is required.
	Tags []Tag

	// The unique and consistent identifier of the Availability Zone in which the file
	// system's One Zone storage classes exist. For example, use1-az1 is an
	// Availability Zone ID for the us-east-1 AWS Region, and it has the same location
	// in every AWS account.
	AvailabilityZoneId *string

	// Describes the AWS Availability Zone in which the file system is located, and is
	// valid only for file systems using One Zone storage classes. For more
	// information, see Using EFS storage classes
	// (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the Amazon
	// EFS User Guide.
	AvailabilityZoneName *string

	// A Boolean value that, if true, indicates that the file system is encrypted.
	Encrypted *bool

	// The Amazon Resource Name (ARN) for the EFS file system, in the format
	// arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id . Example
	// with sample data:
	// arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-01234567
	FileSystemArn *string

	// The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that
	// was used to protect the encrypted file system.
	KmsKeyId *string

	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem. If the file system has a Name tag, Amazon EFS returns the
	// value in this field.
	Name *string

	// The amount of provisioned throughput, measured in MiB/s, for the file system.
	// Valid for file systems using ThroughputMode set to provisioned.
	ProvisionedThroughputInMibps *float64

	// Displays the file system's throughput mode. For more information, see Throughput
	// modes
	// (https://docs.aws.amazon.com/efs/latest/ug/performance.html#throughput-modes) in
	// the Amazon EFS User Guide.
	ThroughputMode ThroughputMode
}

// The latest known metered size (in bytes) of data stored in the file system, in
// its Value field, and the time at which that size was determined in its Timestamp
// field. The value doesn't represent the size of a consistent snapshot of the file
// system, but it is eventually consistent when there are no writes to the file
// system. That is, the value represents the actual size only if the file system is
// not modified for a period longer than a couple of hours. Otherwise, the value is
// not necessarily the exact size the file system was at any instant in time.
type FileSystemSize struct {

	// The latest known metered size (in bytes) of data stored in the file system.
	//
	// This member is required.
	Value int64

	// The time at which the size of data, returned in the Value field, was determined.
	// The value is the integer number of seconds since 1970-01-01T00:00:00Z.
	Timestamp *time.Time

	// The latest known metered size (in bytes) of data stored in the Infrequent Access
	// storage class.
	ValueInIA *int64

	// The latest known metered size (in bytes) of data stored in the Standard storage
	// class.
	ValueInStandard *int64
}

// Describes a policy used by EFS lifecycle management to transition files to the
// Infrequent Access (IA) storage class.
type LifecyclePolicy struct {

	// A value that describes the period of time that a file is not accessed, after
	// which it transitions to the IA storage class. Metadata operations such as
	// listing the contents of a directory don't count as file access events.
	TransitionToIA TransitionToIARules
}

// Provides a description of a mount target.
type MountTargetDescription struct {

	// The ID of the file system for which the mount target is intended.
	//
	// This member is required.
	FileSystemId *string

	// Lifecycle state of the mount target.
	//
	// This member is required.
	LifeCycleState LifeCycleState

	// System-assigned mount target ID.
	//
	// This member is required.
	MountTargetId *string

	// The ID of the mount target's subnet.
	//
	// This member is required.
	SubnetId *string

	// The unique and consistent identifier of the Availability Zone that the mount
	// target resides in. For example, use1-az1 is an AZ ID for the us-east-1 Region
	// and it has the same location in every AWS account.
	AvailabilityZoneId *string

	// The name of the Availability Zone in which the mount target is located.
	// Availability Zones are independently mapped to names for each AWS account. For
	// example, the Availability Zone us-east-1a for your AWS account might not be the
	// same location as us-east-1a for another AWS account.
	AvailabilityZoneName *string

	// Address at which the file system can be mounted by using the mount target.
	IpAddress *string

	// The ID of the network interface that Amazon EFS created when it created the
	// mount target.
	NetworkInterfaceId *string

	// AWS account ID that owns the resource.
	OwnerId *string

	// The virtual private cloud (VPC) ID that the mount target is configured in.
	VpcId *string
}

// The full POSIX identity, including the user ID, group ID, and any secondary
// group IDs, on the access point that is used for all file system operations
// performed by NFS clients using the access point.
type PosixUser struct {

	// The POSIX group ID used for all file system operations using this access point.
	//
	// This member is required.
	Gid *int64

	// The POSIX user ID used for all file system operations using this access point.
	//
	// This member is required.
	Uid *int64

	// Secondary POSIX group IDs used for all file system operations using this access
	// point.
	SecondaryGids []int64
}

// Specifies the directory on the Amazon EFS file system that the access point
// provides access to. The access point exposes the specified file system path as
// the root directory of your file system to applications using the access point.
// NFS clients using the access point can only access data in the access point's
// RootDirectory and it's subdirectories.
type RootDirectory struct {

	// (Optional) Specifies the POSIX IDs and permissions to apply to the access
	// point's RootDirectory. If the RootDirectory > Path specified does not exist, EFS
	// creates the root directory using the CreationInfo settings when a client
	// connects to an access point. When specifying the CreationInfo, you must provide
	// values for all properties. If you do not provide CreationInfo and the specified
	// RootDirectory > Path does not exist, attempts to mount the file system using the
	// access point will fail.
	CreationInfo *CreationInfo

	// Specifies the path on the EFS file system to expose as the root directory to NFS
	// clients using the access point to access the EFS file system. A path can have up
	// to four subdirectories. If the specified path does not exist, you are required
	// to provide the CreationInfo.
	Path *string
}

// A tag is a key-value pair. Allowed characters are letters, white space, and
// numbers that can be represented in UTF-8, and the following characters: + - = .
// _ : /
type Tag struct {

	// The tag key (String). The key can't start with aws:.
	//
	// This member is required.
	Key *string

	// The value of the tag key.
	//
	// This member is required.
	Value *string
}
