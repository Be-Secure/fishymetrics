/*
 * Copyright 2024 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xl420

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics map[string]*prometheus.GaugeVec

func newServerMetric(metricName string, docString string, constLabels prometheus.Labels, labelNames []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        metricName,
			Help:        docString,
			ConstLabels: constLabels,
		},
		labelNames,
	)
}

func NewDeviceMetrics() *map[string]*metrics {
	var (
		ThermalMetrics = &metrics{
			"fanSpeed":          newServerMetric("xl420_thermal_fan_speed", "Current fan speed in the unit of percentage, possible values are 0 - 100", nil, []string{"name"}),
			"fanStatus":         newServerMetric("xl420_thermal_fan_status", "Current fan status 1 = OK, 0 = BAD", nil, []string{"name"}),
			"sensorTemperature": newServerMetric("xl420_thermal_sensor_temperature", "Current sensor temperature reading in Celsius", nil, []string{"name"}),
			"sensorStatus":      newServerMetric("xl420_thermal_sensor_status", "Current sensor status 1 = OK, 0 = BAD", nil, []string{"name"}),
		}

		PowerMetrics = &metrics{
			"supplyOutput":        newServerMetric("xl420_power_supply_output", "Power supply output in watts", nil, []string{"memberId", "sparePartNumber"}),
			"supplyStatus":        newServerMetric("xl420_power_supply_status", "Current power supply status 1 = OK, 0 = BAD", nil, []string{"memberId", "sparePartNumber"}),
			"supplyTotalConsumed": newServerMetric("xl420_power_supply_total_consumed", "Total output of all power supplies in watts", nil, []string{"memberId"}),
			"supplyTotalCapacity": newServerMetric("xl420_power_supply_total_capacity", "Total output capacity of all the power supplies", nil, []string{"memberId"}),
		}

		LogicalDriveMetrics = &metrics{
			"logicalDriveStatus": newServerMetric("xl420_logical_drive_status", "Current logical drive status 1 = OK, 0 = BAD, -1 = DISABLED", nil, []string{"name", "logicalDriveNumber", "raid"}),
		}

		PhysicalDriveMetrics = &metrics{
			"physicalDriveStatus": newServerMetric("xl420_physical_drive_status", "Current physical drive status 1 = OK, 0 = BAD, -1 = DISABLED", nil, []string{"name", "id", "location", "serialnumber"}),
		}

		MemoryMetrics = &metrics{
			"memoryStatus": newServerMetric("xl420_memory_status", "Current memory status 1 = OK, 0 = BAD", nil, []string{"totalSystemMemoryGiB"}),
		}

		Metrics = &map[string]*metrics{
			"thermalMetrics":      ThermalMetrics,
			"powerMetrics":        PowerMetrics,
			"logicalDriveMetrics": LogicalDriveMetrics,
			"driveMetrics":        PhysicalDriveMetrics,
			"memoryMetrics":       MemoryMetrics,
		}
	)

	return Metrics
}
