package gcp

import (
	"context"
	"io"

	"github.com/stellarentropy/gravity-assist-common/metrics/datacounter"

	"github.com/stellarentropy/gravity-assist-common/errors"

	"cloud.google.com/go/storage"
)

// storageClient manages interactions with Google Cloud Storage, facilitating
// operations such as generating signed URLs, uploading and downloading data to
// and from buckets, and establishing data transfer streams.
var storageClient *storage.Client

// init establishes a connection to Google Cloud Storage and initializes the
// global storageClient variable, panicking if the client creation fails.
func init() {
	// Declare an error variable
	var err error

	// Create a new storage client and assign it to the global storageClient variable
	// The context.Background() is used as the context for this operation
	storageClient, err = storage.NewClient(context.Background())

	// If an error occurred while creating the client, panic and stop the execution
	if err != nil {
		panic(err)
	}
}

// Upload transfers data from an [io.Reader] to a specified path within a Google
// Cloud Storage bucket, respecting the context provided for the operation. It
// requires a context for cancelation and deadline control, the name of the
// bucket, the destination path within that bucket, and the data source as an
// [io.Reader]. In case of success, it returns nil; otherwise, it returns an
// error indicating what went wrong during the upload process.
func Upload(ctx context.Context, component string, bucket string, path string, r io.Reader) error {
	// Get an io.WriteCloser for the specified bucket and path
	wc, err := GetUploadWriter(ctx, component, bucket, path)
	if err != nil {
		return err
	}

	// Copy the data from the provided io.Reader to the io.WriteCloser
	if _, err := io.Copy(wc, r); err != nil {
		// If there's an error during the copy, wrap it with a custom error and return
		return errors.Wrap(errors.ErrObjectStorageUpload, err)
	}

	// Close the io.WriteCloser after the copy operation
	if err := wc.Close(); err != nil {
		// If there's an error during the close operation, wrap it with a custom error and return
		return errors.Wrap(errors.ErrObjectStorageUpload, err)
	}

	// If everything went well, return nil indicating a successful operation
	return nil
}

// GetUploadWriter creates an [io.WriteCloser] for uploading data to a specified
// path within a Google Cloud Storage bucket. It takes a context, bucket name,
// and object path as arguments to initiate the upload process. On successful
// creation of the writer, it returns the writer along with any error that may
// have occurred during setup.
func GetUploadWriter(ctx context.Context, component string, bucket string, path string) (io.WriteCloser, error) {
	// Create a new writer for the specified bucket and object path
	wc := storageClient.Bucket(bucket).Object(path).NewWriter(ctx)

	// Create a new counter for the writer to track the amount of data written
	counter := datacounter.NewObjectStorageWriterCounter(ctx, component, wc, storageClient)

	// Return the counter (which also acts as a writer) and nil for the error
	return counter, nil
}

// Download retrieves content from a specified path in a Google Cloud Storage
// bucket and writes it to the provided [io.Writer]. It ensures the data is
// fetched and copied correctly, handling any errors that may arise during the
// operation. This function requires a context for managing the request's
// lifetime, the name of the bucket, the object's path within that bucket, and
// an [io.Writer] to which the data will be written. If any errors occur while
// setting up the reader, transferring the data, or closing the connection, they
// are returned.
func Download(ctx context.Context, component string, bucket string, path string, w io.Writer) error {
	// Get an io.ReadCloser for the specified bucket and path
	rc, err := GetDownloadReader(ctx, component, bucket, path, false)
	if err != nil {
		// If there's an error during the reader creation, return the error
		return err
	}
	// Ensure the reader is closed after the operation
	defer func() { _ = rc.Close() }()

	// Copy the data from the io.ReadCloser to the provided io.Writer
	if _, err := io.Copy(w, rc); err != nil {
		// If there's an error during the copy, wrap it with a custom error and return
		return errors.Wrap(errors.ErrObjectStorageDownload, err)
	}

	// If everything went well, return nil indicating a successful operation
	return nil
}

// GetDownloadReader creates and returns an [io.ReadCloser] for reading data
// from a specified object in a Google Cloud Storage bucket. It takes a context
// for managing the request's lifetime, the name of the bucket, the object's
// path within that bucket, and a seeker flag indicating whether seeking
// operations are supported. In the event of an error during reader creation,
// the error is returned along with a nil reader.
func GetDownloadReader(ctx context.Context, component string, bucket string, path string, seeker bool) (io.ReadCloser, error) {
	// Get a handle to the specified object in the bucket
	handle := storageClient.Bucket(bucket).Object(path)

	// Create a new reader for the object
	rc, err := handle.NewReader(ctx)
	if err != nil {
		// If there's an error during the reader creation, wrap it with a custom error and return
		return nil, errors.Wrap(errors.ErrObjectStorageDownload, err)
	}

	// Create a new counter for the reader to track the amount of data read
	counter := datacounter.NewObjectStorageReaderCounter(ctx, component, rc, storageClient, handle, seeker)

	// Return the counter
	return counter, nil
}
