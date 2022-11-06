package main

import "testing"

func TestParse(t *testing.T) {
	test := struct {
		input  []string
		output []food
	}{
		input: []string{
			"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
			"trh fvjkl sbzzf mxmxvkd (contains dairy)",
			"sqjhc fvjkl (contains soy)",
			"sqjhc mxmxvkd sbzzf (contains fish)",
		},
		output: []food{
			{
				ingredients: []string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				allergens:   []string{"dairy", "fish"},
			},
			{
				ingredients: []string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				allergens:   []string{"dairy"},
			},
			{
				ingredients: []string{"sqjhc", "fvjkl"},
				allergens:   []string{"soy"},
			},
			{
				ingredients: []string{"sqjhc", "mxmxvkd", "sbzzf"},
				allergens:   []string{"fish"},
			},
		},
	}

	got := parseInput(test.input)

	if len(got) != len(test.output) {
		t.Fatalf("Lengths of got and test.output are not equal. Got: %d. Wanted: %d.", len(got), len(test.output))
	}

	for i, o := range test.output {
		g := got[i]
		if len(o.ingredients) != len(g.ingredients) {
			t.Fatalf("Lengths of ingredients are not equal. Got: %v. Wanted: %v.", g.ingredients, o.ingredients)
		}
		if len(o.allergens) != len(g.allergens) {
			t.Fatalf("Lengths of allergens are not equal. Got: %v. Wanted: %v.", g.allergens, o.allergens)
		}

		for j, ing := range g.ingredients {
			if o.ingredients[j] != ing {
				t.Fatalf("Ingredients are not equal. Got: %v. Wanted: %v.", ing, o.ingredients[j])
			}
		}

		for j, all := range g.allergens {
			if o.allergens[j] != all {
				t.Fatalf("Ingredients are not equal. Got: %v. Wanted: %v.", all, o.allergens[j])
			}
		}
	}
}

func TestIntersect(t *testing.T) {
	test := struct {
		input  [][]string
		output []string
	}{
		input: [][]string{
			[]string{
				"one", "two", "three", "three",
			},
			[]string{
				"five", "six", "three", "one",
			},
		},
		output: []string{
			"three", "one",
		},
	}

	got := intersect(test.input[0], test.input[1])

	if len(got) != len(test.output) {
		t.Fatalf("Lengths of got and test.output are not equal. Got: %d. Wanted: %d.", len(got), len(test.output))
	}

	for i, s := range got {
		if test.output[i] != s {
			t.Fatalf("Intersects are not equal. Got: %v. Wanted: %v.", got, test.output)
		}
	}
}
