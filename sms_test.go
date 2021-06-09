package sms_test

import (
	"sms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMS(t *testing.T) {
	t.Run("empty drawer disappoints in the morning", func(t *testing.T) {
		d := sms.Drawer{}
		socksForToday, err := d.Morning()
		assert.Error(t, err)
		assert.Empty(t, socksForToday)
	})

	t.Run("laundry without socks dissappoints as well", func(t *testing.T) {
		d := sms.Drawer{}

		d.Folding([]sms.Sock{})

		socksForToday, err := d.Morning()
		assert.Error(t, err)
		assert.Empty(t, socksForToday)
	})

	t.Run("successful life", func(t *testing.T) {
		d := sms.Drawer{}

		d.Folding([]sms.Sock{
			{"green"}, {"christmas"}, {"with bikes"}, {"christmas"}, {"with bikes"}, {"green"}},
		)

		socksForToday, err := d.Morning()
		assert.NoError(t, err)
		assert.NotEmpty(t, socksForToday)
	})

	t.Run("if sometimes you have lonely sock in your laundry - that's fine", func(t *testing.T) {
		d := sms.Drawer{}

		d.Folding([]sms.Sock{
			{"green"}, {"christmas"}, {"with bikes"}, {"christmas"}, {"with bikes"}},
		)

		socksForToday, err := d.Morning()
		assert.NoError(t, err)
		assert.NotEmpty(t, socksForToday)
	})

	t.Run("if you only have lonely socks in your laundry - it's not fine", func(t *testing.T) {
		d := sms.Drawer{}

		d.Folding([]sms.Sock{
			{"green"}, {"christmas"}, {"with bikes"}},
		)

		socksForToday, err := d.Morning()
		assert.Error(t, err)
		assert.Empty(t, socksForToday)
	})

	t.Run("if you have mismatched sock in the laundry and in the drawer - you should be fine", func(t *testing.T) {
		d := sms.Drawer{}

		d.Folding([]sms.Sock{
			{"green"}, {"christmas"}, {"with bikes"}, {"christmas"}, {"with bikes"}},
		)

		socksForToday, err := d.Morning()
		assert.NoError(t, err)
		assert.NotEmpty(t, socksForToday)

		socksForToday, err = d.Morning()
		assert.NoError(t, err)
		assert.NotEmpty(t, socksForToday)

		d.Folding([]sms.Sock{{"green"}})

		socksForToday, err = d.Morning()
		assert.NoError(t, err)
		assert.NotEmpty(t, socksForToday)
	})

}
