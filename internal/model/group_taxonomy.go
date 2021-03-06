package model

import "github.com/satori/go.uuid"

type GroupTaxonomy struct {
  Id uuid.UUID  `sql:"type:uuid,default:uuid_generate_v4()"`
  Type string `sql:",notnull"`
  Name string `sql:",notnull"`
}
