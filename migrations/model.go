package migrations

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"primaryKey,size:256"`
	Name           string
	Email          string
	Password       string
	Phone          string
	Photo          string
	Status         *bool
	UserRole       []*UserRole
	UserPermission []*UserPermission
	PondID         *uuid.UUID `gorm:"size:256"`
	Pond           Pond
	orm.OrmModel
}

type Role struct {
	ID             uuid.UUID `gorm:"primaryKey,size:256"`
	Code           string
	Name           string
	Scope          string
	RolePermission []*RolePermission
	orm.OrmModel
}

type Permission struct {
	ID             uuid.UUID `gorm:"primaryKey,size:256"`
	Code           string
	Name           string
	Path           string
	RolePermission []*RolePermission
	orm.OrmModel
}

type RolePermission struct {
	ID             uuid.UUID `gorm:"primaryKey,size:256"`
	RoleID         uuid.UUID `gorm:"size:256"`
	Role           Role
	PermissionID   uuid.UUID `gorm:"size:256"`
	PermissionName string
	PermissionPath string
	Permission     Permission
	orm.OrmModel
}

type UserRole struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256"`
	UserID uuid.UUID `gorm:"size:256"`
	User   User
	RoleID uuid.UUID `gorm:"size:256"`
	Role   Role
	orm.OrmModel
}

type UserPermission struct {
	ID             uuid.UUID `gorm:"primaryKey,size:256"`
	UserID         uuid.UUID `gorm:"size:256"`
	User           User
	PermissionID   uuid.UUID `gorm:"size:256"`
	PermissionPath string
	PermissionName string
	Permission     Permission
	orm.OrmModel
}

type Team struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	Name       string
	CountryID  uuid.UUID `gorm:"size:256"`
	ProvinceID uuid.UUID `gorm:"size:256"`
	CityID     uuid.UUID `gorm:"size:256"`
	DistrictID uuid.UUID `gorm:"size:256"`
	Detail     string
	Note       string
	ListPond   []*Pond
	orm.OrmModel
}

type Pond struct {
	ID            uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	OwnerName     string    `json:"ownerName"`
	Name          string    `json:"name"`
	CountryID     uuid.UUID `gorm:"size:256" json:"country_id"`
	ProvinceID    uuid.UUID `gorm:"size:256" json:"province_id"`
	CityID        uuid.UUID `gorm:"size:256" json:"city_id"`
	DistrictID    uuid.UUID `gorm:"size:256" json:"district_id"`
	DetailAddress string    `json:"detailAddress"`
	NoteAddress   string    `json:"noteAddress"`
	Type          string    `json:"type"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	TeamID        uuid.UUID `gorm:"size:256" json:"team_id"`
	Team          Team      `json:"team"`
	Status        string    `json:"status"`
	Image         string    `json:"image"`
	ListPool      []Pool    `json:"list_pool"`
	ListBerkas    []Berkas  `json:"berkas"`
	orm.OrmModel
}

type Berkas struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID `gorm:"size:256"`
	Pond   Pond
	Name   string `json:"name"`
	File   string `json:"file"`
	orm.OrmModel
}

type Pool struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID `gorm:"size:256" json:"pond_id"`
	Pond   Pond      `json:"pond"`
	Name   string    `json:"name"`
	Long   float64   `json:"long"`
	Wide   float64   `json:"wide"`
	Image  string    `json:"image"`
	orm.OrmModel
}

type Budidaya struct {
	ID              uuid.UUID `gorm:"primaryKey,size:256"`
	PoolID          uuid.UUID `gorm:"size:256"`
	Pool            Pool
	DateOfSeed      time.Time
	FishSpeciesID   uuid.UUID `gorm:"size:256"`
	FishSpecies     FishSpecies
	FishSpeciesName string
	EstTonase       float64
	EstPanenDate    time.Time
	EstPrice        int
	Status          string
	orm.OrmModel
}

type FishSpecies struct {
	ID       uuid.UUID `gorm:"primaryKey,size:256"`
	Name     string
	Budidaya []*Budidaya
	orm.OrmModel
}

type Order struct {
	ID          uuid.UUID `gorm:"primaryKey,size:256"`
	Code        string
	BudidayaID  uuid.UUID `gorm:"size:256"`
	Budidaya    model.Budidaya
	UserID      uuid.UUID `gorm:"size:256"`
	User        User
	Qty         int
	BookingDate time.Time
	Status      string
	orm.OrmModel
}
