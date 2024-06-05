package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type DecisionNode struct {
	Feature     string
	Threshold   float64
	Left, Right *DecisionNode
	IsLeaf      bool
	LeafValue   string
}

func main() {
	// Create a decision tree
	root := &DecisionNode{
		Feature:   "X_2",
		Threshold: 1.9,
		Left: &DecisionNode{
			IsLeaf:    true,
			LeafValue: "Iris-setosa",
		},
		Right: &DecisionNode{
			Feature:   "X_3",
			Threshold: 1.5,
			Left: &DecisionNode{
				Feature:   "X_2",
				Threshold: 4.9,
				Left: &DecisionNode{
					IsLeaf:    true,
					LeafValue: "Iris-versicolor",
				},
				Right: &DecisionNode{
					IsLeaf:    true,
					LeafValue: "Iris-virginica",
				},
			},
			Right: &DecisionNode{
				Feature:   "X_2",
				Threshold: 5.0,
				Left: &DecisionNode{
					Feature:   "X_1",
					Threshold: 2.8,
					Left: &DecisionNode{
						IsLeaf:    true,
						LeafValue: "Iris-virginica",
					},
					Right: &DecisionNode{
						IsLeaf:    true,
						LeafValue: "Iris-versicolor",
					},
				},
				Right: &DecisionNode{
					IsLeaf:    true,
					LeafValue: "Iris-virginica",
				},
			},
		},
	}

	// Create a new plot
	p := plot.New()

	// Set the title and axis labels
	p.Title.Text = "Decision Tree"
	p.X.Label.Text = "Feature"
	p.Y.Label.Text = "Threshold"

	// Create a plotter for the decision tree
	tree := DecisionTreePlotter{root, 0.5, 0.5, 1.0}

	// Add the decision tree plotter to the plot
	p.Add(tree)

	// Save the plot to a file
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "decision_tree.png"); err != nil {
		panic(err)
	}
}

type DecisionTreePlotter struct {
	Node       *DecisionNode
	X, Y, Size float64
}

func (dt DecisionTreePlotter) Plot(c draw.Canvas, plt *plot.Plot) {
	dt.plotNode(c, plt, dt.Node, dt.X, dt.Y, dt.Size)
}

func (dt DecisionTreePlotter) plotNode(c draw.Canvas, plt *plot.Plot, node *DecisionNode, x, y, size float64) {
	if node.IsLeaf {
		// Plot leaf node
		dt.drawText(c, node.LeafValue, x, y)
	} else {
		// Plot decision node
		text := fmt.Sprintf("%s <= %.1f", node.Feature, node.Threshold)
		dt.drawText(c, text, x, y)

		// Plot left and right child nodes
		dt.plotNode(c, plt, node.Left, x-size/2, y-size/2, size/2)
		dt.plotNode(c, plt, node.Right, x+size/2, y-size/2, size/2)

		// Draw lines connecting the nodes
		leftLine, err := plotter.NewLine(plotter.XYs{{X: x, Y: y}, {X: x - size/2, Y: y - size/2}})
		if err != nil {
			panic(err)
		}
		rightLine, err := plotter.NewLine(plotter.XYs{{X: x, Y: y}, {X: x + size/2, Y: y - size/2}})
		if err != nil {
			panic(err)
		}
		plt.Add(leftLine, rightLine)
	}
}

func (dt DecisionTreePlotter) drawText(c draw.Canvas, text string, x, y float64) {
	pt := vg.Point{X: vg.Length(x), Y: vg.Length(y)}
	c.FillText(draw.TextStyle{Color: color.Black}, pt, text)
}
