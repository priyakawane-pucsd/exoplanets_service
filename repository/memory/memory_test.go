package memory

import (
	"context"
	"exoplanetservice/models/dao"
	"exoplanetservice/models/filters"
	"reflect"
	"testing"
)

func TestRepository_CreateExoplanets(t *testing.T) {
	type fields struct {
		cfg        *Config
		exoplanets map[string]dao.Exoplanets
	}
	type args struct {
		ctx       context.Context
		exoplanet *dao.Exoplanets
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "jupiter",
			fields: fields{
				exoplanets: make(map[string]dao.Exoplanets),
			},
			args: args{
				ctx: context.Background(),
				exoplanet: &dao.Exoplanets{
					Name: "Jupiter",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				cfg:        tt.fields.cfg,
				exoplanets: tt.fields.exoplanets,
			}
			if err := r.CreateExoplanets(tt.args.ctx, tt.args.exoplanet); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateExoplanets() error = %v, wantErr %v", err, tt.wantErr)
			}
			// check if exoplanet exists in map
			if r.exoplanets[tt.args.exoplanet.ID] != *tt.args.exoplanet {
				t.Errorf("Repository.CreateExoplanets() error = %v, wantErr %v", "exoplanet not saved in map", tt.wantErr)

			}
		})
	}
}

func TestRepository_GetExoplanets(t *testing.T) {
	type fields struct {
		cfg        *Config
		exoplanets map[string]dao.Exoplanets
	}
	type args struct {
		ctx    context.Context
		filter *filters.ExoplanetFilter
		limit  int
		offset int
	}
	exoplanets := map[string]dao.Exoplanets{"1": {ID: "1", Name: "earth"}, "2": {ID: "2", Name: "plueto", Mass: 4.5}}
	r1 := exoplanets["1"]
	r2 := exoplanets["2"]
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dao.Exoplanets
		wantErr bool
	}{
		{
			name: "first",
			fields: fields{
				exoplanets: exoplanets,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filters.ExoplanetFilter{},
				limit:  1,
				offset: 0,
			},
			want:    []*dao.Exoplanets{&r1},
			wantErr: false,
		},
		{
			name: "second",
			fields: fields{
				exoplanets: exoplanets,
			},
			args: args{
				ctx: context.Background(),
				filter: &filters.ExoplanetFilter{
					Mass: 4.5,
				},
				limit:  1,
				offset: 0,
			},
			want:    []*dao.Exoplanets{&r2},
			wantErr: false,
		},
		{
			name: "third",
			fields: fields{
				exoplanets: exoplanets,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filters.ExoplanetFilter{},
				limit:  1,
				offset: 1,
			},
			want:    []*dao.Exoplanets{&r2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				cfg:        tt.fields.cfg,
				exoplanets: tt.fields.exoplanets,
			}
			got, err := r.GetExoplanets(tt.args.ctx, tt.args.filter, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetExoplanets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetExoplanets() = %v, want %v", got, tt.want)
			}
		})
	}
}
