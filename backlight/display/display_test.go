/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package display

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var edid0 = []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0, 0x30, 0xe4, 0x5e, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18, 0x1, 0x4, 0x95, 0x1f, 0x11, 0x78, 0xea, 0xeb, 0xf5, 0x95, 0x59, 0x54, 0x90, 0x27, 0x1e, 0x50, 0x54, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0xd0, 0x1d, 0x56, 0xf4, 0x50, 0x0, 0x16, 0x30, 0x30, 0x20, 0x35, 0x0, 0x36, 0xae, 0x10, 0x0, 0x0, 0x19, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xfe, 0x0, 0x4c, 0x47, 0x20, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0xa, 0x20, 0x20, 0x0, 0x0, 0x0, 0xfe, 0x0, 0x4c, 0x50, 0x31, 0x34, 0x30, 0x57, 0x48, 0x38, 0x2d, 0x54, 0x50, 0x44, 0x31, 0x0, 0x25}

var edid1 = []byte{0x0, 0xff}

func getControllerByName(cs []*Controller, name string) *Controller {
	for _, c := range cs {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func Test_list(t *testing.T) {
	controllers, err := list("./testdata")
	require.NoError(t, err)
	assert.Len(t, controllers, 2)
	controller := getControllerByName(controllers, "intel_backlight")
	assert.Equal(t, controller.MaxBrightness, 937)
	assert.Len(t, controller.DeviceEDID, 128)

	br, err := controller.GetBrightness()
	require.NoError(t, err)
	assert.Equal(t, br, 100)

	abr, err := controller.GetActualBrightness()
	require.NoError(t, err)
	assert.Equal(t, abr, 100)

	controller = getControllerByName(controllers, "acpi_video0")
	assert.Equal(t, controller.Type, ControllerTypeFirmware)
	assert.Equal(t, controller.MaxBrightness, 15)
	require.Nil(t, controller.DeviceEDID)

	br, err = controller.GetBrightness()
	require.NoError(t, err)
	assert.Equal(t, br, 1)

	abr, err = controller.GetActualBrightness()
	require.NoError(t, err)
	assert.Equal(t, abr, 1)

	controller = controllers.GetByEDID(edid0)
	assert.Equal(t, controller.Name, "intel_backlight")

	controller = controllers.GetByEDID(edid1)
	require.Nil(t, controller)
}

func TestList(t *testing.T) {
	controllers, err := List()
	require.NoError(t, err)
	t.Log(err)
	for _, c := range controllers {
		t.Logf("%+v\n", c)
		br, _ := c.GetBrightness()
		t.Log("brightness", br)

		abr, _ := c.GetActualBrightness()
		t.Log("actual_brightness", abr)
	}
}
