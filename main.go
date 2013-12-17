package main

import (
	"github.com/skiesel/expsys/rdb"
	"github.com/skiesel/expsys/plots/standardplots"
	"github.com/skiesel/expsys/tables/standardtables"
)

func main() {
	filters := []map[string]string{
		map[string]string{
			"algorithm": "astar",
		},
		map[string]string{
			"algorithm": "speedy",
		},
	}

	names := []string{"A*", "Speedy",}

	dss := rdb.GetDatasetsWithPathKeys("data", filters, names)

	standardplots.PlotSolutionCosts("Solution Costs", dss)
	standardplots.PlotSolutionCostsFactorOfBest("Solution Costs Factor Best", dss)

	standardtables.SolutionCostSumsTable(dss)
	standardtables.SolutionCostAveragesTable(dss)
}