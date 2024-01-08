package config

type Server struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
