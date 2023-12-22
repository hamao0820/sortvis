package graph

type Graph struct {
	Name        string
	Description string
	Function    func() []int
}

var (
	GraphList = []Graph{
		{"sin", "sin(x) wave", sin},
		{"sin2", "sin(2x) wave", sin2},
		{"sin3", "sin(4x) wave", sin3},
		{"cos", "cos(x) wave", cos},
		{"cos2", "cos(2x) wave", cos2},
		{"cos3", "cos(4x) wave", cos3},
		{"sqr", "square", square},
		{"sqrI", "inverse square", squareInverse},
		{"step", "step function", step},
		{"stepI", "inverse step function", stepInverse},
		{"asc", "ascending", ascending},
		{"des", "descending", descending},
		{"par", "parabola(x^2)", parabola},
		{"parI", "inverse parabola(-x^2)", parabolaInverse},
		{"tri", "triangle", triangle},
		{"triI", "inverse triangle", triangleInverse},
	}
)
