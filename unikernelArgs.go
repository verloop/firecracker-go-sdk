// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package firecracker

import (
	"fmt"
	"strings"
)

func parseUnikernelArgs(unikernelArgs UnikernelArgs, ipBootParam string) string {
	netdevIP := convertIPtoNetdevIP(ipBootParam)
	return fmt.Sprintf("%s %s %s -- %s", unikernelArgs.KernelName, netdevIP, unikernelArgs.ExtraArgs, unikernelArgs.UserArgs)
}

func convertIPtoNetdevIP(ipBootParam string) string {
	parts := strings.Split(ipBootParam, ":")
	ip := parts[0]      // 192.168.124.75 // ip address of machine
	gateway := parts[2] // 192.168.124.1 (index 2 for the gateway)
	// subnetMask := parts[3] // Not required right now for the new format

	// Construct the refactored config
	refactoredConfig := fmt.Sprintf("netdev.ip=%s/24:%s:::", ip, gateway)
	return refactoredConfig
}
