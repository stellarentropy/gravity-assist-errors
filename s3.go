package errors

import (
	"fmt"
)

// ErrObjectStorageUpload represents an error that occurs when there is a
// failure to upload data to an object storage service. This may be due to
// issues like network connectivity, access permissions, or other conditions
// that prevent successful data transfer to the storage service.
var ErrObjectStorageUpload = fmt.Errorf("error uploading data to object storage")

// ErrObjectStorageDownload represents an error encountered when attempting to
// retrieve data from an object storage service. This may be due to network
// disruptions, access permission problems, or other unforeseen factors that
// interfere with the ability to download the desired content.
var ErrObjectStorageDownload = fmt.Errorf("error downloading data from object storage")
