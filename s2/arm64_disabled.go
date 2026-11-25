//go:build arm64

package s2

// arm64 support has been removed from this fork. This file is compiled only
// for arm64, and the undefined reference below fails the build with a clear
// cause. Use upstream klauspost/compress for arm64 support.
var _ = arm64IsNotSupportedByThisFork
