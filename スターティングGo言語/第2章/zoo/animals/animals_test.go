package animals

import (
	"testing"
)

func TestElephantFeed(t *testing.T) {
	except := "Grass"
	actual := ElephantFeed()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}

func TestMonkeyFeed(t *testing.T) {
	except := "Banana"
	actual := MonkeyFeed()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}

func TestRabbitFeed(t *testing.T) {
	except := "Carrot"
	actual := RabbitFeed()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}
