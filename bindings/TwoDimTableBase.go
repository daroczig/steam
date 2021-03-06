/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package bindings

type TwoDimTableBase struct {
	*Schema
	/** Table Name */
	Name string `json:"name"`

	/** Table Description */
	Description string `json:"description"`

	/** Column Specification */
	Columns []*ColumnSpecsBase `json:"columns"`

	/** Number of Rows */
	Rowcount int32 `json:"rowcount"`

	/** Table Data (col-major) */
	Data [][]Polymorphic `json:"data"`
}

func NewTwoDimTableBase() *TwoDimTableBase {
	return &TwoDimTableBase{
		Name:        "",
		Description: "",
		Columns:     nil,
		Rowcount:    0,
		Data:        nil,
	}
}
