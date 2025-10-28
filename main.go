/*
 */
package main

/*
#include <stdlib.h>
#include "libSzConfigMgr.h"
#include "libSz.h"
#include "szhelpers/SzLang_helpers.h"
#cgo linux CFLAGS: -g -I/opt/senzing/er/sdk/c
#cgo linux LDFLAGS: -L/opt/senzing/er/lib -lSz
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/senzing-garage/go-helpers/settings"
)

const (
	noError      = 0
	logLevel     = 0
	instanceName = "test"
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// Build SENZING_ENGINE_CONFIGURATION_JSON.

	senzingEngineSettings, err := settings.BuildSimpleSettingsUsingEnvVars()
	panicOnError(err)
	fmt.Printf("SENZING_ENGINE_CONFIGURATION_JSON: %s\n", senzingEngineSettings)

	// Create Configuration manager.

	err = createEngine(instanceName, senzingEngineSettings, logLevel)
	panicOnError(err)

	err = createConfigManager(instanceName, senzingEngineSettings, logLevel)
	panicOnError(err)

	// Destroy Configuration manager.

	err = destroyConfigManager()
	panicOnError(err)

	err = destroyConfigManager()
	panicOnError(err)

	// Create Configuration manager.

	err = createConfigManager(instanceName, senzingEngineSettings, logLevel)
	panicOnError(err)

	// Destroy Configuration manager.

	err = destroyConfigManager()
	panicOnError(err)

}

// ----------------------------------------------------------------------------
// Supporting functions
// ----------------------------------------------------------------------------

func createEngine(
	instanceName string,
	settings string,
	verboseLogging int64,
) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var err error

	moduleNameForC := C.CString(instanceName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(settings)
	defer C.free(unsafe.Pointer(iniParamsForC))

	result := C.Sz_init(moduleNameForC, iniParamsForC, C.int64_t(verboseLogging))
	if result != noError {
		err = fmt.Errorf("Sz_init failed. Settings: %s", settings)
	}

	return err
}

func createConfigManager(
	instanceName string,
	settings string,
	verboseLogging int64,
) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var err error

	moduleNameForC := C.CString(instanceName)
	defer C.free(unsafe.Pointer(moduleNameForC))

	iniParamsForC := C.CString(settings)
	defer C.free(unsafe.Pointer(iniParamsForC))

	result := C.SzConfigMgr_init(moduleNameForC, iniParamsForC, C.int64_t(verboseLogging))
	if result != noError {
		err = fmt.Errorf("SzConfigMgr_init failed. Settings: %s", settings)
	}

	return err
}

func destroyConfigManager() error {
	var err error

	result := C.SzConfigMgr_destroy()
	if result != noError {
		err = fmt.Errorf("SzConfigMgr_destroy failed")
	}
	return err
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
