package cache

import (
	"testing"
	"time"
)

func TestGetFromFastestCacheLayer(t *testing.T) {
	cacheLayers := make([]Cache, 0)
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(100) * time.Millisecond,
		Data:         map[string]string{},
	})
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(200) * time.Millisecond,
		Data: map[string]string{
			"key_1": "value_1",
		},
	})
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(400) * time.Millisecond,
		Data: map[string]string{
			"key_1": "value_1",
			"key_2": "value_2",
		},
	})

	tests := []struct {
		cacheLayers []Cache
		key         string
		want        string
		maxTookTime time.Duration
		name        string
		err         error
	}{
		{
			cacheLayers: cacheLayers,
			key:         "key_1",
			want:        "value_1",
			maxTookTime: time.Duration(200) * time.Millisecond,
			name:        "fastest cache layer is layer 2",
			err:         nil,
		},
		{
			cacheLayers: cacheLayers,
			key:         "key_2",
			want:        "value_2",
			maxTookTime: time.Duration(400) * time.Millisecond,
			name:        "fastest cache layer is layer 3",
			err:         nil,
		},
		{
			cacheLayers: cacheLayers,
			key:         "key_3",
			want:        "",
			maxTookTime: time.Duration(400) * time.Millisecond,
			name:        "cacheLayers dosn't have key",
			err:         ErrKeyNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			valueGot, err := GetValueFromFastestLayer(tt.cacheLayers, tt.key)
			elapsed := time.Since(start)

			if err != tt.err {
				t.Errorf(
					"Got Error is not correct, %s, want %s",
					err,
					tt.err,
				)
			}

			if elapsed-tt.maxTookTime > 10*time.Millisecond {
				t.Errorf(
					"Too slow! Took %v, max allowed %v, diff %v",
					elapsed,
					tt.maxTookTime,
					elapsed-tt.maxTookTime,
				)
			}

			if valueGot != tt.want {
				t.Errorf(
					"Got Value is not correct, %s, want %s",
					valueGot,
					tt.want,
				)
			}
		})
	}
}

func TestSetValueToAllCacheLayers(t *testing.T) {
	cacheLayers := make([]Cache, 0)
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(100) * time.Millisecond,
		Data:         map[string]string{},
	})
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(200) * time.Millisecond,
		Data:         map[string]string{},
	})
	cacheLayers = append(cacheLayers, Cache{
		ResponseTime: time.Duration(400) * time.Millisecond,
		Data:         map[string]string{},
	})

	tests := []struct {
		cacheLayers []Cache
		key         string
		value       string
		maxTookTime time.Duration
		name        string
	}{
		{
			cacheLayers: cacheLayers,
			key:         "key_1",
			value:       "value_1",
			maxTookTime: time.Duration(400) * time.Millisecond,
			name:        "set value to all cache layers",
		},
		{
			cacheLayers: cacheLayers,
			key:         "key_2",
			value:       "value_2",
			maxTookTime: time.Duration(400) * time.Millisecond,
			name:        "set value to all cache layers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			_ = SetValueToAllLayers(tt.cacheLayers, tt.key, tt.value)
			elapsed := time.Since(start)

			if elapsed-tt.maxTookTime > 10*time.Millisecond {
				t.Errorf(
					"Too slow! Took %v, max allowed %v, diff %v",
					elapsed,
					tt.maxTookTime,
					elapsed-tt.maxTookTime,
				)
			}

			for _, cache := range tt.cacheLayers {
				if cache.Data[tt.key] != tt.value {
					t.Errorf(
						"Got Value is not correct, %s, want %s",
						tt.cacheLayers[0].Data[tt.key],
						tt.value,
					)
				}
			}
		})
	}
}
